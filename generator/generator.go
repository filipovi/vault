package generator

// https://en.wikipedia.org/wiki/Master_Password
// https://en.wikipedia.org/wiki/Scrypt

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"math"

	"golang.org/x/crypto/scrypt"
)

const min = "abcdefghijklmnopqrstuvwxyz"
const max = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const under = "-_/"
const number = "0123456789"
const space = " "
const special = "!@#$%^&*()=+,.?:;{}[]`~"

/*
	Your identity is defined by the master key.

	The master key is the cryptographic result of two components:
	1. Your <name> (identification)
	2. Your <master password> (authentication)
*/
func newMasterKey(name, passphrase, scope string) ([]byte, error) {
	salt := bytes.Buffer{}
	salt.Write([]byte(scope))
	salt.WriteString(fmt.Sprint(len(name)))
	salt.Write([]byte(name))
	mk, err := scrypt.Key([]byte(passphrase), salt.Bytes(), 32768, 8, 2, 64)
	if err != nil {
		return nil, err
	}
	return mk, nil
}

/*
	The site key is a derivative from the master key when it is used to unlock a specific site.

	The site key is the result of three components:
	1. The <site name> (identification)
	2. Your <master key> (authentication)
	3. The <site counter>

	The site counter ensures you can easily create new keys for the site should a key become compromised.
	Together, they create a cryptographic identifier that is unique to your account at this site.
*/
func newSiteKey(mk []byte, service string, counter int, scope string) ([]byte, error) {
	seed := bytes.Buffer{}
	seed.Write([]byte(scope))
	seed.WriteString(fmt.Sprint(len(service)))
	seed.Write([]byte(service))
	seed.WriteString(fmt.Sprint(counter))
	sig := hmac.New(sha512.New, mk)
	sig.Write(seed.Bytes())
	return sig.Sum(nil), nil
}

/*
	Your  site  password  is  an  identifier  derived  from  your  site  key  in  compliance  with  the site’s password policy.
	The purpose of this step is to render the site’s cryptographic key into a format that the site’s password input will accept.
	Master  Password  declares  several  site  password  formats  and  uses  these  pre-defined password “templates” to render the site key legible.
*/
func newSitePassword(seed []byte, length int) (string, error) {
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

// NewPassword generates a Master Password based on https://masterpassword.app/masterpassword-algorithm.pdf
func NewPassword(name, passphrase, service string, length int, counter int, scope string) (string, error) {
	mk, err := newMasterKey(name, passphrase, scope)
	if err != nil {
		return "", err
	}
	sk, _ := newSiteKey(mk, service, counter, scope)
	if err != nil {
		return "", err
	}
	return newSitePassword(sk, length)
}
