package redisService

import (
	"database/sql"
	"fmt"
)

type RedisService struct {
	db *sql.DB
}

func (m *RedisService) Find() {
	fmt.Println("RedisService find")
}
func (m *RedisService) Delete() {
	fmt.Println("RedisService delete")
}
