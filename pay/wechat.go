package pay

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
)

type WxPayTool struct {
	Conf   WxConf
	Client *wechat.ClientV3
}

func NewWxPayTool(conf WxConf) (*WxPayTool, error) {
	client, err := wechat.NewClientV3(conf.MchId, conf.SerialNo, conf.ApiV3Key, conf.PrivateKey)
	if err != nil {
		return nil, err
	}
	fmt.Println(client)
	err = client.AutoVerifySign()
	if err != nil {
		return nil, err
	}
	// 自定义配置http请求接收返回结果body大小，默认 10MB
	//client.SetBodySize() // 没有特殊需求，可忽略此配置
	// 打开Debug开关，输出日志，默认是关闭的
	client.DebugSwitch = gopay.DebugOn
	return &WxPayTool{
		Conf:   conf,
		Client: client}, nil

}

// WxJsPay 支付，在微信支付服务后台生成预支付交易单
func (p *WxPayTool) WxJsPay(ctx context.Context, openId string, total int64, tradeNo string, description string) (*wechat.PrepayRsp, error) {
	//第三方30分钟过期
	expire := time.Now().Add(30 * 60 * time.Second).Format(time.RFC3339)
	bm := make(gopay.BodyMap)
	bm.Set("appid", p.Conf.AppMiniId).
		Set("description", description).
		Set("out_trade_no", tradeNo).
		Set("time_expire", expire).
		Set("notify_url", p.Conf.NotifyUrl).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", total).
				Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) {
			bm.Set("openid", openId)
		})

	wxRsp, err := p.Client.V3TransactionJsapi(ctx, bm)
	if err != nil {
		return nil, err
	}
	if wxRsp.Code == 200 {
		return nil, err
	}
	return wxRsp, nil
}

// WxRefund 退款
func (p *WxPayTool) WxRefund(ctx context.Context, refund int64, total int64, transactionId string, refundNo string, reason string) (*wechat.RefundRsp, string) {

	bm := make(gopay.BodyMap)
	// 商户订单号（支付后返回的，42000开头）
	bm.Set("transaction_id", transactionId).
		Set("sign_type", "MD5").
		// 必填 退款订单号（程序员定义的）
		Set("out_refund_no", refundNo).
		// 选填 退款描述
		Set("reason", reason).
		Set("notify_url", p.Conf.NotifyUrl).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			// 退款金额:单位是分
			bm.Set("refund", refund). //实际退款金额
							Set("total", total). // 折扣前总金额（不是实际退款数）
							Set("currency", "CNY")
		})
	wxRsp, err := p.Client.V3Refund(ctx, bm)
	if err != nil {
		fmt.Println(fmt.Sprintf("-----wxPay WxRefund() -------error:%s", err.Error()))
		return nil, err.Error()
	}
	return wxRsp, wxRsp.Error
}

// WxQueryRefund 查询退款状态
func (p *WxPayTool) WxQueryRefund(refundNo string) (*wechat.RefundQueryRsp, error) {

	wxRsp, err := p.Client.V3RefundQuery(context.Background(), refundNo, nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("-----wxPay WxQueryRefund() -------error:%s", err.Error()))
	}
	return wxRsp, err
}

// WxTestV3Query 交易查询
func (p *WxPayTool) WxTestV3Query(no string) *wechat.QueryOrderRsp {
	wxRsp, err := p.Client.V3TransactionQueryOrder(context.Background(), wechat.OutTradeNo, no)
	if err != nil {
		fmt.Println(fmt.Sprintf("-----wxPay TestV3QueryOrder() -------error:%s", err.Error()))
		return nil
	}
	return wxRsp
}

func (p *WxPayTool) Notify(req *http.Request) (error, *wechat.V3DecryptBusifavorResult) {
	fmt.Println("--------------------- WxPayNotify START ---------------------")
	notifyReq, err := wechat.V3ParseNotify(req)
	if err != nil {
		fmt.Println("------ WxPayNotify V3ParseNotify ERR ------", err.Error())
		return err, nil
	}

	// 验证异步通知的签名
	err = notifyReq.VerifySignByPK(p.Client.WxPublicKey())
	if err != nil {
		fmt.Println("------ WxPayNotify VerifySignByPKMap ERR ------", err.Error())
		return err, nil
	}
	// 普通支付通知解密
	result, err := notifyReq.DecryptBusifavorCipherText(p.Conf.ApiV3Key)
	if err != nil {
		fmt.Println("------ WxPayNotify DecryptCipherText Error ------", err.Error())
		return err, nil
	}
	return nil, result
	//if result != nil && result.TradeState == "SUCCESS" {
	//	fmt.Println("------ WxPayNotify PushMessToPayQueue START 【" + result.OutTradeNo + "】------")
	//	var wxReq = make(map[string]interface{})
	//	promotionAmount := 0
	//	for i := range result.PromotionDetail {
	//		promotionAmount += result.PromotionDetail[i].Amount
	//	}
	//	fmt.Println(fmt.Sprintf("------ WxPayNotify 优惠券总额 promotionAmount 【%d】------", promotionAmount))
	//	//商户订单号:商户系统内部订单号
	//	wxReq["pay_no"] = result.OutTradeNo
	//	//微信支付订单号:微信支付系统生成的订单号。
	//	wxReq["trade_no"] = result.TransactionId
	//	//与支付宝同步
	//	wxReq["trade_status"] = "TRADE_SUCCESS"
	//	wxReq["notify_time"] = result.SuccessTime
	//	wxReq["total_amount"] = result.Amount.Total
	//	wxReq["receipt_amount"] = result.Amount.PayerTotal + promotionAmount
	//	var mapData = make(map[string]interface{})
	//	mapData["data_type"] = "PayNotify"
	//	mapData["param"] = map[string]interface{}{"payType": "wxPay", "notifyReq": wxReq}
	//}

}
