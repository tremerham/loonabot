//go:build client

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"github.com/parthpower/loonabot/cmd/runner"

	_ "embed"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//go:embed cert/ca-cert.pem
var pemServerCA []byte

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	if len(os.Args) < 2 {
		panic("usage: ./client localhost:12345 ./loonabot")
	}
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	dial, err := grpc.Dial(os.Args[1], grpc.WithTransportCredentials(tlsCredentials), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer dial.Close()
	fmt.Println("sending...")
	err = runner.UploadExec(dial, os.Args[2])
	if err != nil {
		panic(err)
	}
	err = runner.Restart(dial)
	if err != nil {
		panic(err)
	}
}
