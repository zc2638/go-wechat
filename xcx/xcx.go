package xcx

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"github.com/zctod/tool/common/utils"
	"go-wechat/core"
	"io/ioutil"
	"strconv"
	"time"
)

type WechatXcx struct {
	Appid       string // 微信分配的小程序ID
	Appsecret   string // 微信分配的小程序密钥
	MchId       string // 微信支付分配的商户号
	ApiKey		string // 微信支付商户密钥
	accessToken string // 微信小程序唯一凭证
}

// 登录凭证校验
func (w *WechatXcx) Code2session(code string) (res ResCode2Session, err error) {

	h := core.HttpReq{Url: core.WECHAT_XCX_CODE2SESSION + "?appid=" + w.Appid + "&secret=" + w.Appsecret + "&js_code=" + code + "&grant_type=authorization_code"}
	err = h.Get(&res)
	return
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
func (w *WechatXcx) SendTemplate(tp Template) (res ResCode, err error) {

	err = w.BuildAccessToken()
	if err != nil {
		return
	}
	params, _ := utils.StrcutToMap(tp)
	h := core.HttpReq{
		Url:    core.WECHAT_XCX_TEMPLATE_SEND + "?access_token=" + w.accessToken,
		Params: params,
	}
	err = h.Post(&res)
	return
}

// 统一下单
func (w *WechatXcx) OrderPay(order UnifiedOrder) (res ResUnifiedOrder, err error) {

	if order.NotifyUrl == "" {
		err = errors.New("notify_url cannot be empty")
		return
	}
	if order.Openid == "" {
		err = errors.New("openid cannot be empty")
		return
	}
	if order.TotalFee == 0 {
		err = errors.New("total_fee must greater than 0")
		return
	}

	switch order.SignType {
	case core.SIGNTYPE_HMAC_SHA256:
		order.SignType = core.SIGNTYPE_HMAC_SHA256
		break
	default:
		order.SignType = core.SIGNTYPE_MD5
		break
	}
	if order.SpbillCreateIp == "" {
		order.SpbillCreateIp = "127.0.0.1"
	}

	var orderReq = UnifiedOrderReq{
		Appid: w.Appid,
		MchId: w.MchId,
		Body: order.Body,
		OutTradeNo: order.OutTradeNo,
		SpbillCreateIp: order.SpbillCreateIp,
		NotifyUrl: order.NotifyUrl,
		TradeType: "JSAPI",
		TotalFee: strconv.Itoa(order.TotalFee),
		NonceStr: utils.RandomStr(32),
		DeviceInfo: order.DeviceInfo,
		SignType: order.SignType,
		Detail: order.Detail,
		Attach: order.Attach,
		TimeStart: order.TimeStart,
		TimeExpire: order.TimeExpire,
		GoodsTag: order.GoodsTag,
		ProductId: order.ProductId,
		LimitPay: order.LimitPay,
		Openid: order.Openid,
		Receipt: order.Receipt,
	}

	// 补全签名
	var req = make(map[string]string)
	params, err := utils.StrcutToMap(orderReq)
	if err != nil {
		return
	}
	for k, v := range params {
		if v.(string) == "" {
			delete(params, k)
		}
		req[k] = v.(string)
	}
	switch order.SignType {
	case core.SIGNTYPE_HMAC_SHA256:
		orderReq.Sign = core.Sign(req, w.ApiKey, hmac.New(sha256.New, []byte(w.ApiKey)))
		break
	default:
		orderReq.Sign = core.Sign(req, w.ApiKey, md5.New())
		break
	}

	params["sign"] = orderReq.Sign
	h := core.HttpReq{
		Url: core.WECHAT_XCX_UNIFIEDORDER,
		Params: params,
	}
	mp, err := h.PostXml()
	if err != nil {
		return
	}
	err = utils.MapToStruct(mp, &res)
	return
}
