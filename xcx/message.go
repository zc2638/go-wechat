package xcx

import (
	"encoding/json"
	"github.com/zc2638/gotool/curlx"
	"github.com/zc2638/wechat"
)

/**
 * Created by zc on 2019/12/10.
 */
// 模板消息发送
type MessageSend struct {
	TemplateMsg TemplateMsg
	Result      wechat.ResCode
}

func (m *MessageSend) Exec(drive wechat.Drive) error {
	accessToken, err := drive.BuildAccessToken()
	if err != nil {
		return err
	}
	tb, err := json.Marshal(m.TemplateMsg)
	if err != nil {
		return err
	}

	h := curlx.HttpReq{
		Url:  drive.GetHost() + "/cgi-bin/message/wxopen/template/uniform_send?access_token=" + accessToken,
		Body: tb,
		Header: map[string]string{
			curlx.HEADER_CONTENT_TYPE: curlx.CT_APPLICATION_JSON_UTF8,
		},
		Method: curlx.METHOD_POST,
	}
	return h.Do().ParseJSON(&m.Result)
}

// 创建被分享动态消息的 activity_id
type MessageActivityCreate struct {
	Result      MessageActivityCreateResult
}

type MessageActivityCreateResult struct {
	ActivityId     string `json:"activity_id"`     // 动态消息的 ID
	ExpirationTime int    `json:"expiration_time"` // activity_id 的过期时间戳。默认24小时后过期
	ErrCode        int    `json:"errcode"`         // 错误码。0：请求成功，-1：系统繁忙。此时请开发者稍候再试，42001：access_token 过期
}

func (m *MessageActivityCreate) Exec(drive wechat.Drive) error {
	accessToken, err := drive.BuildAccessToken()
	if err != nil {
		return err
	}

	h := curlx.HttpReq{
		Url: drive.GetHost() + "/cgi-bin/message/wxopen/activityid/create?access_token=" + accessToken,
	}
	return h.Do().ParseJSON(&m.Result)
}

// 修改被分享的动态消息
type MessageActivityUpdate struct {
	ActivityId   string               `json:"activity_id"`   // 动态消息的 ID
	TargetState  int                  `json:"target_state"`  // 动态消息修改后的状态:0未开始，1已开始
	TemplateInfo ActivityTemplateInfo `json:"template_info"` // 动态消息对应的模板信息
	Result       wechat.ResCode       `json:"-"`
}

func (m *MessageActivityUpdate) Exec(drive wechat.Drive) error {
	accessToken, err := drive.BuildAccessToken()
	if err != nil {
		return err
	}
	mb, err := json.Marshal(m)
	if err != nil {
		return err
	}

	h := curlx.HttpReq{
		Url:    drive.GetHost() + "/cgi-bin/message/wxopen/updatablemsg/send?access_token=" + accessToken,
		Body:   mb,
		Method: curlx.METHOD_POST,
	}
	return h.Do().ParseJSON(&m.Result)
}
