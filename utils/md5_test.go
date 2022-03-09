package utils

import (
	"fmt"
	"testing"
)

func TestMD5NowID(t *testing.T) {
	md5 := MD5NowID()
	fmt.Printf("len:%v, value:%v", len(md5), md5)

}

func TestMD5Now(t *testing.T) {
	md5 := MD5Now()
	fmt.Printf("len:%v, value:%v", len(md5), md5)

}
