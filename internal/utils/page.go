package utils

import "math"

// 分页
type Page struct {
	PageIndex int64 `json:"page_index"` // 页面索引
	PageSize  int64 `json:"page_size"`  // 每页大小
	PageTotal int64 `json:"page_total"` // 总分页数
	Count     int64 `json:"count"`      // 当页记录数
	Total     int64 `json:"total"`      // 总记录数
}

const (
	PageIndexDefault = 1
	PageSizeDefault  = 20
)

/**
初始化分页
*/
func (p *Page) InitPage(total int64) {
	if p.PageIndex <= 0 {
		p.PageIndex = PageIndexDefault
	}
	if p.PageSize <= 0 {
		p.PageSize = PageSizeDefault
	}
	// 最大的索引
	maxIndex := int64(math.Ceil(float64(total) / float64(p.PageSize)))
	if maxIndex != 0 && maxIndex < p.PageIndex {
		p.PageIndex = maxIndex
	}
	// 总记录数
	p.Total = total
	// 总分数页
	if maxIndex <= 0 { // 没数据情况
		p.PageTotal = 0
		p.Count = 0
	} else {
		p.PageTotal = maxIndex
	}
	// 当页记录数
	if p.PageTotal > p.PageIndex { // 总页数大于当前页数
		p.Count = p.PageSize
	}
	if p.PageTotal == p.PageIndex { // 总页数等于当前页数
		p.Count = p.Total - (p.PageSize * (p.PageIndex - 1))
	}
}
