package mysqlService

import (
	"database/sql"
	"fmt"
)

type MySQLService struct {
	db *sql.DB
}

func (m *MySQLService) Find() {
	fmt.Println("MySQLService find")
}
func (m *MySQLService) Delete() {
	fmt.Println("MySQLService delete")
}
