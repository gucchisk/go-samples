package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	// "net"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	tlsConfig := &tls.Config{}
	key, err := ioutil.ReadFile("./server.key")
	if err != nil {
		log.Panic("key error")
	}
	certpem, err := ioutil.ReadFile("./server.crt")
	if err != nil {
		log.Panic("cert error")
	}
	cert, err := tls.X509KeyPair(certpem, key)
	if err != nil {
		log.Panic("key paier error")
	}
	tlsConfig.Certificates = make([]tls.Certificate, 1)
	tlsConfig.Certificates[0] = cert
	tlsConfig.MinVersion = tls.VersionTLS12
	tlsConfig.CipherSuites = []uint16{
		tls.TLS_AES_128_GCM_SHA256,
		tls.TLS_AES_256_GCM_SHA384,
		tls.TLS_CHACHA20_POLY1305_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
	}
	tlsConfig.PreferServerCipherSuites = true
	tlsConfig.SessionTicketsDisabled = true

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("called")
		w.Write([]byte("hello"))
	})
	addr := "localhost:8080"
	// addr := ":8080"
	s := &http.Server{
		Addr: addr,
		Handler: mux,
		// TLSConfig: tlsConfig,
	}
	// listener, err := net.Listen("tcp", "localhost:8080")
	// if err != err {
	// 	log.Panic("listen error")
	// }
	// s.Serve(listener)
	// s.Serve(tls.NewListener(listener, s.TLSConfig))
	s.ListenAndServe()
	
}
