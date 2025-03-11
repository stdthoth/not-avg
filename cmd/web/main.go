package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/joho/godotenv"

	"github.com/stdthoth/not-avg/db"
)

const version = "1.0.0"
const cssVersion = "1.0.0"

type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	paystack struct {
		secret string
		key    string
	}
}

type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	app.infoLog.Printf("starting HTTP server in %s mode on port %d", app.config.env, app.config.port)

	return srv.ListenAndServe()

}

func main() {
	godotenv.Load()

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "server port to listen on")
	flag.StringVar(&cfg.db.dsn, "dsn", "root:foot5print@tcp(localhost:3306)/notaverage?parseTime=true&tls=false", "DSN")
	flag.StringVar(&cfg.env, "env", "development", "application environment {development|production}")
	flag.StringVar(&cfg.env, "api", "http://localhost:4001", "URl to api")

	flag.Parse()

	cfg.paystack.key = os.Getenv("PAYSTACK_KEY")
	cfg.paystack.secret = os.Getenv("PAYSTACK_SECRET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dbConn, err := db.OpenDB()

	tc := make(map[string]*template.Template)

	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
	}
	err := app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}

}
