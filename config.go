package params_validator

// Conf 配置文件
type Conf struct {
	ZhTrans bool `json:"zhTrans,optional,default=true"` // 是否开启中文, 默认中文
}
