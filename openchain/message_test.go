package openchain

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestT(t *testing.T) {
	//now := time.Now().UnixNano()
	//
	//timeInMilli := now / int64(time.Millisecond)
	//fmt.Println(timeInMilli)
	//
	//timeFromMill := time.UnixMilli(timeInMilli)
	//fmt.Println(timeFromMill)
	//
	//fmt.Println(time.Now().UnixMilli())
	//fmt.Println(time.UnixMilli(time.Now().UnixMilli()))

	timeString := strconv.Itoa(int(time.Now().UnixMilli()))

	fmt.Println(timeString)

}

//RefreshExpireAt: time.Unix(int64(claims["exp"].(float64)), 0),
//
