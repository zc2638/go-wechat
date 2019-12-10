package xcx

import (
	"encoding/json"
	"github.com/zc2638/wechat"
	"github.com/zc2638/wechat/config"
	"github.com/zctod/go-tool/common/curlx"
)

/**
 * Created by zc on 2019/12/10.
 */
// 模板消息发送
type MessageSend struct {
	accessToken string
	TemplateMsg TemplateMsg
	Err         error
	Result      wechat.ResCode
}

func (m *MessageSend) Sent(drive wechat.Drive) {
	m.accessToken, m.Err = drive.BuildAccessToken()
}

func (m *MessageSend) Exec() {

	tb, err := json.Marshal(m.TemplateMsg)
	if err != nil {
		m.Err = err
		return
	}

	h := curlx.HttpReq{
		Url: config.XCX_MESSAGE_SEND,
		Query: map[string]string{
			"access_token": m.accessToken,
		},
		Body: tb,
		Header: map[string]string{
			"Content-Type": "application/json; encoding=utf-8",
		},
	}

	b, err := h.Post()
	if err != nil {
		m.Err = err
		return
	}

	var res wechat.ResCode
	if err := json.Unmarshal(b, &res); err != nil {
		m.Err = err
		return
	}
	m.Result = res
}

// 创建被分享动态消息的 activity_id
type MessageActivityCreate struct {
	accessToken string
	Err         error
	Result      MessageActivityCreateResult
}

type MessageActivityCreateResult struct {
	ActivityId     string `json:"activity_id"`     // 动态消息的 ID
	ExpirationTime int    `json:"expiration_time"` // activity_id 的过期时间戳。默认24小时后过期
	ErrCode        int    `json:"errcode"`         // 错误码。0：请求成功，-1：系统繁忙。此时请开发者稍候再试，42001：access_token 过期
}

func (m *MessageActivityCreate) Sent(drive wechat.Drive) {
	m.accessToken, m.Err = drive.BuildAccessToken()
}

func (m *MessageActivityCreate) Exec() {

	h := curlx.HttpReq{
		Url: config.XCX_MESSAGE_ACTIVITY_CREATE,
		Query: map[string]string{
			"access_token": m.accessToken,
		},
	}

	b, err := h.Get()
	if err != nil {
		m.Err = err
		return
	}

	var res MessageActivityCreateResult
	if err := json.Unmarshal(b, &res); err != nil {
		m.Err = err
		return
	}
	m.Result = res
}

// 修改被分享的动态消息
type MessageActivityUpdate struct {
	accessToken  string               `json:"-"`
	ActivityId   string               `json:"activity_id"`   // 动态消息的 ID
	TargetState  int                  `json:"target_state"`  // 动态消息修改后的状态:0未开始，1已开始
	TemplateInfo ActivityTemplateInfo `json:"template_info"` // 动态消息对应的模板信息
	Err          error                `json:"-"`
	Result       wechat.ResCode       `json:"-"`
}

func (m *MessageActivityUpdate) Sent(drive wechat.Drive) {
	m.accessToken, m.Err = drive.BuildAccessToken()
}

func (m *MessageActivityUpdate) Exec() {

	mb, err := json.Marshal(m)
	if err != nil {
		m.Err = err
		return
	}

	h := curlx.HttpReq{
		Url: config.XCX_MESSAGE_ACTIVITY_UPDATE,
		Query: map[string]string{
			"access_token": m.accessToken,
		},
		Body: mb,
	}

	b, err := h.Post()
	if err != nil {
		m.Err = err
		return
	}

	var res wechat.ResCode
	if err := json.Unmarshal(b, &res); err != nil {
		m.Err = err
		return
	}
	m.Result = res
}
