package cookie

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"haibaracode/pkg/config"
)

const (
	CookieName     string = "haibaracode"
	CookieMaxAge   int    = 7 * 60 * 3600
	CookiePath     string = "/"
	CookieDomain   string = "localhost"
	CookieSecure   bool   = false
	CookieHttpOnly bool   = true
)

var (
	hashKey  = []byte(config.Cookie.HashKey)
	blockKey = []byte(config.Cookie.BlockKey)
	s        = securecookie.New(hashKey, blockKey)
)

func SetCookie(c *gin.Context, value string) error {
	encoded, err := s.Encode(CookieName, value)
	if err == nil {
		c.SetCookie(CookieName, encoded, CookieMaxAge, CookiePath, CookieDomain, CookieSecure, CookieHttpOnly)
	}
	return err
}

func ReadCookie(c *gin.Context) (string, error) {
	if cookieValue, err := c.Cookie(CookieName); err == nil {
		value := ""
		if err = s.Decode(CookieName, cookieValue, &value); err == nil {
			return value, nil
		}
	}
	return "", errors.New("获取 Cookie 失败。")
}
