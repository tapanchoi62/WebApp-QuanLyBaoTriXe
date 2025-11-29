package utils

import (
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
	"gorm.io/gorm"
)

func Paginate[T any](db *gorm.DB, page int, pageSize int, search string, searchFields []string) (data []T, pagination models.Pagination, err error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	// Apply search
	if search != "" && len(searchFields) > 0 {
		like := "%" + search + "%"
		for i, field := range searchFields {
			if i == 0 {
				db = db.Where(field+" LIKE ?", like)
			} else {
				db = db.Or(field+" LIKE ?", like)
			}
		}
	}

	// Count total
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil,  models.Pagination{}, err
	}

	// Query data
	if err := db.Limit(pageSize).Offset(offset).Find(&data).Error; err != nil {
		return nil, models.Pagination{}, err
	}

	// Build pagination info
	pagination = models.Pagination{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: (total + int64(pageSize) - 1) / int64(pageSize),
		Search:     search,
	}

	return data, pagination, nil
}