package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("qsjo387865tuifewu4857")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}

// // Valid implements jwt.Claims.
// func (j *JWTClaim) Valid() error {
// 	panic("unimplemented")
// }

func (c *JWTClaim) Valid() error {
	// Tambahkan logika validasi klaim di sini, jika diperlukan
	return nil
}
