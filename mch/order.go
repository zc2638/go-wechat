package mch

import (
	"github.com/zc2638/wechat"
)

/**
 * Created by zc on 2019/12/17.
 */
// 统一下单
type OrderUnified struct {
	Body           string             `json:"body"`             // 必填。商品简单描述
	OutTradeNo     string             `json:"out_trade_no"`     // 必填。商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一。
	SpbillCreateIp string             `json:"spbill_create_ip"` // 必填。支持IPV4和IPV6两种格式的IP地址。调用微信支付API的机器IP（127.0.0.1）
	NotifyUrl      string             `json:"notify_url"`       // 必填。异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。
	TradeType      string             `json:"trade_type"`       // 必填。JSAPI--JSAPI支付（或小程序支付）、NATIVE--Native支付、APP--app支付，MWEB--H5支付
	TotalFee       int                `json:"total_fee"`        // 必填。订单总金额，单位为分
	Openid         string             `json:"openid"`           // 必填。用户唯一openid
	DeviceInfo     string             `json:"device_info"`      // 可选。自定义参数，可以为终端设备号(门店号或收银设备ID)，PC网页或公众号内支付可以传"WEB"
	SignType       string             `json:"sign_type"`        // 可选。签名类型，默认为MD5，支持HMAC-SHA256和MD5。
	Detail         string             `json:"detail"`           // 可选。商品详细描述
	Attach         string             `json:"attach"`           // 可选。附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用。
	TimeStart      string             `json:"time_start"`       // 可选。订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。
	TimeExpire     string             `json:"time_expire"`      // 可选。订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。订单失效时间是针对订单号而言的，由于在请求支付的时候有一个必传参数prepay_id只有两小时的有效期，所以在重入时间超过2小时的时候需要重新请求下单接口获取新的prepay_id。
	GoodsTag       string             `json:"goods_tag"`        // 可选。订单优惠标记，使用代金券或立减优惠功能时需要的参数
	ProductId      string             `json:"product_id"`       // 可选。trade_type=NATIVE时，此参数必传。此参数为二维码中包含的商品ID，商户自行定义。
	LimitPay       string             `json:"limit_pay"`        // 可选。上传此参数no_credit--可限制用户不能使用信用卡支付
	Receipt        string             `json:"receipt"`          // 可选。Y，传入Y时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效
	Result         OrderUnifiedResult `json:"-"`
}

type OrderUnifiedResult struct {
	wechat.Return
	TradeType string `json:"trade_type"` // 交易类型：JSAPI -JSAPI支付，NATIVE -Native支付，APP -APP支付
	PrepayId  string `json:"prepay_id"`  // 预支付交易会话标识。微信生成的预支付会话标识，用于后续接口调用中使用，该值有效期为2小时
	CodeUrl   string `json:"code_url"`   // 二维码链接。trade_type=NATIVE时有返回，此url用于生成支付二维码，然后提供给用户进行扫码支付。注意：code_url的值并非固定，使用时按照URL格式转成二维码即可
}

func (o *OrderUnified) Exec(drive wechat.Drive) error {
	if o.SpbillCreateIp == "" {
		o.SpbillCreateIp = "127.0.0.1"
	}
	if o.TradeType == "" {
		o.TradeType = TRADETYPE_JSAPI
	}
	return drive.GetMerchant().Exec("/pay/unifiedorder", o, &o.Result, false)
}

// 查询订单
type OrderQuery struct {
	TransactionId string           `json:"transaction_id"` // 与商户订单号二选一。微信的订单号，优先使用
	OutTradeNo    string           `json:"out_trade_no"`   // 与微信订单号二选一。商户系统内部订单号
	SignType      string           `json:"sign_type"`      // 可选。签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	Result        OrderQueryResult `json:"-"`
}

