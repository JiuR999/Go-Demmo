package main

import "fmt"

type Payment struct {
	context  *PaymentContext
	strategy PayStrategy
}

type PaymentContext struct {
	Name, ID string
	Money    int
}

type PayStrategy interface {
	Pay(ctx *PaymentContext)
}

func (p Payment) pay() {
	p.strategy.Pay(p.context)
}

type Cash struct{}

func (c Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("%s使用现金支付%d元", ctx.Name, ctx.Money)
}

type Bank struct{}

func (c Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("%s使用银行卡支付%d元", ctx.Name, ctx.Money)
}

func main() {
	pContext := &PaymentContext{
		Name:  "张三",
		ID:    "123456",
		Money: 100,
	}
	cash := &Bank{}
	payment := &Payment{
		context:  pContext,
		strategy: cash,
	}
	payment.pay()
}
