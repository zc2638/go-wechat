package core

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/zctod/go-tool/common/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	METHOD_POST = "POST"
	METHOD_GET  = "GET"
)

type M map[string]interface{}

type HttpReq struct {
	Url      string
	Method   string
	Header   map[string]string
	Params   M
	Body     []byte
	CertFile string
	KeyFile  string
}

func (h *HttpReq) buildBody() {

	if h.Body != nil {
		return
	}
	if h.Params == nil {
		return
	}

	var data string
	for k, v := range h.Params {
		if data != "" {
			data += "&"
		}
		data += k + "=" + v.(string)
	}

	switch h.Method {
	case METHOD_POST:
		h.Body = []byte(data)
		break
	case METHOD_GET:
		urlArr := strings.Split(h.Url, "?")
		if len(urlArr) == 2 {
			if data != "" {
				urlArr[1] = urlArr[1] + "&" + data
			}
			//将GET请求的参数进行转义
			h.Url = urlArr[0] + "?" + url.PathEscape(urlArr[1])
		}
		break
	}
}

func (h *HttpReq) Do() ([]byte, error) {

	var transport *http.Transport
	if h.CertFile != "" {
		cert, err := tls.LoadX509KeyPair(h.CertFile, h.KeyFile)
		if err != nil {
			return nil, err
		}
		transport = &http.Transport{
			DisableCompression: true,
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		}
	}
	var client = &http.Client{}
	if transport != nil {
		client.Transport = transport
	}
	h.buildBody()

	req, err := http.NewRequest(h.Method, h.Url, bytes.NewReader(h.Body))
	if err != nil {
		return nil, err
	}

	if h.Header != nil {
		for k, v := range h.Header {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func (h *HttpReq) Get() ([]byte, error) {

	h.Method = METHOD_GET
	return h.Do()
}

func (h *HttpReq) GetData() (M, error) {

	var res M
	b, err := h.Get()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &res)
	return res, err
}

func (h *HttpReq) Post() ([]byte, error) {

	h.Method = METHOD_POST
	return h.Do()
}

func (h *HttpReq) PostData() (M, error) {

	var res M
	b, err := h.Post()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &res)
	return res, err
}

func (h *HttpReq) Json() ([]byte, error) {

	h.Header = map[string]string{"Content-Type": "application/json; encoding=utf-8"}
	return h.Post()
}

func (h *HttpReq) JsonData() (M, error) {

	var res M
	b, err := h.Json()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &res)
	return res, err
}

func (h *HttpReq) XmlData() (map[string]string, error) {

	h.Method = METHOD_POST
	h.Header = map[string]string{"Content-Type": "text/xml; charset=utf-8"}
	h.Body = utils.MapToXml(utils.MapToStringMap(h.Params))
	b, err := h.Do()
	if err != nil {
		return nil, err
	}
	return utils.XmlToMap(b), nil
}
