package public

import (
	"github.com/zc2638/gotool/curlx"
	"github.com/zc2638/wechat"
	"github.com/zc2638/wechat/config"
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
		Url:  config.PUBLIC_TEMPLATE_SEND + "?access_token=" + t.accessToken,
		Body: t.Message,
		Header: map[string]string{
			curlx.HEADER_CONTENT_TYPE: curlx.CT_APPLICATION_JSON_UTF8,
		},
		Method: curlx.METHOD_POST,
	}
	t.Err = h.Do().ParseJSON(&t.Result)
}
