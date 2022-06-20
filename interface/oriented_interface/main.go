package main

import "fmt"

type ZhiFuBao struct {
	// 支付宝
}

type WeChat struct {
	// 微信
}

// Pay 支付宝的支付方法
func (z *ZhiFuBao) Pay(amount int64) {
	fmt.Printf("使用支付宝付款：%.2f元。\n", float64(amount/100))
}

// Pay 微信的支付方法
func (w *WeChat) Pay(amount int64) {
	fmt.Printf("使用微信付款：%.2f元。\n", float64(amount/100))
}

//// Checkout 结账
//func Checkout(obj *ZhiFuBao) {
//	// 支付100元
//	obj.Pay(100)
//}

// CheckoutWithZFB 支付宝结账
func CheckoutWithZFB(obj *ZhiFuBao) {
	// 支付100元
	obj.Pay(100)
}

// CheckoutWithWX 微信支付结账
func CheckoutWithWX(obj *WeChat) {
	// 支付100元
	obj.Pay(100)
}

// Payer 包含支付方法的接口类型
type Payer interface {
	Pay(int64)
}

// Checkout 结账
func Checkout(obj Payer) {
	// 支付100元
	obj.Pay(100)
}

func main() {
	Checkout(&ZhiFuBao{})
	CheckoutWithZFB(&ZhiFuBao{})
	CheckoutWithWX(&WeChat{})
	Checkout(&WeChat{})

}
