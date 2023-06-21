package pagination

import (
	"github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/internal/helpers/pagination/models"
)

type PaginationImpl struct{}

func NewPaginationImpl() *PaginationImpl {
	return &PaginationImpl{}
}

func (pgn *PaginationImpl) CalculatePagination(limit, page, recordCount int) *models.Pagination {
	pagination := models.Pagination{}

	total := (recordCount / limit)

	// caculator total page
	remainder := (recordCount % total)
	if remainder == 0 {
		pagination.TotalPage = total
	} else {
		pagination.TotalPage = total + 1
	}

	pagination.CurrentPage = page
	pagination.RecordPerPage = limit

	// calculate next/prev
	if page <= 0 {
		pagination.Next = page + 1
	} else if page < pagination.TotalPage {
		pagination.Previous = page - 1
		pagination.Next = page + 1
	} else if page == pagination.TotalPage {
		pagination.Previous = page - 1
		pagination.Next = 0
	}

	return &pagination
}
