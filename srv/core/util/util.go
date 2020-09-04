package util

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"math/rand"
	"time"
)

// rune is a unicode string type
// string
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandStringRunes ...
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// DeferClose ...
// defer DeferClose(&err, rsp.Body.Close)
// Note: https://github.com/carlmjohnson/json-tidy/blob/master/json-tidy.go#L91:L96
func DeferClose(c io.Closer, e *error) {
	var err error

	if err = c.Close(); err != nil {
		if *e == nil {
			*e = err
		}
	}
}

// HashString ...
func HashString(byteArr []byte) string {
	filenameHash := sha256.Sum256(byteArr)
	return base64.URLEncoding.EncodeToString(filenameHash[:])
}
