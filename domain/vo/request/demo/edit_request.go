package demoReq

// 修改请求
type EditRequest struct {
	Id   int64  `json:"id"`   // 主键
	Name string `json:"name"` // 名称
	Sex  string `json:"sex"`  // 性别
}
