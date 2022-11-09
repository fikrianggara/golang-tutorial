package time

import (
	"fmt"
	"time"
)

func GetTime() {
	currentTime :=time.Now()
	fmt.Println(currentTime)
}

func main(){
	fmt.Println("waktu sekarang ialah:")
	fmt.Println(time.Now())
}