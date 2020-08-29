// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jwt

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

func TestClaimToken(t *testing.T) {
	fmt.Printf("\n=================生成Token=================\n")
	tk, err := Generate("abc")
	if err != nil {
		fmt.Println("sorry", err)
		return
	}
	fmt.Println(tk)

	fmt.Printf("\n=================Token解码=================\n")
	items := strings.Split(tk, ".")
	header, _ := base64.URLEncoding.DecodeString(items[0])
	payload, _ := base64.URLEncoding.DecodeString(items[1])
	fmt.Println(string(header))
	fmt.Println(string(payload))
	fmt.Println(items[2])

	fmt.Printf("\n=================Token反解析=================\n")
	token, err := Resolve(tk)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", token)
}

func TestParse(t *testing.T) {
	var token *jwt.Token
	//  修改并验证claims校验项
	var claims Claims = Claims{
		StandardClaims: jwt.StandardClaims{
			Audience:  "abc",
			ExpiresAt: 1591755941,
		},
		Version: 159166954,
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString([]byte(JWTSECRET))

	if tokenStr != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhYmMiLCJleHAiOjE1OTE3NTU5NDEsIlZlciI6MTU5MTY2OTU0fQ.dtvFC8OJ4Hy9n9mUj33zEpS1zvUO92IuO6-KN9eYxc0" {
		t.Errorf("sorry,%s", tokenStr)
		return
	}
}

func TestBearer(t *testing.T) {
	tk := "Bearer Bearer er xxx"
	tk = strings.TrimPrefix(tk, "Bearer ")
	fmt.Println(tk)
}
