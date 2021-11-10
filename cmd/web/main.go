package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"a4lab2.net/snipb/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
)

type application struct {
	infoLog  *log.Logger
	errLog   *log.Logger
	session  *sessions.Session
	snippets *mysql.SnippetModel
	users    *mysql.UserModel
	cache    map[string]*template.Template
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {

	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {

	app.clientError(w, http.StatusNotFound)
}

type contextKey string

var contextKeyUser = contextKey("user")

func main() {
	addr := flag.String("addr", ":8080", "Http address from the command option")
	dsn := flag.String("dsn", "web:PASSHERE@/Snippetbox?parseTime=true", "Mysql DataSource")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)

	db, err := openDb(*dsn)
	if err != nil {
		errLog.Fatal(err)
	}
	defer db.Close()
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errLog.Fatal(err)
	}
	secret := flag.String("secret", "u46IpCV9y5Vlur8YvODJEhgOY8m9JVE4", "Secret")
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	app := &application{
		infoLog:  infoLog,
		errLog:   errLog,
		session:  session,
		snippets: &mysql.SnippetModel{DB: db},
		users:    &mysql.UserModel{DB: db},
		cache:    templateCache,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}
	infoLog.Println("Starting Server at ", *addr)
	errf := srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errLog.Fatal(errf)
}

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
