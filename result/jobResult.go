package result

import (
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

// JobResult job返回
func JobResult(ctx context.Context, resp interface{}, err error) {

	if err == nil {
		//成功返回 ,只有dev环境下才会打印info，线上不显示
		if resp != nil {
			logx.Info("resp:%+v", resp)
		}
		return
	} else {
		errCode := uint32(10001)
		errMsg := "服务器开小差啦，稍后再来试一试"

		//错误返回
		causeErr := errors.Cause(err) // err类型
		var e *CodeError
		if errors.As(causeErr, &e) { //自定义错误类型
			//自定义CodeError
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		}

		logx.WithContext(ctx).Errorf("【JOB-ERR】 : %+v ,errCode:%d , errMsg:%s ", err, errCode, errMsg)
		return
	}
}
