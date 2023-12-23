package main

import (
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

type CurrencyEntity struct {
	Bid float32
	gorm.Model
}

func connect() (*gorm.DB, error) {
	conn, err := gorm.Open(sqlite.Open("./db/exchange.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = conn.AutoMigrate(&CurrencyEntity{})
	if err != nil {
		log.Panicf("Data migration failed: %s", err)
	}

	return conn, nil
}

func SaveExchangeInfo(currency CurrencyResponse) (CurrencyEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	db, err := connect()
	if err != nil {
		log.Panicf("DB Connection failed: %s", err)
	}

	bid, err := strconv.ParseFloat(currency.USDBRL.Bid, 32)
	if err != nil {
		log.Panicf("Invalid value received: %s", err)
	}

	entity := CurrencyEntity{Bid: float32(bid)}

	tx := db.WithContext(ctx).Create(&entity)

	if tx.Error != nil {
		return CurrencyEntity{}, tx.Error
	}

	return entity, nil
}
