package mytime

import (
    "time"
    "fmt"
)

func Now() string {
    today := time.Now()
    return fmt.Sprintf("%v", today)
}