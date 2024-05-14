package config

import (
	"github.com/AbnerEarl/goutils/redisc"
	"github.com/AbnerEarl/sso"
	"golang.org/x/oauth2"
)

var (
	DefaultLimit    = 10
	DefaultMaxLimit = 1000
)

var (
	SSOConfig = sso.OidcConfig{
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
)
