package base

import (
	"fmt"
	"github.com/AbnerEarl/goutils/gins"
	"sso-demo/config"
	"time"
)

var WhitelistApi = map[string]bool{
	"/api/v1/account/login":    true,
	"/api/v1/account/callback": true,
	"/api/v1/account/logout":   true,
}

func CheckToken() gins.HandlerFunc {
	return func(c *gins.Context) {
		apiPath := c.Request.URL.Path
		if WhitelistApi[apiPath] {
			c.Next()
			return
		}
		err := config.SSOConfig.VerifyToken(c.Writer, c.Request)
		if err != nil {
			c.Abort()
			gins.SendResponse(c, config.ErrTokenInvalid, nil)
			return
		}
		value, err := config.SSOConfig.GetValue(c.Writer, c.Request, "expiry")
		if err != nil {
			c.Abort()
			gins.SendResponse(c, config.ErrTokenInvalid, nil)
			return
		}
		expiry, err := time.Parse(time.RFC3339, value)
		if err != nil {
			c.Abort()
			gins.SendResponse(c, config.ErrTokenInvalid, nil)
			return
		}
		if expiry.Before(time.Now()) {
			c.Abort()
			gins.SendResponse(c, config.ErrTokenInvalid, nil)
			return
		}

		if expiry.Unix()-time.Now().Unix() < 0 {
			fmt.Println("超过了制定时间。。。。。。。。。。。")
			c.Abort()
			gins.SendResponse(c, config.ErrTokenInvalid, nil)
			return
		} else if expiry.Unix()-time.Now().Unix() < 10 {
			fmt.Println("开始刷新token。。。。。。。。。。。")
			config.SSOConfig.RefreshToken(c.Writer, c.Request)
		}
		fmt.Println("不需要刷新token。。。。。。。。。。。")
		c.Next()
		return
	}
}
