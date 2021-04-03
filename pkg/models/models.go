package models

import (
	"fmt"
	"time"
)

type Mousetrap struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	OrgId       int64     `json:"org_id"`
	Status      bool      `json:"status"`
	LastTrigger time.Time `json:"last_trigger"`
}

type Organisation struct {
	Id       int64
	Name     string
	Password string
}

type Credentials struct {
	Name     string `json:"name"`
	Password string `json:"pass"`
}

func (c Credentials) CheckNotEmpty() error {
	if c.Password == "" || c.Name == "" {
		return fmt.Errorf("username or password is empty")
	}
	return nil
}
