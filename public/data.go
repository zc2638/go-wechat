package public

const (
	MENU_TYPE_CLICK            = "click"              // 点击推事件用户点击click类型按钮后，微信服务器会通过消息接口推送消息类型为event的结构给开发者（参考消息接口指南），并且带上按钮中开发者填写的key值，开发者可以通过自定义的key值与用户进行交互
	MENU_TYPE_VIEW             = "view"               // 跳转URL
	MENU_TYPE_SCANCODE_PUSH    = "scancode_push"      // 扫码推事件
	MENU_TYPE_SCANCODE_WAITMSG = "scancode_waitmsg"   // 扫码推事件且弹出“消息接收中”提示框
	MENU_TYPE_PIC_SYSPHOTO     = "pic_sysphoto"       // 弹出系统拍照发图用户
	MENU_TYPE_PIC_ALBUM        = "pic_photo_or_album" // 弹出拍照或者相册发图
	MENU_TYPE_PIC_WEIXIN       = "pic_weixin"         // 弹出微信相册发图器
	MENU_TYPE_LOCATION         = "location_select"    // 弹出地理位置选择器
	MENU_TYPE_MEDIA            = "media_id"           // 用户点击media_id类型按钮后，微信服务器会将开发者填写的永久素材id对应的素材下发给用户
	MENU_TYPE_VIEW_LIMITED     = "view_limited"       // 跳转图文消息URL
)

type AuthorizeUrl struct {
	RedirectUri string // 返回跳转链接
	Scope       string // 授权类型：静默snsapi_base, 正常snsapi_userinfo
	State       string // 额外值
}

type Menu struct {
	Button []MenuButton
}

type MenuButton struct {
	Button    MenuSubButton
	SubButton []MenuSubButton
}

type MenuSubButton struct {
	Type     string `json:"type"`     // 菜单的响应动作类型，view表示网页类型，click表示点击类型，miniprogram表示小程序类型
	Name     string `json:"name"`     // 菜单标题，不超过16个字节，子菜单不超过60个字节
	Key      string `json:"key"`      // 菜单KEY值，用于消息接口推送，不超过128字节
	Url      string `json:"url"`      // 网页 链接，用户点击菜单可打开链接，不超过1024字节
	MediaId  string `json:"media_id"` // 调用新增永久素材接口返回的合法media_id
	Appid    string `json:"appid"`    // 小程序的appid（仅认证公众号可配置）
	PagePath string `json:"pagepath"` // 小程序的页面路径
}

type Template struct {
	ToUser      string `json:"touser"`      // 接收者openid
	TemplateId  string `json:"template_id"` // 模板ID
	Url         string `json:"url"`         // 模板跳转链接（海外帐号没有跳转能力）
	Data        string `json:"data"`        // 模板数据
	Color       string `json:"color"`       // 模板内容字体颜色，不填默认为黑色
	MiniProgram MiniProgram                 // 跳小程序所需数据，不需跳小程序可不用传该数据
}

type MiniProgram struct {
	Appid    string `json:"appid"`    // 所需跳转到的小程序appid
	PagePath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布
}
