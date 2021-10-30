//go:build server

package main

import (
	"github.com/parthpower/loonabot/cert"
	"github.com/parthpower/loonabot/cmd/runner"
	"google.golang.org/grpc/credentials"

	"crypto/tls"
	"fmt"
	"os"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	crt, err := cert.Certs.ReadFile("server-cert.pem")
	if err != nil {
		return nil, err
	}
	k, err := cert.Certs.ReadFile("server-key.pem")
	if err != nil {
		return nil, err
	}
	serverCert, err := tls.X509KeyPair(crt, k)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	if len(os.Args) < 1 {
		panic("usage: ./server 0.0.0.0:12345")
	}
	creds, err := loadTLSCredentials()
	if err != nil {
		panic(err)
	}
	fmt.Println("listening")
	if runner.Start(os.Args[1], creds) != nil {
		panic("err")
	}
}
