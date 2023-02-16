package util

import (
	"math"
)

func CountOffsetPagination(page, limit int) int {
	if page == 0 || limit == 0 {
		return 0
	}
	return (page - 1) * limit
}

func CountTotalPagePagination(countDataPerPage, totalData int) int {
	if countDataPerPage == 0 || totalData == 0 {
		return 0
	}
	return int(math.Ceil(float64(totalData) / float64(countDataPerPage)))
}
