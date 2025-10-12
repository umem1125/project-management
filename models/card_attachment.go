package models

import (
	"time"

	"github.com/google/uuid"
)

type CardAttachment struct {
	InternalID	int64		`json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID	uuid.UUID	`json:"public_id" db:"public_id"`
	CardID		int64		`json:"card_id" db:"card_id"`
	UserID		int64		`json:"user_id" db:"user_id"`
	File		string		`json:"file" db:"file"`
	CreatedAt	time.Time	`json:"created_at" db:"created_at"`
}