package wechat

/**
 * Created by zc on 2019/12/9.
 */
type Engine interface {
	Sent(drive Drive)
	Exec()
}

type Drive interface {
	BuildAccessToken() (string, error)
	GetAppId() string
	GetAppSecret() string
}

type LocalAccess struct {
	AccessToken string `json:"access_token"` // 凭证
	ExpireAt    int64    `json:"expire_at"`    // 过期unix时间戳
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值。
	ResCode
}

type ResCode struct {
	ErrCode     int    `json:"errcode"`      // 错误码
	ErrMsg      string `json:"errmsg"`       // 错误信息
}

type RGB struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}