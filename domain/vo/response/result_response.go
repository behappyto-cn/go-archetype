package response

import (
	"github.com/gin-gonic/gin"
	respStatusCode "go-archetype/common/constant/status"
)

// ConvertResult 转换结果
func ConvertResult(context *gin.Context, result bool, msg string) {
	if result {
		context.JSON(respStatusCode.OK, SuccessResultMsg(result, msg))
	} else {
		context.JSON(respStatusCode.OK, Failure(msg))
	}
}

// ConvertPageResult 转换分页结果
func ConvertPageResult(context *gin.Context, result bool, msg string, pageResult *PageResult) {
	if !result {
		context.JSON(respStatusCode.OK, Failure(msg))
	} else if pageResult == nil || pageResult.Total == 0 || pageResult.Records == nil {
		// 按需验证 pageResult.Total == 0 || pageResult.Records == nil
		context.JSON(respStatusCode.OK, Failure(msg))
	} else {
		context.JSON(respStatusCode.OK, SuccessResult(pageResult))
	}
}
