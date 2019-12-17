package xcx

import (
	"github.com/zc2638/gotool/curlx"
	"github.com/zc2638/wechat"
)

/**
 * Created by zc on 2019/12/10.
 */
// 登录凭证校验
type Code2Session struct {
	Code   string
	Result Code2SessionResult
}

type Code2SessionResult struct {
	Openid     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionId    string `json:"unionid"`     // 用户在开放平台的唯一标识符
	wechat.ResCode
}

func (a *Code2Session) Exec(drive wechat.Drive) error {
	h := curlx.HttpReq{
		Url: drive.GetHost() + "/sns/jscode2session",
		Query: map[string]string{
			"appid":      drive.GetAppId(),
			"secret":     drive.GetAppSecret(),
			"js_code":    a.Code,
			"grant_type": "authorization_code",
		},
	}
	return h.Do().ParseJSON(&a.Result)
}

// 检验数据的真实性，并且获取解密后的明文. TODO Result可以结构化
type DecryptData struct {
	EncryptedData string
	SessionKey    string
	Iv            string
	Result        []byte
}

func (a *DecryptData) Exec(drive wechat.Drive) error {
	var err error
	a.Result, err = wechat.AesDecrypt(a.EncryptedData, a.SessionKey, a.Iv)
	return err
}
