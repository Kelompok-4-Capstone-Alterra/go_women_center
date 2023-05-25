package helper

func GetOffsetAndLimit(page, limit int) (int, int) {

	if page <= 0 {
		page = 1
	}

	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 10
	}

	offset := (page - 1) * limit

	return offset, limit
}

func GetTotalPages(totalData, limit int) int {
	if totalData <= 0 {
		return 1
	}

	totalPages := totalData / limit
	if totalData%limit > 0 {
		totalPages++
	}

	return totalPages
}