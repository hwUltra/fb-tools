package pay

import (
	"context"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type AliPayTool struct {
	Conf   AliConf
	Client *alipay.Client
}

func GetAliClient(conf AliConf) *AliPayTool {
	var syOnce sync.Once
	var aliPayClient *alipay.Client
	syOnce.Do(func() {
		aliPayClient, err := alipay.NewClient(conf.AppId, conf.PrivateKey, conf.IsProd)
		if err != nil {
			log.Panic(err.Error())
		}
		//配置公共参数
		aliPayClient.SetCharset("utf-8").
			SetSignType(alipay.RSA2).
			SetNotifyUrl("")
	})
	return &AliPayTool{
		Conf:   conf,
		Client: aliPayClient,
	}
}

// TradeCreate 统一收单交易创建接口
func (p *AliPayTool) TradeCreate(ctx context.Context, openId string, payNo string, payAmount string, description string) (*alipay.TradeCreateResponse, error) {
	//统一收单线下交易预创建
	bm := make(gopay.BodyMap)
	bm.Set("notify_url", p.Conf.NotifyUrl)
	bm.Set("subject", description).
		//买家支付宝用户ID
		Set("buyer_id", openId).
		//商户订单号。由商家自定义
		Set("out_trade_no", payNo).
		//元，小数点两位,去掉千分位
		Set("total_amount", strings.Replace(payAmount, ",", "", -1))

	payParam, err := p.Client.TradeCreate(ctx, bm)
	if err != nil {
		fmt.Println(fmt.Sprintf("-----aliPay TradeCreate()-------error:%s", err.Error()))
		return nil, err
	}
	return payParam, nil
}

// Refund 退款接口
func (p *AliPayTool) Refund(bm gopay.BodyMap) *alipay.TradeRefundResponse {
	//统一收单线下交易预创建
	payParam, err := p.Client.TradeRefund(context.Background(), bm)
	if err != nil {
		fmt.Println(fmt.Sprintf("-----aliPay Refund()-------error:%s", err.Error()))
		return nil
	}
	//接口返回fund_change=Y为退款成功，fund_change=N或无此字段值返回时需通过退款查询接口进一步确认退款状态
	return payParam
}

// QueryRefund 查询退款接口
func (p *AliPayTool) QueryRefund(orderNo, refundNo string) (*alipay.TradeFastpayRefundQueryResponse, error) {
	//统一收单线下交易预创建
	bm := make(gopay.BodyMap)
	//商户订单号
	bm.Set("out_trade_no", orderNo)
	//退款请求号
	bm.Set("out_request_no", refundNo)
	payParam, err := p.Client.TradeFastPayRefundQuery(context.Background(), bm)
	if err != nil {
		fmt.Println(fmt.Sprintf("-----aliPay QueryRefund()-------error:%s", err.Error()))
	}
	//接口返回fund_change=Y为退款成功，fund_change=N或无此字段值返回时需通过退款查询接口进一步确认退款状态
	return payParam, err
}

// TradeQuery 交易查询接口
func (p *AliPayTool) TradeQuery(orderNo string) *alipay.TradeQueryResponse {

	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", orderNo)
	//统一收单线下交易预创建
	payParam, err := p.Client.TradeQuery(context.Background(), bm)
	if err != nil {
		fmt.Println(fmt.Sprintf("-----aliPay TradeQuery(%s)-------error:%s", orderNo, err.Error()))
		return nil
	}
	return payParam
}

func (p *AliPayTool) AlipayNotify(req *http.Request) {
	notifyReq, err := alipay.ParseNotifyToBodyMap(req)
	if err != nil {
		fmt.Println("------ AlipayNotify ERR  ------", err.Error())
		return
	}
	fmt.Println("------ AlipayNotify notifyReq  ------", notifyReq)

	ok, err := alipay.VerifySign(p.Conf.PublicKey, notifyReq)
	if ok {
		if notifyReq["receipt_amount"] == nil {
			fmt.Println("------ AlipayNotify 支付宝退款回调,忽略 ------")
		} else {
			//验签成功
			var alReq = make(map[string]interface{})
			//商户订单号:商户系统内部订单号
			alReq["pay_no"] = notifyReq["out_trade_no"]
			alReq["trade_no"] = notifyReq["trade_no"]
			//与支付宝同步
			alReq["trade_status"] = notifyReq["trade_status"]
			alReq["notify_time"] = notifyReq["notify_time"]
			//订单金额。本次交易支付的订单金额，单位为人民币（元）
			totalAmountF, _ := strconv.ParseFloat(notifyReq["total_amount"].(string), 64)
			totalAmount := int(totalAmountF * 100)
			//实收金额。商家在交易中实际收到的款项，单位为人民币（元）。
			receiptAmountF, _ := strconv.ParseFloat(notifyReq["receipt_amount"].(string), 64)
			receiptAmount := int(receiptAmountF * 100)
			alReq["total_amount"] = totalAmount
			alReq["receipt_amount"] = receiptAmount

			var mapData = make(map[string]interface{})
			mapData["data_type"] = "PayNotify"
			mapData["param"] = map[string]interface{}{"payType": "aliPay", "notifyReq": alReq}
		}
	}

	/*	notifyReq["order_no"] = "订单号"
		notifyReq["trade_no"] = "2022061322001402060501532448"
		notifyReq["trade_status"] = "TRADE_SUCCESS"
		notifyReq["notify_time"] = "2022-06-13 10:09:36"
		notifyReq["total_amount"] = 202000
		notifyReq["receipt_amount"] = 200000
		var mapData = make(map[string]interface{})
		mapData["data_type"] = "PayNotify"
		mapData["param"] = map[string]interface{}{"payType": "aliPay", "notifyReq": notifyReq}
	*/

}
