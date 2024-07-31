package pkg

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		limit, page := GetPageQueries(ctx)

		switch {
		case limit > 1000:
			limit = 1000
		case limit <= 0:
			limit = 1000
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

func PaginationDetails(limit, page, dataLength uint64) (nextPage, currentPage, total uint64) {
	if dataLength > (page * limit) {
		nextPage = page + 1
	} else {
		nextPage = 0
	}

	total = dataLength
	currentPage = page
	return
}

func GetPageQueries(ctx *gin.Context) (limit, page int) {
	limitQuery := ctx.Query("limit")
	pageQuery := ctx.Query("page")

	if limitQuery == "" {
		limit = 20
	} else {
		limit, _ = strconv.Atoi(limitQuery)
	}

	if pageQuery == "" {
		page = 1
	} else {
		page, _ = strconv.Atoi(pageQuery)
	}

	return
}
