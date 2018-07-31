package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/filipovi/vault/generator"
)

func checkFlag(flag string) error {
	if flag == "" {
		return fmt.Errorf("%s is empty", flag)
	}

	return nil
}

func main() {
	fmt.Println("Stateless Secure Password Generator")

	name := flag.String("name", "", "your name")
	service := flag.String("service", "", "the name of the service")
	passphrase := flag.String("passphrase", "", "your passphrase")
	length := flag.Int("length", 35, "password length")
	flag.Parse()

	if err := checkFlag(*name); err != nil {
		log.Fatal(err)
	}

	if err := checkFlag(*service); err != nil {
		log.Fatal(err)
	}

	if err := checkFlag(*passphrase); err != nil {
		log.Fatal(err)
	}

	password, err := generator.NewPassword(*name, *passphrase, *service, *length)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(password)
}
