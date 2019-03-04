package xcx

type ResLocalToken struct {
	AccessToken string `json:"access_token"` // 凭证
	ExpireAt    int    `json:"expire_at"`    // 过期unix时间戳
}

type ResAccessToken struct {
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值。
	ResCode
}

type ResCode2Session struct {
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值。
	ResCode
}

type ResCode struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type ResReturnCode struct {
	ReturnCode string `json:"return_code"` // 返回状态码
	ReturnMsg  string `json:"return_msg"`  // 返回信息
}

type ResUnifiedOrder struct {
	Appid      string `json:"appid"`       // 调用接口提交的小程序ID
	MchId      string `json:"mch_id"`      // 调用接口提交的商户号
	DeviceInfo string `json:"device_info"` // 自定义参数，可以为请求支付的终端设备号等
	NonceStr   string `json:"nonce_str"`   // 微信返回的随机字符串
	PrePayId   string `json:"prepay_id"`   // 预支付交易会话标识
	ResultCode string `json:"result_code"` // 业务结果 SUCCESS/FAIL
	Sign       string `json:"sign"`        // 微信返回的签名值
	TradeType  string `json:"trade_type"`  // 交易类型
	CodeUrl    string `json:"code_url"`    // 二维码链接
	ResReturnCode
}
