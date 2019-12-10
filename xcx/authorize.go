package xcx

import (
	"encoding/json"
	"github.com/zc2638/wechat"
	"github.com/zc2638/wechat/config"
	"github.com/zc2638/wechat/core"
	"github.com/zctod/go-tool/common/curlx"
)

/**
 * Created by zc on 2019/12/10.
 */
// 登录凭证校验
type Code2Session struct {
	appId  string
	secret string
	Code   string
	Err    error
	Result Code2SessionResult
}

type Code2SessionResult struct {
	Openid     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionId    string `json:"unionid"`     // 用户在开放平台的唯一标识符
	wechat.ResCode
}

func (a *Code2Session) Sent(drive wechat.Drive) {
	a.appId = drive.GetAppId()
	a.secret = drive.GetAppSecret()
}

func (a *Code2Session) Exec() {

	h := curlx.HttpReq{
		Url: config.XCX_CODE2SESSION,
		Query: map[string]string{
			"appid":      a.appId,
			"secret":     a.secret,
			"js_code":    a.Code,
			"grant_type": "authorization_code",
		},
	}

	b, err := h.Get()
	if err != nil {
		a.Err = err
		return
	}

	var res Code2SessionResult
	if err := json.Unmarshal(b, &res); err != nil {
		a.Err = err
		return
	}
	a.Result = res
}

// 检验数据的真实性，并且获取解密后的明文. TODO Result可以结构化
type DecryptData struct {
	EncryptedData string
	SessionKey    string
	Iv            string
	Err           error
	Result        []byte
}

func (a *DecryptData) Sent(drive wechat.Drive){}

func (a *DecryptData) Exec() {
	a.Result, a.Err = core.AesDecrypt(a.EncryptedData, a.SessionKey, a.Iv)
}
