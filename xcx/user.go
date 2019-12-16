package xcx

import (
	"github.com/zc2638/gotool/curlx"
	"github.com/zc2638/wechat"
	"github.com/zc2638/wechat/config"
)

/**
 * Created by zc on 2019/12/10.
 */
// 用户支付完成后，获取该用户的 UnionId，无需用户授权
type PaidSearch struct {
	accessToken   string
	Openid        string
	TransactionId string // 微信支付订单号（与商户订单号选其一即可，都填则优先支付订单号）
	MchId         string // 微信支付分配的商户号，和商户订单号配合使用
	OutTradeNo    string // 微信支付商户订单号，和商户号配合使用
	Err           error
	Result        PaidSearchResult
}

type PaidSearchResult struct {
	UnionId string `json:"unionid"`
	wechat.ResCode
}

func (u *PaidSearch) Sent(drive wechat.Drive) {
	u.accessToken, u.Err = drive.BuildAccessToken()
}

func (u *PaidSearch) Exec() {
	query := map[string]string{
		"access_token":   u.accessToken,
		"openid":         u.Openid,
		"transaction_id": u.TransactionId,
	}

	if u.TransactionId == "" {
		delete(query, "transaction_id")
		query["mch_id"] = u.MchId
		query["out_trade_no"] = u.OutTradeNo
	}

	h := curlx.HttpReq{
		Url:   config.XCX_PAIDUNIONID,
		Query: query,
		Method: curlx.METHOD_GET,
	}
	u.Err = h.Do().ParseJSON(&u.Result)
}
