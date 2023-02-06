package demoReq

import "go-archetype/domain/vo/request"

// 查询请求
type ListRequest struct {
	request.PageRequest        // 分页参数
	Id                  int64  `json:"id"`   // 主键
	Name                string `json:"name"` // 名称
	Sex                 string `json:"sex"`  // 性别
}
