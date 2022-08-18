package utils

import (
	"testing"
)

func TestGenerateSalt(t *testing.T) {
	salt := GenerateSalt()
	t.Log(salt)
}

func TestEncryptPasswordHash(t *testing.T) {
	salt := GenerateSalt()
	saltHash := EncryptPasswordHash("password", salt)

	t.Log(saltHash)
	t.Log(salt)
}

func TestValifyPasswordHash(t *testing.T) {
	t.Log(ValifyPasswordHash("password", "d95928e8babe0fb2e3afb9ce70c713a6", "1a5dcdbb0552cf7334ed8b670f433402a6a7d375badbbf949d2637162efb565e"))
}
