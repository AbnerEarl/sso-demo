package base

var Comment = map[string]map[string]interface{}{
	"account_info": map[string]interface{}{
		"###":         "###",
		"username":    "用户账号",
		"nickname":    "用户名称",
		"password":    "用户密码",
		"role":        "用户角色",
		"status":      "用户状态",
		"permissions": "用户权限",
		"last_login":  "最后登录时间",
	},
	"log_info": map[string]interface{}{
		"###":            "###",
		"log_info":       "日志详情",
		"request_params": "请求参数",
		"request_token":  "请求token",
		"response_data":  "相应数据",
	},
}
