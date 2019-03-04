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
	ApiKey      string // 微信支付商户密钥
	accessToken string // 微信小程序唯一凭证
	PemCertFile string
	PemKeyFile  string
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

func (w *WechatXcx) execOrder(order interface{}, url, signType string, isCert bool) (core.M, error) {

	// 补全签名
	params, err := utils.StrcutToMap(order)
	if err != nil {
		return nil, err
	}
	params["nonce_str"] = utils.RandomStr(32)
	req := utils.MapToStringMap(params)

	switch signType {
	case core.SIGNTYPE_HMAC_SHA256:
		params["sign"] = core.Sign(req, w.ApiKey, hmac.New(sha256.New, []byte(w.ApiKey)))
		break
	default:
		params["sign"] = core.Sign(req, w.ApiKey, md5.New())
		break
	}

	h := core.HttpReq{
		Url:    url,
		Params: params,
	}
	var mp map[string]string
	if isCert == true {
		h.CertFile = w.PemCertFile
		h.KeyFile = w.PemKeyFile
		mp, err = h.XmlDataWithCert()
	} else {
		mp, err = h.XmlData()
	}
	if err != nil {
		return nil, err
	}

	var res = make(core.M)
	for k, v := range mp {
		res[k] = v
	}
	return res, nil
}

// 统一下单
func (w *WechatXcx) OrderPay(o UnifiedOrder) (core.M, error) {

	var err error
	if o.NotifyUrl == "" {
		err = errors.New("notify_url cannot be empty")
		return nil, err
	}
	if o.Openid == "" {
		err = errors.New("openid cannot be empty")
		return nil, err
	}
	if o.TotalFee == 0 {
		err = errors.New("total_fee must greater than 0")
		return nil, err
	}
	if o.SpbillCreateIp == "" {
		o.SpbillCreateIp = "127.0.0.1"
	}

	var orderReq = UnifiedOrderReq{
		Appid:          w.Appid,
		MchId:          w.MchId,
		Body:           o.Body,
		OutTradeNo:     o.OutTradeNo,
		SpbillCreateIp: o.SpbillCreateIp,
		NotifyUrl:      o.NotifyUrl,
		TradeType:      "JSAPI",
		TotalFee:       strconv.Itoa(o.TotalFee),
		DeviceInfo:     o.DeviceInfo,
		SignType:       o.SignType,
		Detail:         o.Detail,
		Attach:         o.Attach,
		TimeStart:      o.TimeStart,
		TimeExpire:     o.TimeExpire,
		GoodsTag:       o.GoodsTag,
		ProductId:      o.ProductId,
		LimitPay:       o.LimitPay,
		Openid:         o.Openid,
		Receipt:        o.Receipt,
	}

	return w.execOrder(orderReq, core.WECHAT_XCX_UNIFIEDORDER, o.SignType, false)
}

// 查询订单
func (w *WechatXcx) OrderQuery(o OrderQuery) (core.M, error) {

	var orderReq = OrderQueryReq{
		Appid:         w.Appid,
		MchId:         w.MchId,
		TransactionId: o.TransactionId,
		OutTradeNo:    o.OutTradeNo,
	}
	return w.execOrder(orderReq, core.WECHAT_XCX_QUERYORDER, o.SignType, false)
}

// 关闭订单
func (w *WechatXcx) OrderClose(o OrderClose) (core.M, error) {

	var orderReq = OrderCloseReq{
		Appid:      w.Appid,
		MchId:      w.MchId,
		OutTradeNo: o.OutTradeNo,
	}
	return w.execOrder(orderReq, core.WECHAT_XCX_CLOSEORDER, o.SignType, false)
}

// 订单退款
func (w *WechatXcx) OrderRefund(o OrderRefund) (core.M, error) {

	var orderReq = OrderRefundReq{
		Appid:         w.Appid,
		MchId:         w.MchId,
		TransactionId: o.TransactionId,
		OutTradeNo:    o.OutTradeNo,
		OutRefundNo:   o.OutRefundNo,
		TotalFee:      o.TotalFee,
		RefundFee:     o.RefundFee,
		RefundDesc:    o.RefundDesc,
		NotifyUrl:     o.NotifyUrl,
	}
	return w.execOrder(orderReq, core.WECHAT_XCX_REFUNDORDER, o.SignType, true)
}
