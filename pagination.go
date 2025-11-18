package pagination

// Pagination 带翻页的数据
type Pagination struct {
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
	TotalCount int  `json:"total_count"`
	PageCount  int  `json:"page_count"`
	IsLastPage bool `json:"is_last_page"`
	Items      any  `json:"items"`
}

func New(page, pageSize, total int) *Pagination {
	if pageSize < 1 {
		pageSize = 10
	}

	pageCount := -1
	if total >= 0 {
		pageCount = (total + pageSize - 1) / pageSize
		if page > pageCount {
			page = pageCount
		}
	}
	if page < 1 {
		page = 1
	}

	return &Pagination{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: total,
		PageCount:  pageCount,
		IsLastPage: page >= pageCount,
	}
}

func (p *Pagination) SetItems(items any) *Pagination {
	p.Items = items
	return p
}
