package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
	models "github.com/musale/snippets/pkg/models/mysql"
)

type webApp struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
	session       *sessions.Session
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	addr := flag.String("addr", "4000", "HTTP network address")
	dsn := flag.String("dsn", "root:root@tcp(127.0.0.1:3307)/snippetbox?parseTime=true", "MySQL data source name")
	secret := flag.String("secret", "c9be21e559f9d3172d95cc2f0abed91e", "Secret Key")
	flag.Parse()
	addrString := fmt.Sprintf(":%s", *addr)

	templateCache, err := newTemplateCache("./ui/html")
	if err != nil {
		errorLog.Fatal(err)
	}

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	app := &webApp{
		infoLog:       infoLog,
		errorLog:      errorLog,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
		session:       session,
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
