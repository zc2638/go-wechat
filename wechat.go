package wechat

import (
	"encoding/json"
	"errors"
	"github.com/zc2638/gotool/curlx"
	"github.com/zc2638/gotool/utilx"
	"github.com/zc2638/wechat/config"
	"io/ioutil"
	"time"
)

/**
 * Created by zc on 2019/12/9.
 */
type WeChat struct {
	appId           string
	secret          string
	accessTokenPath string
}

func NewWeChat(appId, secret string) *WeChat {
	return &WeChat{appId: appId, secret: secret}
}

// 获取公众号全局唯一后台接口调用凭据（access_token）。
func (w *WeChat) BuildAccessToken() (string, error) {

	filePath := config.AccessTokenPath
	if w.accessTokenPath != "" {
		filePath = w.accessTokenPath
	}

	var local LocalAccess
	b, err := ioutil.ReadFile(filePath)
	if err == nil {
		err = json.Unmarshal(b, &local)
		if err == nil {
			nowTs := time.Now().Unix()
			if local.ExpireAt > nowTs {
				return local.AccessToken, nil
			}
		}
	}

	h := curlx.HttpReq{
		Url: config.WECHAT_ACCESSTOKEN,
		Query: map[string]string{
			"grant_type": "client_credential",
			"appid":      w.appId,
			"secret":     w.secret,
		},
		Method: curlx.METHOD_GET,
	}
	if err := h.Do().ParseJSON(&local); err != nil {
		return "", err
	}
	if local.ErrCode != 0 {
		return "", errors.New(local.ErrMsg)
	}

	local.ExpireAt = time.Now().Add(time.Second * time.Duration(local.ExpiresIn)).Unix()
	file, err := utilx.CreateFile(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	tb, err := json.Marshal(local)
	if err != nil {
		return "", err
	}

	if _, err = file.Write(tb); err != nil {
		return "", err
	}
	return local.AccessToken, nil
}

func (w *WeChat) GetAppId() string {
	return w.appId
}

func (w *WeChat) GetAppSecret() string {
	return w.secret
}

func (w *WeChat) Exec(e Engine) {
	e.Sent(w)
	e.Exec()
}
