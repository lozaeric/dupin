package utils

import (
	"os"
	"time"

	"github.com/rs/xid"
)

func Now() string {
	return time.Now().Format(time.RFC3339)
}

func GenerateID() string {
	return xid.New().String()
}

func IsProduction() bool {
	return os.Getenv("ENV") == "production"
}
