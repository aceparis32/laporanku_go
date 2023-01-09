package utils

import (
	"net/http"

	"gorm.io/gorm"
)

func SearchSalesKeyword(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		keyword := q.Get("keyword")

		if keyword != "" {
			return db.Where("name ILIKE ?", "%"+keyword+"%")
		} else {
			return db
		}
	}
}

func SearchCompanyKeyword(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		keyword := q.Get("keyword")

		if keyword != "" {
			return db.Where("name ILIKE ? OR address ILIKE ?", "%"+keyword+"%", "%"+keyword+"%")
		} else {
			return db
		}
	}
}

func SearchItemKeyword(r *http.Request, column string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		keyword := q.Get("keyword")

		if keyword != "" {
			return db.Where("? ILIKE ?", column, "%"+keyword+"%")
		} else {
			return db
		}
	}
}
