package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"hash"
	"os"
	"strings"
	"time"
)

type Token struct {
	Exp time.Time `json:"exp"`
	Id  uint      `json:"id"`
}

const (
	TypeToken        = 1
	TypeRefreshToken = 2
)

func NewToken(id uint, typ int) string {
	t := Token{Id: id}
	if typ == TypeToken {
		t.Exp = time.Now().Add(time.Hour)
	} else {
		t.Exp = time.Now().Add(time.Hour * 24 * 15)
	}

	jt, _ := json.Marshal(t)
	payload := base64.RawURLEncoding.EncodeToString(jt)

	var h hash.Hash
	if typ == TypeToken {
		h = hmac.New(sha256.New, []byte(os.Getenv("TOKEN_SECRET")))
	} else {
		h = hmac.New(sha256.New, []byte(os.Getenv("REFRESH_TOKEN_SECRET")))
	}
	h.Write([]byte(payload))
	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	return payload + "." + signature
}

func CheckToken(tokenStr string, typ int) (id uint, err error) {
	tokenSli := strings.Split(tokenStr, ".")
	if len(tokenSli) != 2 {
		err = errors.New("token没有点或者多于1个点")
		return
	}
	payload, signature := tokenSli[0], tokenSli[1]

	var h hash.Hash
	if typ == TypeToken {
		h = hmac.New(sha256.New, []byte(os.Getenv("TOKEN_SECRET")))
	} else {
		h = hmac.New(sha256.New, []byte(os.Getenv("REFRESH_TOKEN_SECRET")))
	}
	h.Write([]byte(payload))
	TrueSignature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	if signature != TrueSignature {
		err = errors.New("token被修改过或者密匙被更改")
		return
	}

	jt, _ := base64.RawURLEncoding.DecodeString(payload)
	var token Token
	if e := json.Unmarshal(jt, &token); e != nil {
		err = errors.New("token字符串绑定Token结构体失败")
		return
	}

	if token.Exp.Before(time.Now()) {
		err = errors.New("token过期") //这行里面的内容不要随便改，因为返回数据时会根据字符串内容而报身份已过期
		return
	}

	id = token.Id
	return
}
