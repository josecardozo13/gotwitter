package bd

import (
	"golang.org/x/crypto/bcrypt"
)

func EncriptarPassword(pass string) (string, error) {
	costo := 8 // el minimo es de 6 pasadas de encriptacion, pero vamos a dejar 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
