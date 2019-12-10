package xcx

import (
	"encoding/json"
	"github.com/zc2638/wechat"
	"github.com/zc2638/wechat/config"
	"github.com/zctod/go-tool/common/curlx"
	"strconv"
)

/**
 * Created by zc on 2019/12/10.
 */
// 获取小程序二维码（数量限制）
type QRCode struct {
	accessToken string
	Path        string // 扫码进入的小程序页面路径，最大长度 128 字节，不能为空；对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
	Width       int    // 二维码的宽度，单位 px。最小 280px，最大 1280px
	Err         error
	Result      CodeResult
}

type CodeResult struct {
	ContentType string `json:"contentType"`
	Buffer      []byte `json:"buffer"`
	wechat.ResCode
}

func (c *QRCode) Sent(drive wechat.Drive) {
	c.accessToken, c.Err = drive.BuildAccessToken()
}

func (c *QRCode) Exec() {

	h := curlx.HttpReq{
		Url: config.XCX_QRCODE,
		Query: map[string]string{
			"access_token": c.accessToken,
		},
		Params: map[string]string{
			"path":  c.Path,
			"width": strconv.Itoa(c.Width),
		},
		Header: map[string]string{
			"Content-Type": "application/json; encoding=utf-8",
		},
	}

	b, err := h.Post()
	if err != nil {
		c.Err = err
		return
	}

	var res CodeResult
	if err := json.Unmarshal(b, &res); err != nil {
		c.Err = err
		return
	}
	c.Result = res
}

// 获取小程序码（数量限制）
type Code struct {
	accessToken string     `json:"-"`
	Path        string     `json:"path"`       // 扫码进入的小程序页面路径，最大长度 128 字节，不能为空；对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
	Width       int        `json:"width"`      // 二维码的宽度，单位 px。最小 280px，最大 1280px
	AutoColor   bool       `json:"auto_color"` // 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
	LineColor   wechat.RGB `json:"line_color"` // auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"} 十进制表示
	IsHyaLine   bool       `json:"is_hyaline"` // 是否需要透明底色，为 true 时，生成透明底色的小程序码
	Err         error      `json:"-"`
	Result      CodeResult `json:"-"`
}

func (c *Code) Sent(drive wechat.Drive) {
	c.accessToken, c.Err = drive.BuildAccessToken()
}

func (c *Code) Exec() {

	cb, err := json.Marshal(c)
	if err != nil {
		c.Err = err
		return
	}

	h := curlx.HttpReq{
		Url: config.XCX_CODE,
		Query: map[string]string{
			"access_token": c.accessToken,
		},
		Body: cb,
		Header: map[string]string{
			"Content-Type": "application/json; encoding=utf-8",
		},
	}
	b, err := h.Post()
	if err != nil {
		c.Err = err
		return
	}

	var res CodeResult
	if err := json.Unmarshal(b, &res); err != nil {
		c.Err = err
		return
	}
	c.Result = res
}

// 获取小程序码（无限制）
type CodeUnlimited struct {
	accessToken string     `json:"-"`
	Scene       string     `json:"scene"`      // 最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）
	Page        string     `json:"page"`       // 必须是已经发布的小程序存在的页面（否则报错），例如 pages/index/index, 根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
	Width       int        `json:"width"`      // 二维码的宽度，单位 px。最小 280px，最大 1280px
	AutoColor   bool       `json:"auto_color"` // 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
	LineColor   wechat.RGB `json:"line_color"` // auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"} 十进制表示
	IsHyaLine   bool       `json:"is_hyaline"` // 是否需要透明底色，为 true 时，生成透明底色的小程序码
	Err         error      `json:"-"`
	Result      CodeResult `json:"-"`
}

func (c *CodeUnlimited) Sent(drive wechat.Drive) {
	c.accessToken, c.Err = drive.BuildAccessToken()
}

func (c *CodeUnlimited) Exec() {

	cb, err := json.Marshal(c)
	if err != nil {
		c.Err = err
		return
	}

	h := curlx.HttpReq{
		Url: config.XCX_CODE_UNLIMITED,
		Query: map[string]string{
			"access_token": c.accessToken,
		},
		Body: cb,
		Header: map[string]string{
			"Content-Type": "application/json; encoding=utf-8",
		},
	}
	b, err := h.Post()
	if err != nil {
		c.Err = err
		return
	}

	var res CodeResult
	if err := json.Unmarshal(b, &res); err != nil {
		c.Err = err
		return
	}
	c.Result = res
}
