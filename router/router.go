package router

import (
	"github.com/AbnerEarl/goutils/gins"
	"net/http"
	"sso-demo/base"
	"sso-demo/handler"
	"sso-demo/task"
)

func Load(s *gins.Server) *gins.Server {
	s.NoRoute(func(c *gins.Context) {
		c.JSON(http.StatusNotFound, map[string]interface{}{"message": "the api path is incorrect."})
	})
	s.GET("/ping", func(c *gins.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{"message": "pong"})
	})
	apiV1 := s.Group("/api/v1")
	apiV1.Use(gins.Cors(), gins.Args(), base.CheckToken(), gins.LogAop(task.DealLogData))
	{
		//docs
		apiV1.GET("/docs", handler.DocsComment)

		//account
		apiV1.GET("/account/login", handler.AccountLogin)
		apiV1.GET("/account/callback", handler.AccountCallback)
		apiV1.GET("/account/logout", handler.AccountLogout)
		apiV1.GET("/account/info", handler.AccountInfo)

	}
	return s
}
