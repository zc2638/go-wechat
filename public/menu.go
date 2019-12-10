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
		Url: config.PUBLIC_MENUCREATE,
		Query: map[string]string{
			"access_token": m.accessToken,
		},
		Body: m.Menu,
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
		Url: config.PUBLIC_MENUGET,
		Query: map[string]string{
			"access_token": m.accessToken,
		},
		Header: map[string]string{
			"Content-Type": "application/json; encoding=utf-8",
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
		Url: config.PUBLIC_MENUDELETE,
		Query: map[string]string{
			"access_token": m.accessToken,
		},
	}

	b, err := h.Get()
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

