package result

import (
	"fmt"
	"net/http"

	"github.com/hwUltra/fb-tools/xerr"

	"github.com/pkg/errors"
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
		errCode := xerr.SERVER_COMMON_ERROR
		errMsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)

		// err类型
		var e *xerr.CodeError
		if errors.As(causeErr, &e) { //自定义错误类型
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			if gStatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gStatus.Code())
				if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errCode = grpcCode * 100
				}
				errMsg = gStatus.Message()
			}
		}

		//if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
		//	errCode = e.GetErrCode()
		//	errMsg = e.GetErrMsg()
		//} else {
		//	if gStatus, ok := status.FromError(causeErr); ok { // grpc err错误
		//		grpcCode := uint32(gStatus.Code())
		//		if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
		//			errCode = grpcCode * 100
		//		}
		//		errMsg = gStatus.Message()
		//	}
		//
		//}

		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v %+v ", errCode, err)
		httpx.WriteJson(w, http.StatusOK, Error(errCode, errMsg))
		//httpx.WriteJson(w, http.StatusOK, errMsg)
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
		errCode := xerr.SERVER_COMMON_ERROR
		errMsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)
		// err类型
		var e *xerr.CodeError
		if errors.As(causeErr, &e) { //自定义错误类型
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			if gStatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gStatus.Code())
				if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errCode = grpcCode * 100
				}
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
	causeErr := errors.Cause(err)
	var e *xerr.CodeError
	if errors.As(causeErr, &e) { //自定义错误类型
		errMsg = fmt.Sprintf("%s ,%s", xerr.MapErrMsg(xerr.REUQES_PARAM_ERROR), e.GetErrMsg())
	} else {
		errMsg = fmt.Sprintf("%s ,%s", xerr.MapErrMsg(xerr.REUQES_PARAM_ERROR), err)
	}

	httpx.WriteJson(w, http.StatusOK, Error(xerr.REUQES_PARAM_ERROR, errMsg))
}
