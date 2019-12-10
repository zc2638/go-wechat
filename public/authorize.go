package public

import (
	"encoding/json"
	"github.com/zc2638/wechat"
	"github.com/zc2638/wechat/config"
	"github.com/zctod/go-tool/common/curlx"
	"net/url"
)

/**
 * Created by zc on 2019/12/9.
 */
// 授权链接，获取code
type AuthorizeUrl struct {
	appId       string
	RedirectUri string // 返回跳转链接
	Scope       string // 授权类型：静默snsapi_base, 正常snsapi_userinfo
	State       string // 额外值
	Result      string // 返回
}

func (a *AuthorizeUrl) Sent(drive wechat.Drive) {
	a.appId = drive.GetAppId()
}

func (a *AuthorizeUrl) Exec() {
	if a.Scope == "" {
		a.Scope = config.SCOPE_SNSAPI_BASE
	}
	a.Result = config.PUBLIC_AUTHORIZEURL + "?appid=" + a.appId + "&redirect_uri=" + url.QueryEscape(a.RedirectUri) + "&response_type=code&scope=" + a.Scope + "&state=" + a.State + "#wechat_redirect"
}

// 根据code获取用户身份信息
type AuthorizeInfo struct {
	appId  string
	secret string
	Code   string
	Err    error
	Result AuthorizeInfoResult
}

type AuthorizeInfoResult struct {
	AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn    int    `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` // 用户刷新access_token
	Openid       string `json:"openid"`        // 用户唯一标识，请注意，在未关注公众号时，用户访问公众号的网页，也会产生一个用户和公众号唯一的OpenID
	Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
	wechat.ResCode
}

func (a *AuthorizeInfo) Sent(drive wechat.Drive) {
	a.appId = drive.GetAppId()
	a.secret = drive.GetAppSecret()
}

func (a *AuthorizeInfo) Exec() {

	h := curlx.HttpReq{
		Url: config.PUBLIC_AUTHORIZEINFO,
		Query: map[string]string{
			"appid":      a.appId,
			"secret":     a.secret,
			"code":       a.Code,
			"grant_type": "authorization_code",
		},
	}
	b, err := h.Get()
	if err != nil {
		a.Err = err
		return
	}

	var res AuthorizeInfoResult
	if err := json.Unmarshal(b, &res); err != nil {
		a.Err = err
		return
	}
	a.Result = res
}

// 刷新用户个人的access_token
type RefreshToken struct {
	appId        string
	RefreshToken string
	Err          error
	Result       AuthorizeInfoResult
}

func (a *RefreshToken) Sent(drive wechat.Drive) {
	a.appId = drive.GetAppId()
}

func (a *RefreshToken) Exec() {

	h := curlx.HttpReq{
		Url: config.PUBLIC_REFRESHTOKEN,
		Query: map[string]string{
			"appid":         a.appId,
			"grant_type":    "refresh_token",
			"refresh_token": a.RefreshToken,
		},
	}

	b, err := h.Get()
	if err != nil {
		a.Err = err
		return
	}

	var res AuthorizeInfoResult
	if err := json.Unmarshal(b, &res); err != nil {
		a.Err = err
		return
	}
	a.Result = res
}

// 拉取用户个人信息
type UserInfo struct {
	AccessToken string // 用户token
	Openid      string
	Err         error
	Result      UserInfoResult
}

type UserInfoResult struct {
	Openid     string   `json:"openid"`     // 用户的唯一标识
	Nickname   string   `json:"nickname"`   // 用户昵称
	Sex        string   `json:"sex"`        // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Province   string   `json:"province"`   // 用户个人资料填写的省份
	City       string   `json:"city"`       // 普通用户个人资料填写的城市
	Country    string   `json:"country"`    // 国家，如中国为CN
	HeadImgUrl string   `json:"headimgurl"` // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	Privilege  []string `json:"privilege"`  // 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	UnionId    string   `json:"unionid"`    // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	wechat.ResCode
}

func (a *UserInfo) Sent(drive wechat.Drive) {}

func (a *UserInfo) Exec() {

	h := curlx.HttpReq{
		Url: config.PUBLIC_USERINFO,
		Query: map[string]string{
			"access_token": a.AccessToken,
			"openid":       a.Openid,
			"lang":         "zh_CN", // 返回国家地区语言版本，zh_CN 简体，zh_TW 繁体，en 英语
		},
	}

	b, err := h.Get()
	if err != nil {
		a.Err = err
		return
	}

	var res UserInfoResult
	if err := json.Unmarshal(b, &res); err != nil {
		a.Err = err
		return
	}
	a.Result = res
}

// 检验授权凭证（access_token）是否有效
type CheckAccessToken struct {
	AccessToken string
	Openid      string
	Err         error
	Result      wechat.ResCode
}

func (a *CheckAccessToken) Sent(drive wechat.Drive) {}

func (a *CheckAccessToken) Exec() {

	h := curlx.HttpReq{
		Url: config.PUBLIC_CHECKTOKEN,
		Query: map[string]string{
			"access_token": a.AccessToken,
			"openid":       a.Openid,
		},
	}

	b, err := h.Get()
	if err != nil {
		a.Err = err
		return
	}

	var res wechat.ResCode
	if err := json.Unmarshal(b, &res); err != nil {
		a.Err = err
		return
	}
	a.Result = res
}

