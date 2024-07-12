package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	PageNumber int `json:"pageNumber" form:"pageNumber"` // 页码
	PageSize   int `json:"pageSize" form:"pageSize"`     // 每页大小
}
