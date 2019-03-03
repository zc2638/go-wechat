package xcx

import "encoding/json"

type Template struct {
	ToUser     string          `json:"touser"`      // 必须, 接受者OpenID
	TemplateId string          `json:"template_id"` // 必须, 模版ID
	FormId     string          `json:"form_id"`     // 必须, 表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
	Page       string          `json:"page"`        // 可选, 跳小程序所需(小程序页面路径)
	Data       json.RawMessage `json:"data"`        // 必须, 模板数据, JSON 格式的 []byte, 满足特定的模板需求
}

type UnifiedOrder struct {
	Body           string `json:"body"`             // 必填。商品简单描述
	OutTradeNo     string `json:"out_trade_no"`     // 必填。商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一。
	SpbillCreateIp string `json:"spbill_create_ip"` // 必填。支持IPV4和IPV6两种格式的IP地址。调用微信支付API的机器IP（127.0.0.1）
	NotifyUrl      string `json:"notify_url"`       // 必填。异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。
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