type OrderQueryResult struct {
	wechat.Return
	DeviceInfo         string `json:"device_info"`          // 设备号
	Openid             string `json:"openid"`               // 用户标识
	IsSubscribe        string `json:"is_subscribe"`         // 是否关注公众账号。Y-关注，N-未关注
	TradeType          string `json:"trade_type"`           // 交易类型。JSAPI，NATIVE，APP，MICROPAY
	TradeState         string `json:"trade_state"`          // 交易状态。SUCCESS—支付成功，REFUND—转入退款，NOTPAY—未支付，CLOSED—已关闭，REVOKED—已撤销（付款码支付），USERPAYING--用户支付中（付款码支付），PAYERROR--支付失败(其他原因，如银行返回失败)
	BankType           string `json:"bank_type"`            // 银行类型，采用字符串类型的银行标识
	TotalFee           int    `json:"total_fee"`            // 订单总金额，单位为分
	SettlementTotalFee int    `json:"settlement_total_fee"` // 当订单使用了免充值型优惠券后返回该参数，应结订单金额=订单金额-免充值优惠券金额。
	FeeType            string `json:"fee_type"`             // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY
	CashFee            int    `json:"cash_fee"`             // 现金支付金额，订单现金支付金额
	CashFeeType        string `json:"cash_fee_type"`        // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY
	CouponFee          int    `json:"coupon_fee"`           // 代金券金额
	CouponCount        int    `json:"coupon_count"`         // 代金券使用数量
	TransactionId      string `json:"transaction_id"`       // 微信支付订单号
	OutTradeNo         string `json:"out_trade_no"`         // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一
	Attach             string `json:"attach"`               // 附加数据
	TimeEnd            string `json:"time_end"`             // 订单支付时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010
	TradeStateDesc     string `json:"trade_state_desc"`     // 交易状态描述
}

func (o *OrderQuery) Exec(drive wechat.Drive) error {
	return drive.GetMerchant().Exec("/pay/orderquery", o, &o.Result, false)
}

// 关闭订单
type OrderClose struct {
	OutTradeNo string        `json:"out_trade_no"` // 必填。商户系统内部订单号
	SignType   string        `json:"sign_type"`    // 可选。签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	Result     wechat.Return `json:"-"`
}

func (o *OrderClose) Exec(drive wechat.Drive) error {
	return drive.GetMerchant().Exec("/pay/closeorder", o, &o.Result, false)
}

// 申请退款
type OrderRefund struct {
	TransactionId string            `json:"transaction_id"` // 与商户订单号二选一。微信的订单号，优先使用
	OutTradeNo    string            `json:"out_trade_no"`   // 与微信订单号二选一。商户系统内部订单号
	SignType      string            `json:"sign_type"`      // 可选。签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	OutRefundNo   string            `json:"out_refund_no"`  // 商户系统内部的退款单号
	TotalFee      int               `json:"total_fee"`      // 订单总金额
	RefundFee     int               `json:"refund_fee"`     // 退款总金额
	RefundDesc    string            `json:"refund_desc"`    // 退款原因
	NotifyUrl     string            `json:"notify_url"`     // 异步接收微信支付退款结果通知的回调地址，通知URL必须为外网可访问的url，不允许带参数。如果参数中传了notify_url，则商户平台上配置的回调地址将不会生效。
	Result        OrderRefundResult `json:"-"`
}

type OrderRefundResult struct {
	wechat.Return
	TransactionId       string `json:"transaction_id"`        // 微信支付订单号
	OutTradeNo          string `json:"out_trade_no"`          // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一
	OutRefundNo         string `json:"out_refund_no"`         // 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	RefundId            string `json:"refund_id"`             // 微信退款单号
	RefundFee           int    `json:"refund_fee"`            // 退款总金额,单位为分,可以做部分退款
	SettlementRefundFee int    `json:"settlement_refund_fee"` // 应结退款金额。去掉非充值代金券退款金额后的退款金额，退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
	TotalFee            int    `json:"total_fee"`             // 订单总金额，单位为分，只能为整数
	SettlementTotalFee  int    `json:"settlement_total_fee"`  // 应结订单金额	
	FeeType             string `json:"fee_type"`              // 订单金额货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY
	CashFee             int    `json:"cash_fee"`              // 现金支付金额，订单现金支付金额
	CashFeeType         string `json:"cash_fee_type"`         // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY
	CashRefundFee       int    `json:"cash_refund_fee"`       // 现金退款金额，单位为分，只能为整数
}

func (o *OrderRefund) Exec(drive wechat.Drive) error {
	return drive.GetMerchant().Exec("/secapi/pay/refund", o, &o.Result, true)
}
