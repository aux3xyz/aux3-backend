package common

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"

    "github.com/akamensky/base58"
)

func GenerateKey() (string, string) {
    key, err := rsa.GenerateKey(rand.Reader, 4096)
    if err != nil {
        return "", ""
    }
    privbytes := x509.MarshalPKCS1PrivateKey(key)
    pubbytes := x509.MarshalPKCS1PublicKey(&key.PublicKey)
    return base58.Encode(privbytes), base58.Encode(pubbytes)
}
