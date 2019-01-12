package model

import (
	"time"
)

type Task struct {
	Id        int       `gorm:"AUTO_INCREMENT"`
	Name      string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
}
