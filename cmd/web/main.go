package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type webApp struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	addr := flag.String("addr", "4000", "HTTP network address")
	dsn := flag.String("dsn", "root:root@tcp(127.0.0.1:3307)/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()
	addrString := fmt.Sprintf(":%s", *addr)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &webApp{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	server := &http.Server{
		Addr:     addrString,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server listening on port %s", addrString)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}

// openDB creates a connection pool to mysql
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
