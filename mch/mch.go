package mch

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"github.com/zc2638/wechat/config"
	"github.com/zc2638/wechat/core"
	"github.com/zctod/go-tool/common/utils"
	"strconv"
)

type Merchant struct {
	Appid       string // 微信分配的小程序ID
	MchId       string // 微信支付分配的商户号
	ApiKey      string // 微信支付商户密钥
	PemCertFile string // 证书路径
	PemKeyFile  string // 证书密钥路径
}

func (m *Merchant) execOrder(data interface{}, url string, isCert bool) (core.M, error) {

	// 补全签名
	params, err := utils.StrcutToMap(data)
	if err != nil {
		return nil, err
	}
	params["nonce_str"] = utils.RandomStr(32)
	req := utils.MapToStringMap(params)

	var signType string
	if _, ok := params["sign_type"]; !ok {
		signType = config.SIGNTYPE_MD5
	} else {
		signType = params["sign_type"].(string)
	}

	switch signType {
	case config.SIGNTYPE_HMAC_SHA256:
		params["sign"] = core.Sign(req, m.ApiKey, hmac.New(sha256.New, []byte(m.ApiKey)))
		break
	case config.SIGNTYPE_SHA1:
		params["sign"] = core.Sign(req, m.ApiKey, hmac.New(sha1.New, []byte(m.ApiKey)))
		break
	default:
		params["sign"] = core.Sign(req, m.ApiKey, md5.New())
		break
	}

	h := core.HttpReq{
		Url:    url,
		Params: params,
	}
	var mp map[string]string
	if isCert == true {
		h.CertFile = m.PemCertFile
		h.KeyFile = m.PemKeyFile
	}

	mp, err = h.XmlData()
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
func (m *Merchant) UnifiedOrder(o UnifiedOrder) (core.M, error) {

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
	if o.TradeType == "" {
		o.TradeType = TRADETYPE_JSAPI
	}

	var orderReq = reqUnifiedOrder{
		Appid:          m.Appid,
		MchId:          m.MchId,
		Body:           o.Body,
		OutTradeNo:     o.OutTradeNo,
		SpbillCreateIp: o.SpbillCreateIp,
		NotifyUrl:      o.NotifyUrl,
		TradeType:      o.TradeType,
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

	return m.execOrder(orderReq, config.MCH_UNIFIEDORDER, false)
}

// 查询订单
func (m *Merchant) OrderQuery(o OrderQuery) (core.M, error) {

	var orderReq = reqOrderQuery{
		Appid:         m.Appid,
		MchId:         m.MchId,
		TransactionId: o.TransactionId,
		OutTradeNo:    o.OutTradeNo,
	}
	return m.execOrder(orderReq, config.MCH_QUERYORDER, false)
}

// 关闭订单
func (m *Merchant) OrderClose(o OrderClose) (core.M, error) {

	var orderReq = reqOrderClose{
		Appid:      m.Appid,
		MchId:      m.MchId,
		OutTradeNo: o.OutTradeNo,
	}
	return m.execOrder(orderReq, config.MCH_CLOSEORDER, false)
}

// 订单退款
func (m *Merchant) OrderRefund(o OrderRefund) (core.M, error) {

	var orderReq = reqOrderRefund{
		Appid:         m.Appid,
		MchId:         m.MchId,
		TransactionId: o.TransactionId,
		OutTradeNo:    o.OutTradeNo,
		OutRefundNo:   o.OutRefundNo,
		TotalFee:      o.TotalFee,
		RefundFee:     o.RefundFee,
		RefundDesc:    o.RefundDesc,
		NotifyUrl:     o.NotifyUrl,
	}
	return m.execOrder(orderReq, config.MCH_REFUNDORDER, true)
}

// 企业付款到零钱
func (m *Merchant) Transfer(t Transfer) (core.M, error) {

	var err error
	if t.Openid == "" {
		err = errors.New("请填写用户openid")
		return nil, err
	}
	if t.Amount <= 0 {
		err = errors.New("请填写大于0的金额")
		return nil, err
	}
	if t.CheckName == "" {
		t.CheckName = CHECK_NAME_FALSE
	}
	if t.CheckName == CHECK_NAME_TRUE && t.ReUserName == "" {
		err = errors.New("选择强制校验真实姓名，必须填写真实姓名")
		return nil, err
	}
	if t.SpbillCreateIp == "" {
		t.SpbillCreateIp = "127.0.0.1"
	}
	var transferReq = reqTransfer{
		MchAppid:       m.Appid,
		MchId:          m.MchId,
		DeviceInfo:     t.DeviceInfo,
		PartnerTradeNo: t.PartnerTradeNo,
		Openid:         t.Openid,
		CheckName:      t.CheckName,
		ReUserName:     t.ReUserName,
		Amount:         t.Amount,
		Desc:           t.Desc,
		SpbillCreateIp: t.SpbillCreateIp,
	}
	return m.execOrder(transferReq, config.MCH_TRANSFERS, true)
}

// 企业付款到零钱查询
func (m *Merchant) TransferGet(partnerTradeNo string) (core.M, error) {

	var transferReq = reqTransferGet{
		Appid:          m.Appid,
		MchId:          m.MchId,
		PartnerTradeNo: partnerTradeNo,
	}
	return m.execOrder(transferReq, config.MCH_TRANSFERSGET, true)
}
