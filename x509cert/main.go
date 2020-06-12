package main

import (
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		panic("require x509 certificate file")
	}

	certFile := args[0]

	f, err := os.Open("ca.crt")
	if err != nil {
		panic(err.Error())
	}
	b, err := ioutil.ReadAll(f)
	pool := x509.NewCertPool()
	ok := pool.AppendCertsFromPEM(b)
	if !ok {
		panic("failed to parse root certificate")
	}
	f.Close()

	f, err = os.Open(certFile)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	b, err = ioutil.ReadAll(f)
	if err != nil {
		panic(err.Error())
	}

	block, _ := pem.Decode(b)
	if block == nil {
		panic("block is nil")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Organization: ")
	for i, v := range cert.Subject.Organization {
		var format string
		if i == 0 {
			format = "%s\n"
		} else {
			format = "              %s\n"
		}
		fmt.Printf(format, v)
	}
	fmt.Printf("CommonName: %s\n", cert.Subject.CommonName)

	fmt.Printf("DNS: ")
	for i, v := range cert.DNSNames {
		var format string
		if i == 0 {
			format = "%s\n"
		} else {
			format = "     %s\n"
		}
		fmt.Printf(format, v)
	}

	fmt.Printf("Not Before: %s\n", cert.NotBefore.String())
	fmt.Printf("Not After: %s\n", cert.NotAfter.String())

	hostname := "gucchi.info"

	if err := cert.VerifyHostname(hostname); err == nil {
		fmt.Println("VerifyHostname: OK")
	} else {
		fmt.Printf("VerifyHostname: NG, %s\n", err.Error())
	}

	opts := x509.VerifyOptions{
		Roots: pool,
		DNSName: hostname,
	}
	if _, err := cert.Verify(opts); err == nil {
		fmt.Println("Verify: OK")
	} else {
		fmt.Printf("Verify: NG, %s\n", err.Error())
	}
}
