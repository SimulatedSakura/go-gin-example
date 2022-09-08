package util

import (
	"time"

	"github.com/SimulatedSakura/go-gin-example/pkg/setting"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 生成JWT的Token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	//首项参数为加密方法的选择
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// valid 验证基于时间的声明exp(Expiration Time), iat(Issued At), nbf(Not Before)
		// 注意如果没有任何声明在令牌中仍然会被认为是有效的。
		// 并且对于时区偏差没有计算方法用于验证
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
