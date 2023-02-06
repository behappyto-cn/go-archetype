package demoService

import (
	log "github.com/sirupsen/logrus"
	demoReq "go-archetype/domain/vo/request/demo"
	"go-archetype/domain/vo/response"
)

// Add 新增数据
func Add(addRequest demoReq.AddRequest) (bool, string) {
	log.Info(addRequest)
	// --------- 数据库操作 ---------

	return true, "新增成功"
}

// Delete 删除数据
func Delete(id int) (bool, string) {
	log.Info(id)
	// --------- 数据库操作 ---------

	return true, "删除成功"
}

// Modify 修改数据
func Modify(editRequest demoReq.EditRequest) (bool, string) {
	log.Info(editRequest)
	// --------- 数据库操作 ---------

	return true, "修改成功"
}

// List 查询数据
func List(listRequest demoReq.ListRequest) (bool, string, *response.PageResult) {
	log.Info(listRequest)

	// --------- 数据库操作 ---------
	var total int64         // 总条数
	var records interface{} // 列表数据

	// 返回结果
	return true, "查询成功", response.PageResultData(total, records)
}
