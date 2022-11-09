package secondModule

import (
	"fmt"
	"time"
)

func GetTime() string {
	now := fmt.Sprintf("waktu sekarang : %v", time.Now())
	return now
}