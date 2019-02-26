package main

import (
	"fmt"
	"regexp"
)

const text = `My name is YuWan,My Email is yuwang@abc.qq.com 1234433344
aadd  cumian@163.com dddd
addk  yugang@qq.com.cn ...
yugang0127@gmail.com ...
abcdefg geniusxiaoyu@gmai.com abc
...eieie
`

func main() {
	compile := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9.]+`)
	match := compile.FindAllString(text, -1)

	for _, m := range match {
		fmt.Println(m)
	}
}
