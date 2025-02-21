package handler

// CalculateOffsetAndLimit calculates the offset and limit based on the page number and page size
func CalculateOffsetAndLimit(page, pageSize int) (pages int, offset int, limit int) {
	if page < 1 {
		pages = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	offset = (pages - 1) * pageSize
	if pageSize > 1000 {
		limit = 1000
	} else {
		limit = pageSize
	}

	return pages, offset, limit
}

// CalculateTotalPages calculates the total pages based on the number of items and page size
func CalculateTotalPages(totalItems, pageSize int64) int64 {
	if pageSize < 1 {
		pageSize = 10
	}

	totalPages := totalItems / pageSize
	if totalItems%pageSize != 0 {
		totalPages++
	}

	return totalPages
}
