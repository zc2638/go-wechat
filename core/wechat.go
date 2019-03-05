package core

const (
	// 微信业务接口域名
	WECHAT_DOMAIN = "https://api.weixin.qq.com"
	// 通用异地容灾业务接口域名
	WECHAT_DOMAIN2 = "https://api2.weixin.qq.com"
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
	SCOPE_SNSAPI_BASE = "snsapi_base"
	SCOPE_SNSAPI_USERINFO = "snsapi_userinfo"
)

const (
	// 获取access_token
	WECHAT_ACCESSTOKEN = WECHAT_DOMAIN + "/cgi-bin/token"
)

const (
	// 小程序access_token本地路径
	WECHAT_XCX_ACCESSTOKEN_PATH = "access_token_xcx.txt"
	// 公众号access_token本地路径
	WECHAT_PUBLIC_ACCESSTOKEN_PATH = "access_token_public.txt"
)

const (
	// 商户统一下单
	MCH_UNIFIEDORDER = WECHAT_MCH_DOMAIN + "/pay/unifiedorder"
	// 商户查询订单
	MCH_QUERYORDER = WECHAT_MCH_DOMAIN + "/pay/orderquery"
	// 商户关闭订单
	MCH_CLOSEORDER = WECHAT_MCH_DOMAIN + "/pay/closeorder"
	// 商户订单退款
	MCH_REFUNDORDER = WECHAT_MCH_DOMAIN + "/secapi/pay/refund"
)

const (
	// 小程序code2session
	XCX_CODE2SESSION = WECHAT_DOMAIN + "/sns/jscode2session"
	// 小程序模板发送
	XCX_TEMPLATE_SEND = WECHAT_DOMAIN + "/cgi-bin/message/wxopen/template/send"
)

const (
	// 网页授权
	PUBLIC_AUTHORIZEURL = "https://open.weixin.qq.com/connect/oauth2/authorize"
	// 获取用户信息
	PUBLIC_AUTHORIZEINFO = WECHAT_DOMAIN + "/sns/oauth2/access_token"
	// 刷新用户access_token
	PUBLIC_REFRESHTOKEN = WECHAT_DOMAIN + "/sns/oauth2/refresh_token"
	// 拉取用户信息
	PUBLIC_USERINFO = WECHAT_DOMAIN + "/sns/userinfo"
	//
	PUBLIC_CHECKTOKEN = WECHAT_DOMAIN + "/sns/auth"
)