package core

import (
	"encoding/json"
	"github.com/zctod/tool/common/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	METHOD_POST = "POST"
	METHOD_GET  = "GET"
	METHOD_XML  = "XML"
)

type M map[string]interface{}

type HttpReq struct {
	Url    string
	Params map[string]interface{}
}

func (h *HttpReq) Curl(method string) ([]byte, error) {

	var req *http.Request
	var data string

	if h.Params != nil {
		if method == METHOD_XML {
			var xmlData = make(map[string]string)
			for k, v := range h.Params {
				xmlData[k] = v.(string)
			}
			b, err := utils.MapToXml(xmlData)
			if err == nil {
				data = string(b)
			}
		} else {
			for k, v := range h.Params {
				if data != "" {
					data += "&"
				}
				data += k + "=" + v.(string)
			}
		}
	}

	switch method {
	case METHOD_GET:
		urlArr := strings.Split(h.Url, "?")
		if len(urlArr) == 2 {
			//将GET请求的参数进行转义
			h.Url = urlArr[0] + "?" + url.PathEscape(urlArr[1])
		}
		if h.Params != nil && data != "" {
			urlArr := strings.Split(h.Url, "?")
			if len(urlArr) == 1 {
				h.Url = urlArr[0] + "?" + url.PathEscape(data)
			} else if len(urlArr) == 2 {
				h.Url = urlArr[0] + "&" + url.PathEscape(data)
			}
		}
		req, _ = http.NewRequest(method, h.Url, nil)
		break
	case METHOD_POST:
		req, _ = http.NewRequest(method, h.Url, strings.NewReader(data))
		break
	case METHOD_XML:
		req, _ = http.NewRequest(METHOD_POST, h.Url, strings.NewReader(data))
		req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	default:
		req, _ = http.NewRequest(method, h.Url, strings.NewReader(data))
		break
	}

	resp, err := new(http.Client).Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func (h *HttpReq) Get(res interface{}) error {

	b, err := h.Curl(METHOD_GET)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &res)
}

func (h *HttpReq) GetData() (res M, err error) {

	err = h.Get(&res)
	return
}

func (h *HttpReq) Post(res interface{}) error {

	b, err := h.Curl(METHOD_POST)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &res)
}

func (h *HttpReq) PostData() (res M, err error) {

	err = h.Post(&res)
	return
}

func (h *HttpReq) PostXml() (map[string]string, error) {

	b, err := h.Curl(METHOD_XML)
	if err != nil {
		return nil, err
	}
	return utils.XmlToMap(b)
}
