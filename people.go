package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type People struct {
	gorm.Model

	Name        string
	Age         int
	AddressList AddressList
}

type AddressList []*Address

func (p *People) TableName() string {
	return "people"
}

type Address struct {
	Country  string
	Province string
	City     string
	Detail   string
}

func (a *AddressList) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	data, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}
	return json.Unmarshal(data, &a)
}

func (a AddressList) Value() (driver.Value, error) {
	return json.Marshal(a)
}
