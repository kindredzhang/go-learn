package main

import (
	"fmt"
)

// 定义函数类型
type PromotionFunc func(price float64) float64

// 为函数类型实现方法 ApplyDiscount
func (f PromotionFunc) ApplyDiscount(price float64) float64 {
	return f(price)
}

// 为函数类型实现方法 ApplyDiscount2
func (f PromotionFunc) ApplyDiscount2(price float64) float64 {
	return f(price)
}

// 定义一个无折扣策略
var noDiscount = PromotionFunc(func(price float64) float64 {
	return price
})

func main() {
	// 调用 ApplyDiscount 方法
	discountedPrice := noDiscount.ApplyDiscount(100)
	fmt.Printf("ApplyDiscount: %.2f\n", discountedPrice) // 输出：ApplyDiscount: 100.00

	// 调用 ApplyDiscount2 方法
	discountedPrice2 := noDiscount.ApplyDiscount2(100)
	fmt.Printf("ApplyDiscount2: %.2f\n", discountedPrice2) // 输出：ApplyDiscount2: 100.00
}
