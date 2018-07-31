package main

// https://en.wikipedia.org/wiki/Master_Password
// https://en.wikipedia.org/wiki/Scrypt

import (
	"encoding/base64"
	"flag"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func newPassword(salt, passphrase string, length int) ([]byte, error) {
	dk, err := scrypt.Key([]byte(passphrase), []byte(salt), 1<<15, 8, 1, length)
	if err != nil {
		return nil, err
	}

	return dk, nil
}

func checkFlag(flag string) error {
	if flag == "" {
		return fmt.Errorf("is empty")
	}

	return nil
}

func main() {
	fmt.Println("Stateless Secure Password Generator")

	name := flag.String("name", "", "your name")
	service := flag.String("service", "", "the name of the service")
	passphrase := flag.String("passphrase", "", "your passphrase")
	length := flag.Int("length", 20, "password length")
	flag.Parse()

	fmt.Println(*name)
	fmt.Println(*service)
	fmt.Println(*passphrase)

	password, _ := newPassword(*name+*service, *passphrase, *length)
	fmt.Println(base64.StdEncoding.EncodeToString(password))
}
