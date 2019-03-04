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

type ResCode struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type ResReturnCode struct {
	ReturnCode string `json:"return_code"` // 返回状态码
	ReturnMsg  string `json:"return_msg"`  // 返回信息
}
