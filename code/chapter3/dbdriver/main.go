// Sample program to show how to show you how to briefly work
// with the sql package.
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/goinaction/code/chapter3/dbdriver/postgres"
)

// main is the entry point for the application.
func main() {
	fmt.Println("main")
	sql.Open("postgres", "mydb")
}
