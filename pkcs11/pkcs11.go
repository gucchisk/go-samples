package main

import (
	"fmt"
	"os"
	"github.com/miekg/pkcs11"
)

func main() {
	tokenPin := "1234"

	modulePath := "/usr/local/opt/softhsm/lib/softhsm/libsofthsm2.so"

	var slotId uint = 0x4da79941

	p := pkcs11.New(modulePath)
	p.Initialize()

	tokenInfo, err := p.GetTokenInfo(slotId)
	fmt.Println("token label: " + tokenInfo.Label)
	fmt.Println("token model: " + tokenInfo.Model)
	
	session, err := p.OpenSession(slotId, pkcs11.CKF_SERIAL_SESSION | pkcs11.CKF_RW_SESSION)

	err = p.Login(session, pkcs11.CKU_USER, tokenPin)

	if err != nil {
		fmt.Print(err)
	}

	privateKeyTemplate := []*pkcs11.Attribute{
		pkcs11.NewAttribute(pkcs11.CKA_CLASS, pkcs11.CKO_PRIVATE_KEY),
	}

	err = p.FindObjectsInit(session, privateKeyTemplate)
	if err != nil {
		fmt.Print(err)
	}

	objects, _, err := p.FindObjects(session, 1)

	if err != nil {
		fmt.Print(err)
	}

	privateKey := objects[0]

	err = p.FindObjectsFinal(session)

	if err != nil {
		fmt.Print(err)
	}

	file, err := os.Open("data.txt")
	// data := []byte("Hello, world")
	if err != nil {
		fmt.Print(err)
	}
	buf := make([]byte, 64)
	n, err := file.Read(buf)
	if n == 0 {
		fmt.Print("error")
	}
	if err != nil {
		fmt.Print(err)
	}
	data := buf[:n]

	var signature []byte
	err = p.SignInit(session, []*pkcs11.Mechanism{pkcs11.NewMechanism(pkcs11.CKM_SHA256_RSA_PKCS, nil)}, privateKey)

	signature, err = p.Sign(session, data)

	file, err = os.Create("signature.txt")

	if err != nil {
		fmt.Print(err)
	}

	file.Write(signature)

}


