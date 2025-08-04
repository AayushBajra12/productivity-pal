package server

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"

	"productivity-pal/backend/internal/db"
	"productivity-pal/backend/internal/handlers"
)

func StartServer() error {

	svc := &handlers.Svc{
		Db: db.DB, // Initialize your database connection here
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", svc.UserHandler)
	mux.HandleFunc("/signup", svc.SignupHandler)

	tlsConfig, err := configureTLS()
	if err != nil {
		log.Println("error while configuring tls: ", err)
		return err
	}

	server := &http.Server{
		Addr:      addr,
		Handler:   mux,
		TLSConfig: tlsConfig,
	}

	log.Println("Starting server!!!")
	err = server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Println("error while starting server: ", err)
		return err
	}

	return nil
}

func configureTLS() (*tls.Config, error) {

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Println("error loading cert and key files: ", err)
		return nil, err
	}

	caCert, err := os.ReadFile(caCertFile)
	if err != nil {
		log.Println("error loading ca cert files: ", err)
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		MinVersion:   tls.VersionTLS12,
	}, nil

}
