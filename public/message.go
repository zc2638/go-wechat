package public

import (
	"encoding/json"
	"github.com/zc2638/wechat"
	"github.com/zc2638/wechat/config"
	"github.com/zctod/go-tool/common/curlx"
)

/**
 * Created by zc on 2019/12/10.
 */
// 发送消息模板
type TemplateSend struct {
	accessToken string
	Message     []byte
	Err         error
	Result      TemplateSendResult
}

type TemplateSendResult struct {
	MsgId string `json:"msgid"`
	wechat.ResCode
}

func (t *TemplateSend) Sent(drive wechat.Drive) {
	t.accessToken, t.Err = drive.BuildAccessToken()
}

func (t *TemplateSend) Exec() {

	h := curlx.HttpReq{
		Url: config.PUBLIC_TEMPLATE_SEND,
		Query: map[string]string{
			"access_token": t.accessToken,
		},
		Body: t.Message,
		Header: map[string]string{
			"Content-Type": "application/json; encoding=utf-8",
		},
	}

	b, err := h.Post()
	if err != nil {
		t.Err = err
		return
	}

	var res TemplateSendResult
	if err := json.Unmarshal(b, &res); err != nil {
		t.Err = err
		return
	}
	t.Result = res
}
