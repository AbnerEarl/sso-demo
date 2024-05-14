package config

import "github.com/AbnerEarl/goutils/gins"

var (
	ErrVerifyCode      = &gins.Errno{Code: 10004, Message: "the verification code is error.", Tips: "验证码错误"}
	ErrVerifyLogin     = &gins.Errno{Code: 10005, Message: "the username or password is error.", Tips: "用户账号或密码错误"}
	ErrParam           = &gins.Errno{Code: 10008, Message: "the parameter is error.", Tips: "参数错误"}
	ErrPassword        = &gins.Errno{Code: 10009, Message: "the password was error.", Tips: "密码错误"}
	ErrPasswordDisable = &gins.Errno{Code: 100010, Message: "the password was disabled.", Tips: "密码被禁用"}
	ErrPasswordSame    = &gins.Errno{Code: 100011, Message: "the new password was same as the old password.", Tips: "新密码和旧密码相同"}
	ErrUserStatus      = &gins.Errno{Code: 100015, Message: "the user was forbid.", Tips: "用户被禁用"}
	ErrRepeatUrl       = &gins.Errno{Code: 100016, Message: "the url was repeated.", Tips: "网址重复"}
	ErrDomainUrl       = &gins.Errno{Code: 100017, Message: "the url is not a correct domain name.", Tips: "域名格式错误"}
	ErrNodeLink        = &gins.Errno{Code: 100018, Message: "no outgoing nodes or links were found.", Tips: "出网节点或链路没有配置"}

	ErrDataLength     = &gins.Errno{Code: 20004, Message: "the data length exceed limit.", Tips: "数据长度超出限制"}
	ErrDataIp         = &gins.Errno{Code: 20004, Message: "the ip was error.", Tips: "IP格式错误"}
	ErrDataPort       = &gins.Errno{Code: 20020, Message: "the port was error.", Tips: "端口配置错误"}
	ErrCertFile       = &gins.Errno{Code: 20008, Message: "the 'cert' file error.", Tips: "证书文件错误"}
	ErrNodeList       = &gins.Errno{Code: 20008, Message: "the node list info error.", Tips: "节点信息错误"}
	ErrPasswordLength = &gins.Errno{Code: 20011, Message: "the password should include digit,upper,lower,special character, the length 10 to 20.", Tips: "密码应该包含大写、小写、数字、特殊符合且10到20位"}
	ErrNameLength     = &gins.Errno{Code: 20013, Message: "the node name length exceed limit.", Tips: "节点名称长度超出限制"}
	ErrNodeHop        = &gins.Errno{Code: 20015, Message: "the node hop is error.", Tips: "节点跳数错误"}
	ErrNameExist      = &gins.Errno{Code: 20017, Message: "the name had exist.", Tips: "名称已经存在"}
	ErrMappingStatus  = &gins.Errno{Code: 20021, Message: "the using host is not be deleted.", Tips: "使用中的节点不允许删除"}

	ErrTokenInvalid = &gins.Errno{Code: 20103, Message: "the 'token' was invalid.", Tips: "Token信息无效"}
	ErrNetInfo      = &gins.Errno{Code: 20108, Message: "the network config is failed.", Tips: "网络配置错误"}
	ErrStartInfo    = &gins.Errno{Code: 20108, Message: "the node start is failed.", Tips: "节点启动出现错误"}
	ErrNodeAbnormal = &gins.Errno{Code: 20222, Message: "the node start is failed.", Tips: "节点出现异常错误"}
	ErrUpgrade      = &gins.Errno{Code: 20224, Message: "the upgrade is failed.", Tips: "升级出现错误"}
)
