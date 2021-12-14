package lib

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aqaurius6666/go-utils/cryptography"
)

var (
	SEC_KEY = os.Getenv("SERVER_KEY")
)

func RandomPassword() *string {
	str := fmt.Sprintf("%d", time.Now().UnixMilli())
	out := base64.StdEncoding.EncodeToString([]byte(str))
	out = fmt.Sprintf("%10s", out)
	out = strings.ReplaceAll(out, "=", "")
	return &out
}

func HashPassword(in string, sec string) string {
	out := cryptography.Hash256(append([]byte(in), []byte(sec)...))
	return cryptography.BytesToBase64(out)
}

func MyHashPassword(in *string) *string {
	out := HashPassword(*in, SEC_KEY)
	return &out
}
