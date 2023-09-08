package entity

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	IdUser    int    `json:"id_user"`
	KodeAkses string `json:"kode_akses"`
	jwt.RegisteredClaims
}
