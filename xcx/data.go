package xcx

import "encoding/json"

type Template struct {
	ToUser     string          `json:"touser"`      // 必须, 接受者OpenID
	TemplateId string          `json:"template_id"` // 必须, 模版ID
	FormId     string          `json:"form_id"`     // 必须, 表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
	Page       string          `json:"page"`        // 可选, 跳小程序所需(小程序页面路径)
	Data       json.RawMessage `json:"data"`        // 必须, 模板数据, JSON 格式的 []byte, 满足特定的模板需求
}