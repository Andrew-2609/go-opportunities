package schema

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role      string
	Company   string
	Location  string
	WorkModel string
	Link      string
	Salary    int64
}
