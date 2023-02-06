package demoReq

// 新增请求
type AddRequest struct {
	Name string `json:"name"` // 名称
	Sex  string `json:"sex"`  // 性别
}
