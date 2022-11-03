package keys

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Olegas/dnstools/pkg/config"
	"github.com/golang-jwt/jwt/v4"
)

func signedToken() string {
	serviceAccountID := config.GetServiceAccountID()
	keyID := config.GetKeyID()
	claims := jwt.RegisteredClaims{
		Issuer:    serviceAccountID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Audience:  []string{"https://iam.api.cloud.yandex.net/iam/v1/tokens"},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodPS256, claims)
	token.Header["kid"] = keyID

	privateKey := loadPrivateKey()
	signed, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}
	return signed
}

func loadPrivateKey() *rsa.PrivateKey {
	keyFile := config.GetKeyFileName()
	data, err := ioutil.ReadFile(keyFile)
	if err != nil {
		panic(err)
	}
	rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(data)
	if err != nil {
		panic(err)
	}
	return rsaPrivateKey
}

func GetIAMToken() string {
	jot := signedToken()
	resp, err := http.Post(
		"https://iam.api.cloud.yandex.net/iam/v1/tokens",
		"application/json",
		strings.NewReader(fmt.Sprintf(`{"jwt":"%s"}`, jot)),
	)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		panic(fmt.Sprintf("%s: %s", resp.Status, body))
	}
	var data struct {
		IAMToken string `json:"iamToken"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	return data.IAMToken
}
