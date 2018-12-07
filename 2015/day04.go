package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

const input = "bgvyzdsv"

func main() {

	for i := 1; i < 10000000; i++ {
		data := input + strconv.Itoa(i)
		md5 := fmt.Sprintf("%x", md5.Sum([]byte(data)))
		if strings.HasPrefix(md5, "000000") {
			fmt.Println(i)
			break
		}
	}
}
