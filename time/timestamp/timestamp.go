package timestamp

import (
	"fmt"
	"time"
)

func GetTime() {
	currentTime :=time.Now()
	fmt.Println(currentTime)
}