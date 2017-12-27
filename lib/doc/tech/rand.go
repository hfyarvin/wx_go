package main

import (
	"fmt"
	"math/rand" //伪随机数
)

func GetRandInt(max int) int {
	r := rand.Intn(max)
	return r
}

func main() {
	// 0 <= n < 10
	fmt.Println(GetRandInt(10), GetRandInt(10), GetRandInt(10), GetRandInt(10), GetRandInt(10))
	r1, r2 := GetRandFloat(5)
	fmt.Println(r1, r2)
	fmt.Println(SetSeed(100))
}

//获取float
func GetRandFloat(min int) (float64, float64) {
	r1 := (rand.Float64() * 5) + float64(min)
	r2 := (rand.Float64() * 5) + float64(min)
	return r1, r2
}

func SetSeed(max int) int {
	// 为了使随机数生成器具有确定性，可以给它一个seed
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)
	return r1.Intn(max)
}
