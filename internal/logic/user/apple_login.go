package user

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"math/big"
	"strings"
	"sviwo/internal/consts/enums"
)

const (
	PublicKeyReqUrl     = "https://appleid.apple.com/auth/keys"
	AppleUrl            = "https://appleid.apple.com"
	ApplicationClientId = "com.sviwo.atv"
)

type (
	JwtClaims struct {
		jwt.StandardClaims
	}
	JwtHeader struct {
		Kid string `json:"kid"`
		Alg string `json:"alg"`
	}
	JwtKeys struct {
		Kty string `json:"kty"`
		Kid string `json:"kid"`
		Use string `json:"use"`
		Alg string `json:"alg"`
		N   string `json:"n"`
		E   string `json:"e"`
	}
)

func VerifyIdentityToken(cliToken, cliUserID string) error {
	cliTokenArr := strings.Split(cliToken, ".")
	if len(cliTokenArr) < 3 {
		return gerror.NewCode(enums.IdentityTokenFormatError)
	}
	cliHeader, err := jwt.DecodeSegment(cliTokenArr[0])
	if err != nil {
		return err
	}
	var jHeader JwtHeader
	err = json.Unmarshal(cliHeader, &jHeader)
	if err != nil {
		return err
	}
	token, err := jwt.ParseWithClaims(cliToken, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		pk := GetRSAPublicKey(jHeader.Kid)
		return pk, nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		if claims.Issuer != AppleUrl || claims.Audience != ApplicationClientId || claims.Subject != cliUserID {
			return gerror.NewCode(enums.IdentityTokenVerifyError)
		}
	} else {
		return gerror.NewCode(enums.IdentityTokenParseFail)
	}
	return nil
}

func GetRSAPublicKey(kid string) *rsa.PublicKey {
	var body []byte
	resp, err := g.Client().Get(context.Background(), PublicKeyReqUrl)
	if err != nil {
		panic(err)
	}
	defer func(resp *gclient.Response) {
		err = resp.Close()
		if err != nil {
			panic(err)
		}
	}(resp)
	_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	body = _body

	var jKeys map[string][]JwtKeys
	err = json.Unmarshal(body, &jKeys)
	if err != nil {
		return nil
	}
	var pubKey rsa.PublicKey
	for _, data := range jKeys {
		for _, val := range data {
			if val.Kid == kid {
				n_bin, _ := base64.RawURLEncoding.DecodeString(val.N)
				n_data := new(big.Int).SetBytes(n_bin)

				e_bin, _ := base64.RawURLEncoding.DecodeString(val.E)
				e_data := new(big.Int).SetBytes(e_bin)

				pubKey.N = n_data
				pubKey.E = int(e_data.Uint64())
				break
			}
		}
	}
	if pubKey.E <= 0 {
		return nil
	}
	return &pubKey
}
