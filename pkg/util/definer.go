package util

func DefinePageLimitPagination(page, limit int) (pagePagination int, limitPagination int) {
	pagePagination = 1
	limitPagination = 10

	if page != 0 {
		pagePagination = page
	}
	if limit != 0 {
		limitPagination = limit
	}

	return
}
