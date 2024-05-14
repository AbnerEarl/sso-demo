/**
 * @author: yangchangjia
 * @email 1320259466@qq.com
 * @date: 2024/4/26 11:05
 * @desc: about the role of class.
 */

package task

import (
	"encoding/json"
	"fmt"
	"github.com/AbnerEarl/goutils/dbs"
	"github.com/AbnerEarl/goutils/gins"
	"github.com/AbnerEarl/goutils/times"
	"sso-demo/model"
	"strconv"
	"time"
)

func DealLogData(data gins.LogData) {
	info, _ := json.Marshal(data.LogInfo)
	params, _ := json.Marshal(data.RequestParams)
	logInfo := model.LogModel{
		LogInfo:       string(info),
		RequestParams: string(params),
		RequestToken:  data.RequestToken,
		ResponseData:  data.ResponseData,
	}

	model.DB.Create(&logInfo)
}

func ClearLogData(logCycle string) {
	cycle := 6
	result, e := strconv.Atoi(logCycle)
	if e == nil && result > 0 {
		cycle = result
	} else {
		gins.LogWarn("the 'log.cycle' config parameter is error, it will be ignored.")
	}
	DeleteLogData(cycle)
}

func DeleteLogData(logCycle int) {
	validTime := time.Now().AddDate(0, -logCycle, 0).Format(times.TmFmtWithMS1)
	model.DB.Transaction(func(tx *dbs.TX) error {
		sql := fmt.Sprintf("delete from log_info where created_at < '%s';", validTime)
		return tx.Exec(sql).Error
	})

	model.DB.Transaction(func(tx *dbs.TX) error {
		sql := fmt.Sprintf("delete from login_log where created_at < '%s';", validTime)
		return tx.Exec(sql).Error
	})

	model.DB.Transaction(func(tx *dbs.TX) error {
		sql := fmt.Sprintf("delete from operate_log where created_at < '%s';", validTime)
		return tx.Exec(sql).Error
	})

	model.DB.Transaction(func(tx *dbs.TX) error {
		sql := fmt.Sprintf("delete from node_log where created_at < '%s';", validTime)
		return tx.Exec(sql).Error
	})

	model.DB.Transaction(func(tx *dbs.TX) error {
		sql := fmt.Sprintf("delete from link_log where created_at < '%s';", validTime)
		return tx.Exec(sql).Error
	})
}
