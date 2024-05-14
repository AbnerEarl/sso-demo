/**
 * @author: yangchangjia
 * @email 1320259466@qq.com
 * @date: 2024/5/13 14:51
 * @desc: about the role of class.
 */

package model

import (
	"github.com/AbnerEarl/goutils/dbs"
)

type LogModel struct {
	dbs.BaseModel
	LogInfo       string `json:"log_info" gorm:"column:log_info;null;comment:'日志详情'"`
	RequestParams string `json:"request_params" gorm:"column:request_params;null;comment:'请求参数'"`
	RequestToken  string `json:"request_token" gorm:"column:request_token;null;comment:'请求token'"`
	ResponseData  string `json:"response_data" gorm:"column:response_data;null;comment:'相应数据'"`
}

func (m *LogModel) TableName() string {
	return "log_info"
}
