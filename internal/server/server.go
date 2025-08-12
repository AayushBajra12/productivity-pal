package server

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"

	"productivity-pal/internal/auth"
	"productivity-pal/internal/handlers"
)

func StartServer() error {

	svc := &handlers.Svc{ // Assuming DB is initialized in handlers package
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", svc.UserHandler)

	mux.HandleFunc("/refresh", auth.CorsMiddleware(auth.JwtMiddleware(auth.RefreshTokenHandler)))

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
		ClientAuth:   tls.NoClientCert,
		MinVersion:   tls.VersionTLS12,
	}, nil

}
