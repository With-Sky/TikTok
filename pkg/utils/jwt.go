package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT(config Config) *JWT {
	return &JWT{
		[]byte(config.Viper.GetString("JWT.SigningKey")),
	}
}

type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func (j *JWT) GenToken(userId int64, config Config) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.Viper.GetInt("JWT.ExpiresTime")) * time.Second * 18).Unix(), // 过期时间
			Issuer:    "XDream",                                                                                        // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(config.Viper.GetString("JWT.SigningKey")))
}

// ParseToken 解析JWT
func (j *JWT) ParseToken(tokenString string, config Config) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(config.Viper.GetString("JWT.SigningKey")), nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
func (j *JWT) GetIdByToken(token string, config Config) (int64, error) {

	myClaims, err := j.ParseToken(token, config)
	if err != nil {
		return 0, err
	}
	return myClaims.UserID, nil
}
