package functions

import (
	"fmt"
	"time"
)

func Today() string {

	today := time.Now()

	return fmt.Sprintf("Dia: %d-%d-%d Hora: %d:%d:%d", today.Day(), int(today.Month()), today.Year(),
		today.Hour(), today.Minute(), today.Second())

}

func Log(message string) {

	fmt.Println(message, Today())

}
