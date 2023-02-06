package request

// 分页参数
type PageRequest struct {
	Offset   int64 `json:"offset"`
	PageSize int64 `json:"pageSize"`
}
