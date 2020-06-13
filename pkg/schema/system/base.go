package system

import "time"

type BaseModel struct {
	CreatedBy string
	UpdatedBy string
	CreatedAt *time.Time `sql:"index:created"`
	UpdatedAt *time.Time `sql:"index:modified"`
	DeletedAt *time.Time `sql:"index:deleted"`
}
