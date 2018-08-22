package generator_test

import (
	"testing"

	"github.com/filipovi/vault/generator"
	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	expected := "F-x+e,C/K41R6{pHzg30GJY9/[UR6(LN"
	p, _ := generator.NewPassword("name", "passphrase", "service", 32, 1, "")
	assert.Equal(t, expected, p)
}
