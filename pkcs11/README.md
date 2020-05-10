# PKCS#11 with softhsm

- data.txt
    - input
    - plain data
- signature.txt
    - output
	- signature (RSASSA-PKCS1-v1_5)
	- encrypted data

### generate key pair

#### private key
```
$ openssl genrsa 2024 > private.key
```

#### public key
```
$ openssl rsa -pubout < private.key > public.key
```

### softhsm (register private key)

1. initialize
```
$ softhsm2-util --init-token --slot 0 --label test --so-pin 4321 --pin 1234
```
1. convert key
```
$ openssl pkcs8 -topk8 -inform PEM -outform PEM -nocrypt -in private.key -out private.pem
```
1. import private key
```
$ softhsm2-util --import private.pem --token test --label key --id 0000 --pin 1234
```

### sign
```
$ go run pkcs11.go
```

### verify
```
$ openssl dgst -sha256 -verify public.key -signature signature.txt data.txt
```
