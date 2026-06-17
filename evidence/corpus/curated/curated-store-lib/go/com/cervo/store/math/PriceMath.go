package math

func AddTax(subtotal int, tax int) int {
	return subtotal + tax
}
func Discount(subtotal int, amount int) int {
	if subtotal > amount {
		return subtotal - amount
	}
	return 0
}
