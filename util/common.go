package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"time"
)

// SignatureResource 对资源进行签名(若name为空则不签名)
func SignatureResource(typ, name, expiresString string) string {
	if name == "" {
		return ""
	}
	if expiresString == "" {
		expires := time.Now().Add(time.Hour * 24)
		expiresString = strconv.FormatInt(expires.Unix(), 10)
	}

	h := hmac.New(sha256.New, []byte(os.Getenv("SIGNATURE_SECRET")))
	h.Write([]byte(typ + name + expiresString))
	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	return fmt.Sprintf("%s/%s/%s?expires=%s&signature=%s", os.Getenv("RESOURCE_URL"), typ, name, expiresString, signature)
}
