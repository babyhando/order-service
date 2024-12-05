package types

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	Content string
	To      uint
	Type    uint8
}
