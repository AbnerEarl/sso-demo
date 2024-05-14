package model

import (
	"github.com/AbnerEarl/goutils/dbs"
	"github.com/spf13/viper"
	"sso-demo/tools"
)

var DB *dbs.DB

func InitDb() {
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	ip := viper.GetString("db.ip")
	DB = dbs.OpenDBMySQL(username, password, ip, 3306, "anon", false, 0, 0, "")
	DB.Migration([]interface{}{
		&dbs.UpdateModel{},
		&AccountModel{},
		&LogModel{},
	})
	basePath := tools.AbPathByCaller() + "/../scheme/"
	DB.InitDefaultData(basePath)
}
