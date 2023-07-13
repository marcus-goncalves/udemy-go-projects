package nap

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
)

type AuthToken struct {
	Token string
}

type AuthBasic struct {
	Username string
	Password string
}

type Authentication interface {
	AuthorizationHeader() string
}

func NewAuthToken(token string) *AuthToken {
	return &AuthToken{
		Token: token,
	}
}

func NewAuthBasic(username, password string) *AuthBasic {
	return &AuthBasic{
		Username: username,
		Password: password,
	}
}

func (a *AuthToken) AuthorizationHeader() string {
	return fmt.Sprintf("token %s", a.Token)
}

func (a *AuthBasic) AuthorizationHeader() string {
	buffer := &bytes.Buffer{}
	enc := base64.NewEncoder(base64.URLEncoding, buffer)
	encBody := fmt.Sprintf("%s:%s", a.Username, a.Password)

	enc.Write([]byte(encBody))
	enc.Close()

	content, err := io.ReadAll(buffer)
	if err != nil {
		log.Fatalln("read failed:", err)
	}

	return fmt.Sprintf("Basic %s", string(content))
}
