package types

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Outbox struct {
	gorm.Model
	Data   datatypes.JSON
	RefID  uint
	Type   uint8
	Status uint8
}
