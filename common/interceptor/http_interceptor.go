package interceptor

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	respStatusCode "go-archetype/common/constant/status"
	"go-archetype/common/util/date"
	"go-archetype/common/util/ip"
	"go-archetype/common/util/sign"
	"go-archetype/domain/vo/response"
	"time"
)

// HttpInterceptor http 请求拦截器
func HttpInterceptor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 响应
		realIp := ipUtil.GetRealIp(ctx.Request)
		fullPath := ctx.FullPath()
		log.Info("请求地址：", fullPath, "，请求ip：", realIp, "，Referer：", ctx.Request.Referer())

		// 签名验证
		sign := signUtil.GetMd5ByStr(dateUtil.FormatYyyyMmDd(time.Now()))
		signParam := ctx.GetHeader("sign")
		if sign != signParam {
			log.Error("签名不正确：", sign, " : ", signParam)
			ctx.JSON(respStatusCode.OK, response.Failure("参数不合法"))
			ctx.Abort()
			return
		}

		// 执行下一个
		ctx.Next()
	}
}
