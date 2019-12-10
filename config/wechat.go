package config

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
	// access_token文件路径
	AccessTokenPath = "access_token.txt"
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
	// 企业付款到零钱
	MCH_TRANSFERS = WECHAT_MCH_DOMAIN + "/mmpaymkttransfers/promotion/transfers"
	// 查询企业付款到零钱
	MCH_TRANSFERSGET = WECHAT_MCH_DOMAIN + "/mmpaymkttransfers/gettransferinfo"
)

const (
	// 小程序code2session
	XCX_CODE2SESSION = WECHAT_DOMAIN + "/sns/jscode2session"
	// 获取支付用户unionId
	XCX_PAIDUNIONID = WECHAT_DOMAIN + "/wxa/getpaidunionid"
	// 小程序模板发送
	XCX_MESSAGE_SEND = WECHAT_DOMAIN + "/cgi-bin/message/wxopen/template/uniform_send"
	// 创建被分享动态消息的 activity_id
	XCX_MESSAGE_ACTIVITY_CREATE = WECHAT_DOMAIN + "/cgi-bin/message/wxopen/activityid/create"
	// 修改被分享的动态消息
	XCX_MESSAGE_ACTIVITY_UPDATE = WECHAT_DOMAIN + "/cgi-bin/message/wxopen/updatablemsg/send"
	// 获取小程序二维码
	XCX_QRCODE = WECHAT_DOMAIN + "/cgi-bin/wxaapp/createwxaqrcode"
	// 获取小程序码
	XCX_CODE = WECHAT_DOMAIN + "/wxa/getwxacode"
	// 获取小程序码（无限制）
	XCX_CODE_UNLIMITED = WECHAT_DOMAIN + "/wxa/getwxacodeunlimit"
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
	// 检查用户token是否过期
	PUBLIC_CHECKTOKEN = WECHAT_DOMAIN + "/sns/auth"
	// 自定义菜单创建
	PUBLIC_MENUCREATE = WECHAT_DOMAIN + "/cgi-bin/menu/create"
	// 自定义菜单查询
	PUBLIC_MENUGET = WECHAT_DOMAIN + "/cgi-bin/menu/get"
	// 自定义菜单删除
	PUBLIC_MENUDELETE = WECHAT_DOMAIN + "/cgi-bin/menu/delete"
	// 发送模板消息
	PUBLIC_TEMPLATE_SEND = WECHAT_DOMAIN + "/cgi-bin/message/template/send"
)