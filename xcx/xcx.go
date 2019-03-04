package xcx

import (
	"encoding/json"
	"github.com/zctod/tool/common/utils"
	"go-wechat/core"
	"io/ioutil"
	"time"
)

type WechatXcx struct {
	Appid       string // 微信分配的小程序ID
	Appsecret   string // 微信分配的小程序密钥
	accessToken string // 微信小程序唯一凭证
}

// 登录凭证校验
func (w *WechatXcx) Code2session(code string) (core.M, error) {

	h := core.HttpReq{Url: core.WECHAT_XCX_CODE2SESSION + "?appid=" + w.Appid + "&secret=" + w.Appsecret + "&js_code=" + code + "&grant_type=authorization_code"}
	return h.GetData()
}

// 获取小程序全局唯一后台接口调用凭据（access_token）。
func (w *WechatXcx) BuildAccessToken() error {

	var token ResLocalToken
	b, err := ioutil.ReadFile(core.WECHAT_XCX_ACCESSTOKEN_PATH)
	if err == nil {
		err = json.Unmarshal(b, &token)
		if err == nil {
			nowTs := time.Now().Unix()
			if token.ExpireAt > int(nowTs) {
				w.accessToken = token.AccessToken
				return nil
			}
		}
	}

	var res ResAccessToken
	h := core.HttpReq{Url: core.WECHAT_XCX_ACCESSTOKEN + "?grant_type=client_credential&appid=" + w.Appid + "&secret=" + w.Appsecret}
	err = h.Get(&res)
	if err == nil {
		w.accessToken = res.AccessToken
		token.AccessToken = res.AccessToken
		token.ExpireAt = int(time.Now().Add(time.Hour * 2).Unix())
		file, err := utils.CreateFile(core.WECHAT_XCX_ACCESSTOKEN_PATH)
		defer file.Close()
		if err == nil {
			tb, _ := json.Marshal(token)
			_, _ = file.Write(tb)
		}
		return nil
	}
	return err
}

// 检验数据的真实性，并且获取解密后的明文.
func (w *WechatXcx) DecryptData(encryptedData, sessionKey, iv string) (core.M, error) {

	b, err := core.AesDecrypt(encryptedData, sessionKey, iv)
	if err != nil {
		return nil, err
	}
	return utils.StrcutToMap(b)
}

// 发送模板消息
func (w *WechatXcx) SendTemplate(tp Template) (core.M, error) {

	err := w.BuildAccessToken()
	if err != nil {
		return nil, err
	}
	params, _ := utils.StrcutToMap(tp)
	h := core.HttpReq{
		Url:    core.WECHAT_XCX_TEMPLATE_SEND + "?access_token=" + w.accessToken,
		Params: params,
	}
	return h.PostData()
}