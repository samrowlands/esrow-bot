package main

import (
	"flag"
	"fmt"
)

var (
	Token string
)

const KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	m := make(map[int]int)

	m[1] = 1
	m[2] = 2
	m[3] = 3
	m[4] = 4
	m[5] = 5

	s := []*int{}
	for _, v := range m {
		n := v
		s = append(s, &n)
	}

	for _, i := range s {
		fmt.Println(*i)
	}
}
