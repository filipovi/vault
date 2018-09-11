package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	proto "github.com/filipovi/vault/api/proto"
	micro "github.com/micro/go-micro"
)

func checkFlag(key, flag string) error {
	if flag == "" {
		return fmt.Errorf("%s is empty", key)
	}

	return nil
}

func main() {
	fmt.Println("Stateless Secure Password Generator")

	name := flag.String("name", "", "your name")
	service := flag.String("service", "", "the name of the service")
	passphrase := flag.String("passphrase", "", "your passphrase")
	length := flag.Int("length", 35, "password length")
	counter := flag.Int("counter", 1, "counter")
	scope := flag.String("scope", "", "your scope")
	flag.Parse()

	if err := checkFlag("name", *name); err != nil {
		log.Fatal(err)
	}
	if err := checkFlag("service", *service); err != nil {
		log.Fatal(err)
	}
	if err := checkFlag("passphrase", *passphrase); err != nil {
		log.Fatal(err)
	}
	if err := checkFlag("scope", *scope); err != nil {
		log.Fatal(err)
	}

	s := micro.NewService()
	generator := proto.NewGeneratorService("master-password-generator", s.Client())
	newPassword, err := generator.NewPassword(context.TODO(), &proto.NewPasswordRequest{
		Name:       *name,
		Passphrase: *passphrase,
		Service:    *service,
		Length:     int32(*length),
		Counter:    int32(*counter),
		Scope:      *scope,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(newPassword.Password)
}
