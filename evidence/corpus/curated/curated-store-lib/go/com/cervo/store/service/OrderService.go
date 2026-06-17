package service

import "github.com/cervantesh/cervo-transpiler-java-to-go/evidence/corpus/curated/curated-store-lib/go/com/cervo/store/math"

func Total(subtotal int, tax int) int {
	return math.AddTax(subtotal, tax)
}
func DiscountedTotal(subtotal int, amount int) int {
	return math.Discount(subtotal, amount)
}
