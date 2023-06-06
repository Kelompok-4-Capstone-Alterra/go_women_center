package helper

func GetPaginateData(page, limit int, device ...string) (int, int, int) {

	if page <= 0 {
		page = 1
	}
	maxLimit := 100
	minLimit := 10

	if len(device) > 0 && device[0] == "mobile" {
		maxLimit = 50
		minLimit = 5
	}

	switch {
	case limit > maxLimit:
		limit = maxLimit
	case limit <= 0:
		limit = minLimit
	}

	offset := (page - 1) * limit

	return page, offset, limit
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