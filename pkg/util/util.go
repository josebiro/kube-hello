package util

import (
	"os"
)

func GetStage() string {
	return os.Getenv("STAGE")
}
