package result

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

// HttpResult http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errCode := uint32(10001)
		errMsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Unwrap(err)

		// err类型
		var e *CodeError
		if errors.As(causeErr, &e) { //自定义错误类型
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			if gStatus, ok := status.FromError(causeErr); ok { // grpc err错误
				errCode = uint32(gStatus.Code())
				errMsg = gStatus.Message()
			}
		}
		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v %+v ", errCode, err)
		httpx.WriteJson(w, http.StatusOK, Error(errCode, errMsg))
	}
}

// AuthHttpResult 授权的http方法
func AuthHttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errCode := uint32(10001)
		errMsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Unwrap(err)
		// err类型
		var e *CodeError
		if errors.As(causeErr, &e) { //自定义错误类型
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			if gStatus, ok := status.FromError(causeErr); ok { // grpc err错误
				errCode = uint32(gStatus.Code())
				errMsg = gStatus.Message()
			}
		}

		logx.WithContext(r.Context()).Errorf("【GATEWAY-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusOK, Error(errCode, errMsg))
	}
}

// ParamErrorResult http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := ""
	causeErr := errors.Unwrap(err)
	var e *CodeError
	if errors.As(causeErr, &e) { //自定义错误类型
		errMsg = fmt.Sprintf("参数错误 ,%s", e.GetErrMsg())
	} else {
		errMsg = fmt.Sprintf("参数错误 ,%s", err)
	}

	httpx.WriteJson(w, http.StatusOK, Error(10001, errMsg))
}
