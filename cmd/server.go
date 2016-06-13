package cmd

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

type serverconfig struct {
	httpsPort           uint
	grpcPort            uint
	httpsAddr           string
	grpcAddr            string
	concurrency         uint
	queuesize           uint
	vaultTLSCertPath    string
	vaultTLSKeyPath     string
	tlsCert             []byte
	tlsKey              []byte
	logToSumo           bool
	sumoURL             string
	vaultSumoURLPath    string
	healthcheckHTTPport uint
}

var serverConfig serverconfig
var kafkaConfig kafkaconfig

var logger *log.Logger

var version = "0"
var description = "unknown"

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run Furan server",
	Long:  `Furan API server (see docs)`,
	Run:   server,
}

func init() {
	serverCmd.PersistentFlags().UintVar(&serverConfig.httpsPort, "https-port", 4000, "REST HTTPS TCP port")
	serverCmd.PersistentFlags().UintVar(&serverConfig.grpcPort, "grpc-port", 4001, "gRPC TCP port")
	serverCmd.PersistentFlags().UintVar(&serverConfig.healthcheckHTTPport, "healthcheck-port", 4002, "Healthcheck HTTP port (listens on localhost only)")
	serverCmd.PersistentFlags().StringVar(&serverConfig.httpsAddr, "https-addr", "0.0.0.0", "REST HTTPS listen address")
	serverCmd.PersistentFlags().StringVar(&serverConfig.grpcAddr, "grpc-addr", "0.0.0.0", "gRPC listen address")
	serverCmd.PersistentFlags().UintVar(&serverConfig.concurrency, "concurrency", 10, "Max concurrent builds")
	serverCmd.PersistentFlags().UintVar(&serverConfig.queuesize, "queue", 100, "Max queue size for buffered build requests")
	serverCmd.PersistentFlags().StringVar(&serverConfig.vaultTLSCertPath, "tls-cert-path", "/tls/cert", "Vault path to TLS certificate")
	serverCmd.PersistentFlags().StringVar(&serverConfig.vaultTLSKeyPath, "tls-key-path", "/tls/key", "Vault path to TLS private key")
	serverCmd.PersistentFlags().BoolVar(&serverConfig.logToSumo, "log-to-sumo", true, "Send log entries to SumoLogic HTTPS collector")
	serverCmd.PersistentFlags().StringVar(&serverConfig.vaultSumoURLPath, "sumo-collector-path", "/sumologic/url", "Vault path SumoLogic collector URL")
	RootCmd.AddCommand(serverCmd)
}

func setupServerLogger() {
	var url string
	if serverConfig.logToSumo {
		url = serverConfig.sumoURL
	}
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalf("error getting hostname: %v", err)
	}
	stdlog := NewStandardLogger(os.Stderr, url)
	logger = log.New(stdlog, fmt.Sprintf("%v: ", hn), log.LstdFlags)
}

// Separate server because it's HTTP on localhost only
// (simplifies Consul health check)
func healthcheck() {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods("GET")
	addr := fmt.Sprintf("127.0.0.1:%v", serverConfig.healthcheckHTTPport)
	server := &http.Server{Addr: addr, Handler: r}
	logger.Printf("HTTP healthcheck listening on: %v", addr)
	logger.Println(server.ListenAndServe())
}

func server(cmd *cobra.Command, args []string) {
	setupVault()
	if serverConfig.logToSumo {
		getSumoURL()
	}
	setupServerLogger()
	setupDB(initializeDB)
	setupKafka()
	certPath, keyPath := writeTLSCert()
	defer rmTempFiles(certPath, keyPath)
	err := getDockercfg()
	if err != nil {
		logger.Fatalf("error reading dockercfg: %v", err)
	}

	startgRPC()
	go healthcheck()

	r := mux.NewRouter()
	r.HandleFunc("/", versionHandler).Methods("GET")
	r.HandleFunc("/build", buildRequestHandler).Methods("POST")
	r.HandleFunc("/build/{id}", buildStatusHandler).Methods("GET")
	r.HandleFunc("/build/{id}", buildCancelHandler).Methods("DELETE")

	tlsconfig := &tls.Config{MinVersion: tls.VersionTLS12}
	addr := fmt.Sprintf("%v:%v", serverConfig.httpsAddr, serverConfig.httpsPort)
	server := &http.Server{Addr: addr, Handler: r, TLSConfig: tlsconfig}
	logger.Printf("HTTPS REST listening on: %v", addr)
	logger.Println(server.ListenAndServeTLS(certPath, keyPath))
}

func setupVersion() {
	bv := make([]byte, 20)
	bd := make([]byte, 2048)
	fv, err := os.Open("VERSION.txt")
	if err != nil {
		return
	}
	defer fv.Close()
	sv, err := fv.Read(bv)
	if err != nil {
		return
	}
	fd, err := os.Open("DESCRIPTION.txt")
	if err != nil {
		return
	}
	defer fd.Close()
	sd, err := fd.Read(bd)
	if err != nil {
		return
	}
	version = string(bv[:sv])
	description = string(bd[:sd])
}
