package main

import (
	"context"
	"log"

	"github.com/filipovi/vault/generator"
	proto "github.com/filipovi/vault/proto"
	"github.com/micro/go-micro"
)

// Generator is a struct containing the password
type Generator struct{}

// NewPassword send the name preceded with NewPassword
func (g *Generator) NewPassword(ctx context.Context, req *proto.NewPasswordRequest, rsp *proto.NewPasswordResponse) error {
	password, err := generator.NewPassword(req.Name, req.Passphrase, req.Service, req.Length, req.Counter, req.Scope)
	if err == nil {
		rsp.Password = password
	}

	return err
}

func main() {
	log.Println("Starting...")
	defer log.Println("...bye bye!")

	service := micro.NewService(
		micro.Name("master-password-generator"),
	)

	proto.RegisterGeneratorHandler(service.Server(), new(Generator))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
