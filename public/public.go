package public

import (
	"encoding/json"
	"github.com/zctod/tool/common/utils"
	"go-wechat/core"
	"go-wechat/wechat"
	"io/ioutil"
	"net/url"
	"time"
)

type WechatPublic struct {
	Appid           string // 微信分配的小程序ID
	Appsecret       string // 微信分配的小程序密钥
	AccessTokenPath string // access_token文件存储路径
	accessToken     string // 微信小程序唯一凭证
}

// 获取公众号全局唯一后台接口调用凭据（access_token）。
func (w *WechatPublic) BuildAccessToken() error {

	var filePath string
	if w.AccessTokenPath == "" {
		filePath = core.WECHAT_PUBLIC_ACCESSTOKEN_PATH
	} else {
		filePath = w.AccessTokenPath
	}
	var token wechat.LocalAccessToken
	b, err := ioutil.ReadFile(filePath)
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

	var res wechat.ResAccessToken
	h := core.HttpReq{Url: core.WECHAT_ACCESSTOKEN + "?grant_type=client_credential&appid=" + w.Appid + "&secret=" + w.Appsecret}
	err = h.Get(&res)
	if err == nil {
		w.accessToken = res.AccessToken
		token.AccessToken = res.AccessToken
		token.ExpireAt = int(time.Now().Add(time.Hour * 2).Unix())
		file, err := utils.CreateFile(filePath)
		defer file.Close()
		if err == nil {
			tb, _ := json.Marshal(token)
			_, _ = file.Write(tb)
		}
		return nil
	}
	return err
}

// 授权链接，获取code
func (w *WechatPublic) AuthorizeUrl(a AuthorizeUrl) string {
	if a.Scope == "" {
		a.Scope = core.SCOPE_SNSAPI_BASE
	}
	return core.PUBLIC_AUTHORIZEURL + "?appid=" + w.Appid + "&redirect_uri=" + url.QueryEscape(a.RedirectUri) + "&response_type=code&scope=" + a.Scope + "&state=" + a.State + "#wechat_redirect"
}

// 获取用户身份信息
func (w *WechatPublic) AuthorizeInfo(code string) (core.M, error) {

	var h = core.HttpReq{
		Url: core.PUBLIC_AUTHORIZEINFO + "?appid=" + w.Appid + "&secret=" + w.Appsecret + "&code=" + code + "&grant_type=authorization_code",
	}
	return h.GetData()
}

// 刷新用户个人的access_token
func (w *WechatPublic) RefreshToken(token string) (core.M, error) {

	var h = core.HttpReq{
		Url: core.PUBLIC_REFRESHTOKEN + "?appid=" + w.Appid + "&grant_type=refresh_token&refresh_token=" + token,
	}
	return h.GetData()
}

// 拉取用户个人信息
func (w *WechatPublic) UserInfo(token, openid string) (core.M, error) {

	var h = core.HttpReq{
		Url: core.PUBLIC_USERINFO + "?access_token=" + token + "&openid=" + openid + "&lang=zh_CN",
	}
	return h.GetData()
}

// 检验授权凭证（access_token）是否有效
func (w *WechatPublic) CheckToken(token, openid string) (core.M, error) {

	var h = core.HttpReq{
		Url: core.PUBLIC_CHECKTOKEN + "?access_token=" + token + "&openid=" + openid,
	}
	return h.GetData()
}

// 自定义菜单创建
func (w *WechatPublic) MenuCreate(menu Menu) (core.M, error) {

	b, err := json.Marshal(menu)
	if err != nil {
		return nil, err
	}
	if err := w.BuildAccessToken(); err != nil {
		return nil, err
	}
	var h = core.HttpReq{
		Url: core.PUBLIC_MENUCREATE + "?access_token=" + w.accessToken,
		Body: string(b),
	}
	return h.JsonData()
}

// 自定义菜单查询
func (w *WechatPublic) MenuGet() (string, error) {

	if err := w.BuildAccessToken(); err != nil {
		return "", err
	}
	var h = core.HttpReq{
		Url: core.PUBLIC_MENUGET + "?access_token=" + w.accessToken,
	}
	return h.JsonStr()
}

// 自定义菜单删除
func (w *WechatPublic) MenuDelete() (core.M, error) {

	if err := w.BuildAccessToken(); err != nil {
		return nil, err
	}
	var h = core.HttpReq{
		Url: core.PUBLIC_MENUDELETE + "?access_token=" + w.accessToken,
	}
	return h.GetData()
}

// 发送模板消息
func (w *WechatPublic) TemplateSend(t Template) (core.M, error) {

	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	if err := w.BuildAccessToken(); err != nil {
		return nil, err
	}
	var h = core.HttpReq{
		Url: core.PUBLIC_TEMPLATE_SEND + "?access_token=" + w.accessToken,
		Body: string(b),
	}
	return h.JsonData()
}