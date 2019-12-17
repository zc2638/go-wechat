package wechat

import (
	"encoding/json"
	"errors"
	"github.com/zc2638/gotool/curlx"
	"github.com/zc2638/gotool/utilx"
	"io/ioutil"
	"time"
)

/**
 * Created by zc on 2019/12/9.
 */
type WeChat struct {
	appId           string
	secret          string
	pemCertPath     string // 证书路径
	pemKeyPath      string // 证书密钥路径
	host            string
	accessTokenPath string
	storage         Storage
	mch             *Merchant
}

func NewWeChat(appId, secret string, options ...Option) *WeChat {
	w := &WeChat{
		appId: appId,
		secret: secret,
		mch: &Merchant{},
	}
	for _, option := range options {
		option(w)
	}
	if w.host == "" {
		w.host = DOMAIN
	}
	return w
}

// 获取公众号全局唯一后台接口调用凭据（access_token）。
func (w *WeChat) BuildAccessToken() (string, error) {

	filePath := AccessTokenPath
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
		Url: w.host + "/cgi-bin/token",
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

func (w *WeChat) GetHost() string {
	return w.host
}

func (w *WeChat) GetMerchant() *Merchant {
	return w.mch
}

func (w *WeChat) Exec(e Engine) error {
	return e.Exec(w)
}

type Option func(w *WeChat)

func NewMerchant(appId, mchId, apiKey string) Option {
	return func(w *WeChat) {
		w.mch.AppId = appId
		w.mch.MchId = mchId
		w.mch.ApiKey = apiKey
	}
}

func NewPemCert(certPath, keyPath string, InsecureSkipVerify bool) (Option, error) {
	tr, err := curlx.NewTransport(certPath, keyPath, InsecureSkipVerify)
	if err != nil {
		return nil, err
	}
	return func(w *WeChat) {
		w.pemCertPath = certPath
		w.pemKeyPath = keyPath
		w.mch.Transport = tr
	}, nil
}

func NewDomain(host string) Option {
	return func(w *WeChat) {
		w.host = host
	}
}

func NewTokenPath(path string) Option {
	return func(w *WeChat) {
		w.accessTokenPath = path
	}
}

func NewTokenStorage(storage Storage) Option {
	return func(w *WeChat) {
		w.storage = storage
	}
}
