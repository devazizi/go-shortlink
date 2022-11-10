package adapter

import "time"

type Link struct {
	ID         uint   `gorm:"primaryKey"`
	Short      string `gorm:"type:varchar(255)"`
	FullLink   string
	ClickCount uint `gorm:"default(0)"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
