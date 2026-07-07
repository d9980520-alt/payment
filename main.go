package main

import (
	"errors"
)

type transaction struct {
	description string
	usd         int
	isCancelled bool
}

type paymentModule struct {
	mapInfo map[int]transaction
	nextID  int
}

func NewPaymentModule() *paymentModule {
	return &paymentModule{
		mapInfo: make(map[int]transaction),
		nextID:  1,
	}
}

func (p *paymentModule) Buy(description string, usd int) (int, error) {
	if usd <= 0 {
		return 0, errors.New("ошибка: сумма должна быть больше нуля")
	}

	id := p.nextID
	p.nextID++

	p.mapInfo[id] = transaction{
		description: description,
		usd:         usd,
		isCancelled: false,
	}

	return id, nil
}

func (p *paymentModule) Cancell(id int) error {
	tx, exists := p.mapInfo[id]
	if !exists {
		return errors.New("ошибка: транзакция не найдена")
	}

	if tx.isCancelled {
		return errors.New("ошибка: транзакция уже была отменена")
	}

	tx.isCancelled = true
	p.mapInfo[id] = tx

	return nil
}

func (p *paymentModule) Info(id int) (transaction, error) {
	tx, exists := p.mapInfo[id]
	if !exists {
		return transaction{}, errors.New("ошибка: транзакция не найдена")
	}
	return tx, nil
}

func (p *paymentModule) AllInfo() map[int]transaction {
	return p.mapInfo
}

func main() {}
