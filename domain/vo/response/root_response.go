package response

import "strings"

// Result 返回结果
type Result struct {
	BaseResult
	Data interface{} `json:"data"`
}

// BaseResult 基本返回信息
type BaseResult struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	UserMsg string `json:"userMsg"`
}

// Result 返回分页结果
type PageResult struct {
	Total   int64       `json:"total"`
	Records interface{} `json:"records"`
}

// SuccessResult 成功返回信息
func SuccessResult(data interface{}) *Result {
	r := new(Result)
	r.Code = 0
	r.Msg = "OK"
	r.UserMsg = "操作成功"
	r.Data = data
	return r
}

// SuccessResult 成功返回信息
func SuccessResultMsg(data interface{}, msg string) *Result {
	if len(strings.TrimSpace(msg)) == 0 {
		msg = "操作成功"
	}
	r := new(Result)
	r.Code = 0
	r.Msg = msg
	r.UserMsg = msg
	r.Data = data
	return r
}

// Failure 失败返回信息
func Failure(msg string) *Result {
	if len(strings.TrimSpace(msg)) == 0 {
		msg = "操作成功"
	}
	r := new(Result)
	r.Code = -1
	r.Msg = msg
	r.UserMsg = msg
	return r
}

// FailureResult 失败返回信息
func FailureResult(code int, msg string) *Result {
	if len(strings.TrimSpace(msg)) == 0 {
		msg = "操作成功"
	}
	r := new(Result)
	r.Code = code
	r.Msg = msg
	r.UserMsg = msg
	return r
}

// 分页结果
func PageResultData(total int64, records interface{}) *PageResult {
	r := new(PageResult)
	r.Total = total
	r.Records = records
	return r
}
