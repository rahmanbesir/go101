package main

type Order struct {
}

type OrderRepository interface {
	getById(int) Order
}

type test struct {
}

func (o *test) getById(id int) Order {
	return Order{}
}
