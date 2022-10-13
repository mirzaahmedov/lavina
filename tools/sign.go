package tools

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"net/http"
)

func CompareSign(sign string, r *http.Request, body []byte, secret string) bool {
	protocol := "http"

	if r.TLS != nil {
		protocol += "s"
	}

	payload := r.Method + protocol + "://" + r.Host + r.RequestURI + string(body) + secret

	sum := make([]byte, 16)
	if _, err := hex.Decode(sum, []byte(sign)); err != nil {
		return false
	} else {
		result := md5.Sum([]byte(payload))
		return bytes.Equal(result[:], sum)
	}
}