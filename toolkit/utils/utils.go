package utils

import (
	"time"

	"github.com/rs/xid"
)

func Now() string {
	return time.Now().Format(time.RFC3339)
}

func GenerateID() string {
	return xid.New().String()
}
