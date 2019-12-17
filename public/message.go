package public

import (
	"github.com/zc2638/gotool/curlx"
	"github.com/zc2638/wechat"
)

/**
 * Created by zc on 2019/12/10.
 */
// 发送消息模板
type TemplateSend struct {
	Message     []byte
	Result      TemplateSendResult
}

type TemplateSendResult struct {
	MsgId string `json:"msgid"`
	wechat.ResCode
}

func (t *TemplateSend) Exec(drive wechat.Drive) error {
	accessToken, err := drive.BuildAccessToken()
	if err != nil {
		return err
	}
	h := curlx.HttpReq{
		Url:  drive.GetHost() + "/cgi-bin/message/template/send?access_token=" + accessToken,
		Body: t.Message,
		Header: map[string]string{
			curlx.HEADER_CONTENT_TYPE: curlx.CT_APPLICATION_JSON_UTF8,
		},
		Method: curlx.METHOD_POST,
	}
	return h.Do().ParseJSON(&t.Result)
}
