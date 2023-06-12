package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)
var Conn *pgx.Conn

func DatabaseConnect() {

	databaseUrl := "postgres://postgres:1234@localhost:5432/db_myblogs"
	var err error

	Conn, err = pgx.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect")
		os.Exit(1)
	}

	fmt.Println("Successfully Connected Database")

}