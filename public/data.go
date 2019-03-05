package public

type AuthorizeUrl struct {
	RedirectUri string // 返回跳转链接
	Scope string // 授权类型：静默snsapi_base, 正常snsapi_userinfo
	State string // 额外值
}