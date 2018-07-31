package main

// https://en.wikipedia.org/wiki/Master_Password
// https://en.wikipedia.org/wiki/Scrypt

import (
	"crypto/hmac"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"

	"golang.org/x/crypto/scrypt"
)

const Base string = "com.filipovicz.vault"

const min = "abcdefghijklmnopqrstuvwxyz"
const max = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const under = "-_/"
const number = "0123456789"
const space = " "
const special = "!@#$%^&*()=+,.?:;{}[]`~"

func newMasterKey(name, passphrase string) ([]byte, error) {
	salt := Base + strconv.Itoa(len(name)) + name
	dk, err := scrypt.Key([]byte(passphrase), []byte(salt), 1<<15, 8, 1, 64)
	if err != nil {
		return nil, err
	}
	return dk, nil
}

func newSeed(mk []byte, message string) ([]byte, error) {
	sig := hmac.New(sha512.New, mk)
	sig.Write([]byte(message))
	return sig.Sum(nil), nil
}

func newMessage(service string, length int) string {
	return Base + strconv.Itoa(len(service)) + service + strconv.Itoa(length)
}

func newPassword(seed []byte, length int) (string, error) {
	if length > 64 {
		return "", fmt.Errorf("length > 64")
	}

	chars := min + max + number + under + space + special
	charsLength := len(chars)

	password := ""
	for index := 0; index < length; index++ {
		r := seed[index]

		if int(r) < charsLength {
			c := chars[int(r)]
			password = password + string(c)
		} else {
			m := math.Mod(float64(r), float64(charsLength))
			password = password + string(chars[int(m)])
		}
	}

	return password, nil
}

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

	mk, err := newMasterKey(*name, *passphrase)
	if err != nil {
		log.Fatal(err)
	}
	seed, _ := newSeed(mk, newMessage(*service, *length))
	if err != nil {
		log.Fatal(err)
	}
	password, err := newPassword(seed, *length)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(password)
}
