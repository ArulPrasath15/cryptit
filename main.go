package main

import (
	"fmt"

	"github.com/arulprasath15/cryptit/decrypt"
	"github.com/arulprasath15/cryptit/encrypt"
	"github.com/pborman/uuid"
)

func main() {
	pass := "hello"
	encryptStr := encrypt.Nimbus(pass)
	fmt.Println("Encrypted String: ", encryptStr)
	decryptStr := decrypt.NimbusDecrypt(encryptStr)
	fmt.Println("Decrypted String: ", decryptStr)

	uuid := uuid.New()
	fmt.Println("UUID: ", uuid)
}
