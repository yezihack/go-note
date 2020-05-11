package mongdbService

import (
	"database/sql"
	"fmt"
)

type MongDBService struct {
	db *sql.DB
}

func (m *MongDBService) Find() {
	fmt.Println("MongDBService find")
}
func (m *MongDBService) Delete() {
	fmt.Println("MongDBService delete")
}
