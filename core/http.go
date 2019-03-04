package core

import (
	"crypto/tls"
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
	Url      string
	Params   map[string]interface{}
	CertFile string
	KeyFile  string
}

func (h *HttpReq) CurlWithCert(method string) ([]byte, error) {

	cert, err := tls.LoadX509KeyPair(h.CertFile, h.KeyFile)
	if err != nil {
		return nil, err
	}
	//certBytes, err := ioutil.ReadFile("client.pem")
	//if err != nil {
	//	panic("Unable to read cert.pem")
	//}
	//clientCertPool := x509.NewCertPool()
	//ok := clientCertPool.AppendCertsFromPEM(certBytes)
	//if !ok {
	//	panic("failed to parse root certificate")
	//}
	conf := &tls.Config{
		//RootCAs:            clientCertPool,
		Certificates: []tls.Certificate{cert},
	}
	transport := &http.Transport{
		TLSClientConfig:    conf,
		DisableCompression: true,
	}
	var req = &http.Client{Transport: transport}
	var resp *http.Response
	var data string

	if h.Params != nil {
		if method == METHOD_XML {
			b := utils.MapToXml(utils.MapToStringMap(h.Params))
			data = string(b)
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
		resp, err = req.Get(h.Url)
		break
	case METHOD_POST:
		resp, err = req.Post(h.Url, "application/x-www-form-urlencoded", strings.NewReader(data))
		break
	case METHOD_XML:
		resp, err = req.Post(h.Url, "text/xml; charset=utf-8", strings.NewReader(data))
		break
	default:
		resp, err = req.Post(h.Url, "application/x-www-form-urlencoded", strings.NewReader(data))
		break
	}
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
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
			b := utils.MapToXml(xmlData)
			data = string(b)
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

func (h *HttpReq) XmlData() (map[string]string, error) {

	b, err := h.Curl(METHOD_XML)
	if err != nil {
		return nil, err
	}
	return utils.XmlToMap(b), nil
}

func (h *HttpReq) XmlDataWithCert() (map[string]string, error) {

	b, err := h.CurlWithCert(METHOD_XML)
	if err != nil {
		return nil, err
	}
	return utils.XmlToMap(b), nil
}
