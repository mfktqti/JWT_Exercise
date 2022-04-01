package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/*
JWT 全称JSON Web Token 是一种跨域认证解决方案,属于一个开放的标准,它规定了一种Token实现方式,目前多用于前后端分离项目
*/

func main() {

	fmt.Println("hello jwt ")

	s, err := GenToken("jason zhou ", "123456", time.Now().Add(time.Hour*2).Unix())
	if err != nil {
		fmt.Println("生成Token出错:", err)
		return
	}
	token, err := ParseToken(s)
	if err != nil {
		fmt.Println("解析Token出错:", err)
		return
	}

	fmt.Printf("token.Username: %v\n", token.Username)
	fmt.Printf("token.Password: %v\n", token.Password)

	s, err = GenToken("jason zhou", "123546", time.Now().Add(time.Second).Unix())

	if err != nil {
		fmt.Println("生成Token出错:", err)
		return
	}
	//延时3秒,等待Token过期
	time.Sleep(time.Second * 3)
	token, err = ParseToken(s)
	if err != nil {
		fmt.Println("解析Token出错:", err)
		return
	}

	fmt.Printf("token.Username: %v\n", token.Username)
	fmt.Printf("token.Password: %v\n", token.Password)
}

var key = []byte("secret")

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenToken(username, password string, expirseAt int64) (string, error) {
	c := MyClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expirseAt,    //过期时间
			Issuer:    "jason zhou", //签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c) //使用指定的签名方法 创建签名对象
	//使用指定的secret签名并获得完整的编码后的字符串Token
	//注意这个地方一定要是字节切片,不能是字符串
	return token.SignedString(key)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("解析Token出错1:", err, err.Error())
		fmt.Printf("%t", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")

}
