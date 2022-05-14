package dbent

import (
	"fmt"
	"os"

	"github.com/gushikem01/usa-kabu-go/server/ent"
	_ "github.com/lib/pq"
)

type pgConn struct {
	host     string
	dbname   string
	user     string
	password string
	port     string
}

func New() (*ent.Client, error) {
	conn := getConnectionSetting()
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conn.host, conn.port, conn.user, conn.dbname, conn.password))

	return client, err
}

// getConnectionSetting 接続情報
func getConnectionSetting() *pgConn {
	c := &pgConn{
		host:     os.Getenv("POSTGRES_HOST"),
		dbname:   os.Getenv("POSTGRES_DBNAME"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		port:     os.Getenv("POSTGRES_PORT"),
	}
	fmt.Println(c)
	return c
}
