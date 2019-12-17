package public

import (
	"github.com/zc2638/gotool/curlx"
	"github.com/zc2638/wechat"
)

/**
 * Created by zc on 2019/12/10.
 */
// 自定义菜单创建 TODO 可以将菜单结构化
type MenuCreate struct {
	Menu        []byte
	Result      wechat.ResCode
}

func (m *MenuCreate) Exec(drive wechat.Drive) error {
	accessToken, err := drive.BuildAccessToken()
	if err != nil {
		return err
	}
	h := curlx.HttpReq{
		Url:  drive.GetHost() + "/cgi-bin/menu/create?access_token=" + accessToken,
		Body: m.Menu,
		Header: map[string]string{
			curlx.HEADER_CONTENT_TYPE: curlx.CT_APPLICATION_JSON_UTF8,
		},
		Method: curlx.METHOD_POST,
	}
	return h.Do().ParseJSON(&m.Result)
}

// 自定义菜单查询 TODO 返回数据可以结构化
type MenuSearch struct {
	Result      []byte
}

func (m *MenuSearch) Exec(drive wechat.Drive) error {
	accessToken, err := drive.BuildAccessToken()
	if err != nil {
		return err
	}
	h := curlx.HttpReq{
		Url: drive.GetHost() + "/cgi-bin/menu/get?access_token=" + accessToken,
		Header: map[string]string{
			curlx.HEADER_CONTENT_TYPE: curlx.CT_APPLICATION_JSON_UTF8,
		},
	}
	m.Result, err = h.Post()
	return err
}

// 自定义菜单删除
type MenuDelete struct {
	Result      wechat.ResCode
}

func (m *MenuDelete) Exec(drive wechat.Drive) error {
	accessToken, err := drive.BuildAccessToken()
	if err != nil {
		return err
	}
	h := curlx.HttpReq{
		Url:    drive.GetHost() + "/cgi-bin/menu/delete?access_token=" + accessToken,
		Method: curlx.METHOD_GET,
	}
	return h.Do().ParseJSON(&m.Result)
}
