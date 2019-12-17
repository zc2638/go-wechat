package wechat

import (
	"bufio"
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/zc2638/gotool/curlx"
	"github.com/zc2638/gotool/utilx"
	"net/http"
	"sort"
)

/**
 * Created by zc on 2019/12/9.
 */
type Engine interface {
	Exec(drive Drive) error
}

type Drive interface {
	BuildAccessToken() (string, error)
	GetAppId() string
	GetAppSecret() string
	GetHost() string
	GetMerchant() *Merchant
}

type Merchant struct {
	AppId     string
	MchId     string // 微信支付分配的商户号
	ApiKey    string // 微信支付商户密钥
	Transport *http.Transport
}

func (m *Merchant) GetHost() string {
	return DOMAIN_MCH
}

func (m *Merchant) sign(data interface{}) (map[string]string, error) {
	params, err := m.buildParams(data)
	if err != nil {
		return nil, err
	}

	var h = md5.New()
	switch params["sign_type"] {
	case SIGNTYPE_HMAC_SHA256:
		h = hmac.New(sha256.New, []byte(m.ApiKey))
	case SIGNTYPE_SHA1:
		h = hmac.New(sha1.New, []byte(m.ApiKey))
	}

	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	bufw := bufio.NewWriterSize(h, 128)
	for _, k := range keys {
		v := params[k]
		if v == "" {
			continue
		}
		_, _ = bufw.WriteString(k + "=" + v + "&")
	}
	_, _ = bufw.WriteString("key=" + m.ApiKey)
	_ = bufw.Flush()

	signature := make([]byte, hex.EncodedLen(h.Size()))
	hex.Encode(signature, h.Sum(nil))
	params["sign"] = string(bytes.ToUpper(signature))
	return params, nil
}

func (m *Merchant) buildParams(data interface{}) (map[string]string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var order map[string]interface{}
	if err := json.Unmarshal(b, &order); err != nil {
		return nil, err
	}
	params, err := utilx.MapToStringMap(order)
	if err != nil {
		return nil, err
	}
	params["appid"] = m.AppId
	params["mch_id"] = m.MchId
	params["nonce_str"] = utilx.RandomStr(32)
	return params, nil
}

func (m *Merchant) Exec(path string, data interface{}, res interface{}, isCert bool) error {
	params, err := m.sign(data)
	if err != nil {
		return err
	}
	h := curlx.HttpReq{
		Url:  m.GetHost() + path,
		Body: utilx.MapToXml(params),
		Header: map[string]string{
			curlx.HEADER_CONTENT_TYPE: curlx.CT_TEXT_XML,
		},
		Method: curlx.METHOD_POST,
	}
	if isCert {
		h.SetTransport(m.Transport)
	}
	return h.Do().ParseXML(res)
}

type Storage uint

const (
	DRIVER_DEFAULT Storage = iota
	DRIVER_FILE
	DRIVER_REDIS
)

// access_token文件路径
const AccessTokenPath = "access_token.txt"

const (
	// 微信业务接口域名
	DOMAIN = "https://api.weixin.qq.com"
	// 微信业务接口域名（通用异地容灾）
	DOMAIN2 = "https://api2.weixin.qq.com"
	// 微信商户接口域名
	DOMAIN_MCH = "https://api.mch.weixin.qq.com"
	// 微信商户接口域名（通用异地容灾）
	DOMAIN_MCH2 = "https://api2.mch.weixin.qq.com"
)

const (
	SIGNTYPE_MD5         = "MD5"
	SIGNTYPE_SHA1        = "SHA1"
	SIGNTYPE_HMAC_SHA256 = "HMAC-SHA256"
)

type LocalAccess struct {
	AccessToken string `json:"access_token"` // 凭证
	ExpireAt    int64  `json:"expire_at"`    // 过期unix时间戳
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值。
	ResCode
}

type ResCode struct {
	ErrCode int    `json:"errcode"` // 错误码
	ErrMsg  string `json:"errmsg"`  // 错误信息
}

type Return struct {
	ReturnCode string `json:"return_code"`             // 返回状态码。SUCCESS/FAIL 此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	ReturnMsg  string `json:"return_msg"`              // 返回信息。当return_code为FAIL时返回信息为错误原因
	ResultCode string `json:"result_code"`             // 业务结果。SUCCESS/FAIL
	ErrCode    string `json:"err_code, omitempty"`     // 错误代码。当result_code为FAIL时返回错误代码
	ErrCodeDes string `json:"err_code_des, omitempty"` // 错误代码描述。当result_code为FAIL时返回错误代码
	AppId      string `json:"appid"`                   // 公众账号ID
	MchId      string `json:"mch_id"`                  // 商户号
	NonceStr   string `json:"nonce_str"`               // 随机字符串
	Sign       string `json:"sign"`                    // 签名
}

type RGB struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}
