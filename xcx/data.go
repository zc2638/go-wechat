package xcx

import "encoding/json"

type TemplateMsg struct {
	ToUser   string `json:"touser"` // 必须, 接受者OpenID
	Template Template
}

type Template struct {
	TemplateId      string          `json:"template_id"`      // 必须, 模版ID
	FormId          string          `json:"form_id"`          // 必须, 表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
	Page            string          `json:"page"`             // 可选, 跳小程序所需(小程序页面路径)
	Data            json.RawMessage `json:"data"`             // 可选, 模板内容，不填则下发空模板
	EmphasisKeyword string          `json:"emphasis_keyword"` // 可选, 模板需要放大的关键词，不填则默认无放大
}

type ActivityTemplateInfo struct {
	ParameterList []ParameterList `json:"parameter_list"`
}

type ParameterList struct {
	Name  string
	Value string
}
