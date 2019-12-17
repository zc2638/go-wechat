package mch

const (
	TRADETYPE_JSAPI  = "JSAPI"  // JSAPI支付或小程序支付
	TRADETYPE_NATIVE = "NATIVE" // native支付
	TRADETYPE_APP    = "APP"    // app支付
	TRADETYPE_WEB    = "MWEB"   // H5支付
)
const (
	CHECK_NAME_FALSE = "NO_CHECK"    // 不校验真实姓名
	CHECK_NAME_TRUE  = "FORCE_CHECK" // 强制校验真实姓名
)

const (
	BANKCODE_ICBC = "1002" // 工商银行
	BANKCODE_ABC  = "1005" // 农业银行
	BANKCODE_BOC  = "1026" // 中国银行
	BANKCODE_CCB  = "1003" // 建设银行
	BANKCODE_CMB  = "1001" // 招商银行
	BANKCODE_PSBC = "1066" // 邮储银行
	BANKCODE_BCM  = "1020" // 交通银行
	BANKCODE_SPDB = "1004" // 浦发银行
	BANKCODE_CMSB = "1006" // 民生银行
	BANKCODE_CIB  = "1009" // 兴业银行
	BANKCODE_PAB  = "1010" // 平安银行
	BANKCODE_ZXB  = "1021" // 中信银行
	BANKCODE_HXB  = "1025" // 华夏银行
	BANKCODE_CGB  = "1027" // 广发银行
	BANKCODE_CEB  = "1022" // 光大银行
	BANKCODE_BOB  = "1032" // 北京银行
	BANKCODE_NBCB = "1056" // 宁波银行
)

type UnifiedOrder struct {
	Body           string `json:"body"`             // 必填。商品简单描述
	OutTradeNo     string `json:"out_trade_no"`     // 必填。商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一。
	SpbillCreateIp string `json:"spbill_create_ip"` // 必填。支持IPV4和IPV6两种格式的IP地址。调用微信支付API的机器IP（127.0.0.1）
	NotifyUrl      string `json:"notify_url"`       // 必填。异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。
	TradeType      string `json:"trade_type"`       // 必填。JSAPI--JSAPI支付（或小程序支付）、NATIVE--Native支付、APP--app支付，MWEB--H5支付
	TotalFee       int    `json:"total_fee"`        // 必填。订单总金额，单位为分
	Openid         string `json:"openid"`           // 必填。用户唯一openid
	DeviceInfo     string `json:"device_info"`      // 可选。自定义参数，可以为终端设备号(门店号或收银设备ID)，PC网页或公众号内支付可以传"WEB"
	SignType       string `json:"sign_type"`        // 可选。签名类型，默认为MD5，支持HMAC-SHA256和MD5。
	Detail         string `json:"detail"`           // 可选。商品详细描述
	Attach         string `json:"attach"`           // 可选。附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用。
	TimeStart      string `json:"time_start"`       // 可选。订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。
	TimeExpire     string `json:"time_expire"`      // 可选。订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。订单失效时间是针对订单号而言的，由于在请求支付的时候有一个必传参数prepay_id只有两小时的有效期，所以在重入时间超过2小时的时候需要重新请求下单接口获取新的prepay_id。
	GoodsTag       string `json:"goods_tag"`        // 可选。订单优惠标记，使用代金券或立减优惠功能时需要的参数
	ProductId      string `json:"product_id"`       // 可选。trade_type=NATIVE时，此参数必传。此参数为二维码中包含的商品ID，商户自行定义。
	LimitPay       string `json:"limit_pay"`        // 可选。上传此参数no_credit--可限制用户不能使用信用卡支付
	Receipt        string `json:"receipt"`          // 可选。Y，传入Y时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效
}

type OrderQuerys struct {
	TransactionId string `json:"transaction_id"` // 与商户订单号二选一。微信的订单号，优先使用
	OutTradeNo    string `json:"out_trade_no"`   // 与微信订单号二选一。商户系统内部订单号
	SignType      string `json:"sign_type"`      // 可选。签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
}

type OrderCloses struct {
	OutTradeNo string `json:"out_trade_no"` // 必填。商户系统内部订单号
	SignType   string `json:"sign_type"`    // 可选。签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
}

type OrderRefunds struct {
	TransactionId string `json:"transaction_id"` // 与商户订单号二选一。微信的订单号，优先使用
	OutTradeNo    string `json:"out_trade_no"`   // 与微信订单号二选一。商户系统内部订单号
	SignType      string `json:"sign_type"`      // 可选。签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	OutRefundNo   string `json:"out_refund_no"`  // 商户系统内部的退款单号
	TotalFee      int    `json:"total_fee"`      // 订单总金额
	RefundFee     int    `json:"refund_fee"`     // 退款总金额
	RefundDesc    string `json:"refund_desc"`    // 退款原因
	NotifyUrl     string `json:"notify_url"`     // 异步接收微信支付退款结果通知的回调地址，通知URL必须为外网可访问的url，不允许带参数。如果参数中传了notify_url，则商户平台上配置的回调地址将不会生效。
}

type Transfers struct {
	DeviceInfo     string `json:"device_info"`      // 可选。设备号。微信支付分配的终端设备号
	PartnerTradeNo string `json:"partner_trade_no"` // 必填。商户订单号，需保持唯一性
	Openid         string `json:"openid"`           // 必填。用户openid
	CheckName      string `json:"check_name"`       // 必填。校验用户姓名选项。NO_CHECK：不校验真实姓名  FORCE_CHECK：强校验真实姓名
	ReUserName     string `json:"re_user_name"`     // 可选。收款用户真实姓名
	Amount         int    `json:"amount"`           // 必填。金额
	Desc           string `json:"desc"`             // 必填。备注
	SpbillCreateIp string `json:"spbill_create_ip"` // 必填。ip地址
}

type PayBank struct {
	PartnerTradeNo string `json:"partner_trade_no"` // 必填。商户订单号，需保持唯一性
	EncBankNo      string `json:"enc_bank_no"`      // 必填。收款方银行卡号
	EncTrueName    string `json:"enc_true_name"`    // 必填。收款方用户名
	BankCode       string `json:"bank_code"`        // 必填。收款方开户行编号
	Amount         int    `json:"amount"`           // 必填。付款金额，单位分
	Desc           string `json:"desc"`             // 可选。说明
}