package mch

import (
	"errors"
	"github.com/zc2638/wechat"
)

/**
 * Created by zc on 2019/12/17.
 */
// 企业付款到零钱
type Transfer struct {
	DeviceInfo     string         `json:"device_info"`      // 可选。设备号。微信支付分配的终端设备号
	PartnerTradeNo string         `json:"partner_trade_no"` // 必填。商户订单号，需保持唯一性
	Openid         string         `json:"openid"`           // 必填。用户openid
	CheckName      string         `json:"check_name"`       // 必填。校验用户姓名选项。NO_CHECK：不校验真实姓名  FORCE_CHECK：强校验真实姓名
	ReUserName     string         `json:"re_user_name"`     // 可选。收款用户真实姓名
	Amount         int            `json:"amount"`           // 必填。金额
	Desc           string         `json:"desc"`             // 必填。备注
	SpbillCreateIp string         `json:"spbill_create_ip"` // 必填。ip地址
	Result         TransferResult `json:"-"`
}

type TransferResult struct {
	wechat.Return
	PartnerTradeNo string `json:"partner_trade_no"` // 商户订单号，需保持历史全局唯一性(只能是字母或者数字，不能包含有其他字符)
	PaymentNo      string `json:"payment_no"`       // 企业付款成功，返回的微信付款单号
	PaymentTime    string `json:"payment_time"`     // 企业付款成功时间
}

func (t *Transfer) Exec(drive wechat.Drive) error {
	if t.CheckName == "" {
		t.CheckName = CHECK_NAME_FALSE
	}
	if t.CheckName == CHECK_NAME_TRUE && t.ReUserName == "" {
		return errors.New("选择强制校验真实姓名，必须填写真实姓名")
	}
	if t.SpbillCreateIp == "" {
		t.SpbillCreateIp = "127.0.0.1"
	}
	return drive.GetMerchant().Exec("/mmpaymkttransfers/promotion/transfers", t, &t.Result, true)
}

// 企业付款到零钱查询
type TransferSearch struct {
	PartnerTradeNo string               `json:"partner_trade_no"` // 商户调用企业付款API时使用的商户订单号
	Result         TransferSearchResult `json:"-"`
}

type TransferSearchResult struct {
	wechat.Return
	PartnerTradeNo string `json:"partner_trade_no"` // 商户使用查询API填写的单号的原路返回.
	DetailId       string `json:"detail_id"`        // 付款单号。调用企业付款API时，微信系统内部产生的单号
	Status         string `json:"status"`           // 转账状态。SUCCESS:转账成功，FAILED:转账失败，PROCESSING:处理中
	Reason         string `json:"reason"`           // 失败原因
	Openid         string `json:"openid"`           // 收款用户openid
	TransferName   string `json:"transfer_name"`    // 收款用户姓名
	PaymentAmount  int    `json:"payment_amount"`   // 付款金额。单位为“分”
	TransferTime   string `json:"transfer_time"`    // 转账时间
	PaymentTime    string `json:"payment_time"`     // 付款成功时间
	Desc           string `json:"desc"`             // 企业付款备注
}

func (t *TransferSearch) Exec(drive wechat.Drive) error {
	return drive.GetMerchant().Exec("/mmpaymkttransfers/gettransferinfo", t, &t.Result, true)
}
