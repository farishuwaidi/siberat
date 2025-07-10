package helper

import (
	"crypto/md5"
	"encoding/binary"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// for register
func HashPassword(password string) string {
	passwordHash,_ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHash)
}

func GenerateHexID() int64 {
	uid := uuid.New().String()[:9]
	h := md5.Sum([]byte(uid))
	id := int64(binary.BigEndian.Uint64(h[:8]))
	return id
}
// for login
func VerifyPassword(HashPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(HashPassword), []byte(password))
	return err
}
