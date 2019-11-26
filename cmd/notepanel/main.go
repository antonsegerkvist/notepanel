package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/notepanel/internal/notepanel"
)

func main() {
	fmt.Println(`==> Booting notepanel server`)
	notepanel.Run()
	fmt.Println(`==> Ending notepanel server`)
}
