package demoApi

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	respStatusCode "go-archetype/common/constant/status"
	demoReq "go-archetype/domain/vo/request/demo"
	"go-archetype/domain/vo/response"
	demoService "go-archetype/service/demo"
	"strconv"
)

// Demo 示例
func Demo(routerGroup *gin.RouterGroup) {

	// 新增数据
	routerGroup.POST("/demo/add/v1", func(context *gin.Context) {
		body, err := context.GetRawData()
		if err != nil {
			context.JSON(respStatusCode.OK, response.FailureResult(-1, "参数不合法"))
			return
		}
		// 反序列化
		var request demoReq.AddRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			log.Error("反序列化异常 ", err)
			context.JSON(respStatusCode.OK, response.FailureResult(-1, "参数不合法"))
			return
		}
		// 执行结果
		result, msg := demoService.Add(request)

		// 装换输出
		response.ConvertResult(context, result, msg)
	})

	// 删除数据
	routerGroup.DELETE("/demo/delete/v1", func(context *gin.Context) {
		// 执行结果
		idStr, ok := context.GetQuery("id")
		if !ok {
			context.JSON(respStatusCode.OK, response.FailureResult(-1, "参数不合法"))
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			context.JSON(respStatusCode.OK, response.FailureResult(-1, "参数不合法"))
			return
		}
		// 删除
		result, msg := demoService.Delete(id)
		// 装换输出
		response.ConvertResult(context, result, msg)
	})

	// 修改数据
	routerGroup.PUT("/demo/modify/v1", func(context *gin.Context) {
		body, err := context.GetRawData()
		if err != nil {
			context.JSON(respStatusCode.OK, response.FailureResult(-1, "参数不合法"))
			return
		}
		// 反序列化
		var request demoReq.EditRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			log.Error("反序列化异常 ", err)
			context.JSON(respStatusCode.OK, response.FailureResult(-1, "参数不合法"))
			return
		}
		// 执行结果
		result, msg := demoService.Modify(request)

		// 装换输出
		response.ConvertResult(context, result, msg)
	})

	// 查询数据
	routerGroup.POST("/demo/list/v1", func(context *gin.Context) {
		body, err := context.GetRawData()
		if err != nil {
			context.JSON(respStatusCode.OK, response.FailureResult(-1, "参数不合法"))
			return
		}
		// 反序列化
		var request demoReq.ListRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			log.Error("反序列化异常 ", err)
			context.JSON(respStatusCode.OK, response.FailureResult(-1, "参数不合法"))
			return
		}
		// 执行结果
		result, msg, pageResult := demoService.List(request)

		// 装换输出
		response.ConvertPageResult(context, result, msg, pageResult)
	})
}
