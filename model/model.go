package model

import "gorm.io/gorm"

type User struct {
	gorm.Model        //Model a basic GoLang struct which includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `json:"name"`
	Address    string `json:"address"`
	Age        int    `json:"age"`
}
