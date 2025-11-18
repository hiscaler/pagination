package pagination

// Pagination 翻页数据
type Pagination[T any] struct {
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
	TotalCount int  `json:"total_count"`
	PageCount  int  `json:"page_count"`
	IsLastPage bool `json:"is_last_page"`
	Items      []T  `json:"items"`
}

func New[T any](page, pageSize, total int) *Pagination[T] {
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

	return &Pagination[T]{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: total,
		PageCount:  pageCount,
		IsLastPage: page >= pageCount,
	}
}

func (p *Pagination[T]) SetItems(items []T) *Pagination[T] {
	p.Items = items
	return p
}

func (p *Pagination[T]) AddItem(item T) *Pagination[T] {
	p.Items = append(p.Items, item)
	return p
}
