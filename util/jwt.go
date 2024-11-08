package utils

import (
	"EStarExchange/global"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtInfo struct {
	Username string `json:"username"`
	Role     uint   `json:"role"`
}

type JwtClaims struct {
	JwtInfo
	jwt.RegisteredClaims
}

// 为新登录的用户创建一个加密的token
func CreateToken(user JwtInfo) (string, error) {
	// 创建Claims对象
	claims := JwtClaims{
		JwtInfo{
			Username: user.Username,
			Role:     user.Role,
		},
		jwt.RegisteredClaims{
			Issuer:    global.Config.Jwt.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.Jwt.Expiration_time) * time.Second)),
		},
	}
	// 创建一个新的token对象，使用HS256算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//fmt.Println("Time now is ", time.Now())
	//fmt.Println("Token expires at: ", claims.ExpiresAt.Time)
	// 使用密钥签名并生成token字符串
	tokenString, err := token.SignedString([]byte(global.Config.Jwt.Secret_key))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AnalysToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// 将 Secret_key 转换为 []byte
			return []byte(global.Config.Jwt.Secret_key), nil
		})
	if err != nil {
		global.Log.Errorln(err)
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效token")
}
