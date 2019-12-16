package public

import (
	"github.com/zc2638/gotool/curlx"
	"github.com/zc2638/wechat"
	"github.com/zc2638/wechat/config"
)

/**
 * Created by zc on 2019/12/10.
 */
// 自定义菜单创建 TODO 可以将菜单结构化
type MenuCreate struct {
	accessToken string
	Menu        []byte
	Err         error
	Result      wechat.ResCode
}

func (m *MenuCreate) Sent(drive wechat.Drive) {
	m.accessToken, m.Err = drive.BuildAccessToken()
}

func (m *MenuCreate) Exec() {
	h := curlx.HttpReq{
		Url:  config.PUBLIC_MENUCREATE + "?access_token=" + m.accessToken,
		Body: m.Menu,
		Header: map[string]string{
			curlx.HEADER_CONTENT_TYPE: curlx.CT_APPLICATION_JSON_UTF8,
		},
		Method: curlx.METHOD_POST,
	}
	m.Err = h.Do().ParseJSON(&m.Result)
}

// 自定义菜单查询 TODO 返回数据可以结构化
type MenuSearch struct {
	accessToken string
	Err         error
	Result      []byte
}

func (m *MenuSearch) Sent(drive wechat.Drive) {
	m.accessToken, m.Err = drive.BuildAccessToken()
}

func (m *MenuSearch) Exec() {
	h := curlx.HttpReq{
		Url: config.PUBLIC_MENUGET + "?access_token=" + m.accessToken,
		Header: map[string]string{
			curlx.HEADER_CONTENT_TYPE: curlx.CT_APPLICATION_JSON_UTF8,
		},
	}
	m.Result, m.Err = h.Post()
}

// 自定义菜单删除
type MenuDelete struct {
	accessToken string
	Err         error
	Result      wechat.ResCode
}

func (m *MenuDelete) Sent(drive wechat.Drive) {
	m.accessToken, m.Err = drive.BuildAccessToken()
}

func (m *MenuDelete) Exec() {
	h := curlx.HttpReq{
		Url:    config.PUBLIC_MENUDELETE + "?access_token=" + m.accessToken,
		Method: curlx.METHOD_GET,
	}
	m.Err = h.Do().ParseJSON(&m.Result)
}
