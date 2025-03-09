package servertype

import (
	"crypto/tls"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/acme/autocert"
	"log"
)

func StartProductionServer(app *fiber.App) {
	m := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("api.example.com"), // Change to your domain
		Cache:      autocert.DirCache("./certs"),
	}

	cfgs := &tls.Config{
		GetCertificate: m.GetCertificate,
		NextProtos:     []string{"http/1.1", "acme-tls/1"},
	}

	ln, err := tls.Listen("tcp", ":443", cfgs)
	if err != nil {
		log.Fatalf("Failed to create TLS listener: %v", err)
	}

	log.Println("Server running in PRODUCTION mode on https://api.example.com")
	log.Fatal(app.Listener(ln))
}

func StartLocalHttpsServer(app *fiber.App) {
	certFile := "certs/cert.pem"
	keyFile := "certs/key.pem"

	// Ensure self-signed cert exists (Run `mkcert -key-file key.pem -cert-file cert.pem localhost` if missing)
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to load SSL certificate: %v", err)
	}

	cfgs := &tls.Config{Certificates: []tls.Certificate{cert}}

	ln, err := tls.Listen("tcp", "localhost:443", cfgs)
	if err != nil {
		log.Fatalf("Failed to create local TLS listener: %v", err)
	}

	log.Println("Server running in LOCAL mode on https://localhost:443")
	log.Fatal(app.Listener(ln))
}

func StartLocalHttpServer(app *fiber.App) {
	log.Println("Server running in LOCAL mode on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
