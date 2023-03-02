package utils

import (
	"testing"
)

func TestHash(t *testing.T) {
	testPassword := "testPassword"
	hash, _ := HashPassword(testPassword)
	want_true := CheckPasswordHash(testPassword, hash)
	if want_true != true {
		t.Fatalf("Hash password comparison failed")
	}

}
