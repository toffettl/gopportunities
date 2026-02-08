package schemas

import (
	"gorm.io/gorm"
)

// struc - entity
type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Ling     string
	Salary   int64
}
