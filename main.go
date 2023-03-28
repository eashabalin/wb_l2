package main

import (
	"fmt"
	"wb_l2/pattern"
)

func main() {
	builder := pattern.MinskBuilder{}
	director := pattern.NewMotorbikeBuildDirector(&builder)
	director.Construct()
	motorbike := director.GetResult()
	fmt.Println(motorbike)
}
