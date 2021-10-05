package helpers

import (
	"testing"
)

var stringLen = 12
var tokenLen = 32

func TestRandToken(t *testing.T) {
	token := RandToken()
	if len(token) != tokenLen {
		t.Error("wrong len token: ", len(token))
	}
}

func TestGenerateRandomString(t *testing.T) {
	randomString := GenerateRandomString(stringLen)
	if len(randomString) != stringLen {
		t.Error("wrong len string: ", len(randomString))
	}
}
