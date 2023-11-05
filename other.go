package es6

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Includes(arrs *[]string, str string) bool {
	for _, a := range *arrs {
		if a == str {
			return true
		}
	}
	return false
}

func GetUUID() (uuid string) {
	b := make([]byte, 16)
	fmt.Println(b)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}

func GetCurrentTimestamp() string {
	currentTime := time.Now().UTC()
	timestamp := currentTime.Format("2006-01-02 1:04:05")
	return timestamp

}
