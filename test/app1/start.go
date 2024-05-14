package main

import (
	"github.com/AbnerEarl/goutils/gins"
	"github.com/AbnerEarl/goutils/redisc"
	"github.com/AbnerEarl/sso"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"os/signal"
	"sso-demo/config"
	_ "sso-demo/docs"
	"sso-demo/model"
	"sso-demo/router"
	"sso-demo/task"
	"syscall"
)

var (
	cfg = pflag.StringP("config", "c", "conf/config1.yaml", "project config")
)

// @title 项目API文档
// @version 1.0
// @description 项目前后端联调及测试API文档。
// @termsOfService https://github.com
// @contact.name API Support
// @contact.url http://www.cnblogs.com
// @contact.email ×××@qq.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name token
// @BasePath /api/v1
func main() {
	pflag.Parse()
	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	mode := viper.GetString("web.runmode")
	addr := viper.GetString("web.addr")
	logCycle := viper.GetString("log.cycle")

	exitChan := make(chan int)
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		exitChan <- 1
	}()

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGKILL)

	model.InitDb()

	go task.PeriodicTasks(logCycle)

	config.SSOConfig = sso.OidcConfig{
		OidcProvider: "http://0.0.0.0:5556/provider",
		Config: &oauth2.Config{
			ClientID:     "aaa-client",
			ClientSecret: "ZXhhbXBsZS1hcHAtc2VjcmV1",
			RedirectURL:  "http://www.aaa.com:8080/api/v1/account/callback",
			Scopes:       []string{"openid", "offline_access", "profile", "email", "groups"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "http://127.0.0.1:5556/provider/auth",
				TokenURL: "http://127.0.0.1:5556/provider/token",
			},
		},
		AppTopDomain: []string{".aaa.com", ".bbb.com"},
		//RedisClusterCli:  redisc.InitRedisCluster([]string{"127.0.0.1:6379"}, 100, 10, "", ""),
		RedisCli: redisc.InitRedis("127.0.0.1:6379", 0, 100, 10, "", "ycj52011"),
	}

	g := gins.DefaultServer(mode)
	g.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Load(g)
	go func() {
		zap.L().Info(http.ListenAndServe(addr, g).Error())
	}()

	<-exitChan

}
