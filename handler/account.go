package handler

import (
	"context"
	"github.com/AbnerEarl/goutils/gins"
	"golang.org/x/oauth2"
	"net/http"
	"sso-demo/config"
)

// AccountLogin
// @Summary 用户登录
// @Description 功能描述
// @Tags 用户管理
// @Schemes http https
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param username body string true "用户账号"
// @Param password body string true "用户密码"
// @Param sid body string true "会话ID"
// @Param vericode body string true "验证码"
// @Success 200 {object} gins.Response "{"code":0,"data":{},"msg":"请求成功"}"
// @Failure 200 {object} gins.Response "{"code":1,"data":{},"msg":"内部错误"}"
// @Failure 200 {object} gins.Response "{"code":2,"data":{},"msg":"业务错误"}"
// @Router /account/login [get]
func AccountLogin(c *gins.Context) {
	_, err := config.SSOConfig.GetToken(c.Writer, c.Request)
	if err == nil {
		http.Redirect(c.Writer, c.Request, "/api/v1/account/info", http.StatusFound)
	} else {
		url := config.SSOConfig.AuthCodeURL("state", oauth2.AccessTypeOnline)
		http.Redirect(c.Writer, c.Request, url, http.StatusFound)
	}
}

// AccountCallback
// @Summary 登录回调
// @Description 功能描述
// @Tags 用户管理
// @Schemes http https
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} gins.Response "{"code":0,"data":{},"msg":"请求成功"}"
// @Failure 200 {object} gins.Response "{"code":1,"data":{},"msg":"内部错误"}"
// @Failure 200 {object} gins.Response "{"code":2,"data":{},"msg":"业务错误"}"
// @Router /account/callback [get]
func AccountCallback(c *gins.Context) {
	_, err := config.SSOConfig.GetToken(c.Writer, c.Request)
	if err == nil {
		http.Redirect(c.Writer, c.Request, "/api/v1/account/info", http.StatusFound)
	} else {
		code := c.Request.URL.Query().Get("code")
		if code == "" {
			http.Error(c.Writer, "Missing authorization code", http.StatusBadRequest)
			return
		}

		token, err := config.SSOConfig.Exchange(context.Background(), code)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusBadRequest)
			return
		}

		config.SSOConfig.SetTokenIntoCookie(c.Writer, c.Request, token, "", config.SSOConfig.CookieMaxAge)

		http.Redirect(c.Writer, c.Request, "/api/v1/account/info", http.StatusFound)
	}
}

// AccountLogout
// @Summary 用户退出
// @Description 功能描述
// @Tags 用户管理
// @Schemes http https
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} gins.Response "{"code":0,"data":{},"msg":"请求成功"}"
// @Failure 200 {object} gins.Response "{"code":1,"data":{},"msg":"内部错误"}"
// @Failure 200 {object} gins.Response "{"code":2,"data":{},"msg":"业务错误"}"
// @Router /account/logout [get]
func AccountLogout(c *gins.Context) {
	config.SSOConfig.DelTokenIntoCookie(c.Writer, c.Request)
	gins.SendResponse(c, nil, map[string]interface{}{
		"result":  "success",
		"message": "the user logout",
	})
}

// AccountInfo
// @Summary 个人信息
// @Description 功能描述
// @Tags 用户管理
// @Schemes http https
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} gins.Response "{"code":0,"data":{},"msg":"请求成功"}"
// @Failure 200 {object} gins.Response "{"code":1,"data":{},"msg":"内部错误"}"
// @Failure 200 {object} gins.Response "{"code":2,"data":{},"msg":"业务错误"}"
// @Router /account/info [get]
func AccountInfo(c *gins.Context) {
	info, err := config.SSOConfig.GetUserInfo(c.Writer, c.Request)
	if err != nil {
		gins.SendResponse(c, err, nil)
		return
	}
	gins.SendResponse(c, nil, info)
}
