package core

const (
	// 微信业务接口域名
	WECHAT_DOMAIN = "https://api.weixin.qq.com"
	// 微信商户接口域名
	WECHAT_MCH_DOMAIN = "https://api.mch.weixin.qq.com"
)

// 签名方式
const (
	SIGNTYPE_MD5         = "MD5"
	SIGNTYPE_SHA1        = "SHA1"
	SIGNTYPE_HMAC_SHA256 = "HMAC-SHA256"
)

const (
	// 小程序access_token本地路径
	WECHAT_XCX_ACCESSTOKEN_PATH = "access_token_xcx.txt"
)

const (
	// 小程序code2session
	WECHAT_XCX_CODE2SESSION = WECHAT_DOMAIN + "/sns/jscode2session"
	// 小程序获取access_token
	WECHAT_XCX_ACCESSTOKEN = WECHAT_DOMAIN + "/cgi-bin/token"
	// 小程序模板发送
	WECHAT_XCX_TEMPLATE_SEND = WECHAT_DOMAIN + "/cgi-bin/message/wxopen/template/send"
	// 小程序统一下单
	WECHAT_XCX_UNIFIEDORDER = WECHAT_MCH_DOMAIN + "/pay/unifiedorder"
	// 小程序查询订单
	WECHAT_XCX_QUERYORDER = WECHAT_MCH_DOMAIN + "/pay/orderquery"
	// 小程序关闭订单
	WECHAT_XCX_CLOSEORDER = WECHAT_MCH_DOMAIN + "/pay/closeorder"
	// 小程序订单退款
	WECHAT_XCX_REFUNDORDER = WECHAT_MCH_DOMAIN + "/secapi/pay/refund"
)
