package models

import (
	"github.com/google/uuid"
)

const (
	DEBTOR   = "DEBTOR"
	DEFERRED = "DEFERRED"
	CLEAR    = "CLEAR"

	BUY  = "BUY"
	SELL = "SELL"
)

type ID = uuid.UUID

func NewID() ID {
	return uuid.New()
}
