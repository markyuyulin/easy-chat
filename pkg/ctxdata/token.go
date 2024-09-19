package ctxdata

import "github.com/golang-jwt/jwt"

const Identify = "imooc.com"

// secretKey秘钥
// iat生成token的时间
// seconds这是token过期的时间
// uid
func GetJwtToken(secretKey string, iat, seconds int64, uid string) (string, error) {
	claims := make(jwt.MapClaims)
	//过期时间
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[Identify] = uid

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
