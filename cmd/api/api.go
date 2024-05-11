package main

import (
	"biomassx/internal/api"
	"biomassx/internal/member"
	"os"

	"git.doolta.com/doolta/go-kit/gormdb"
	"git.doolta.com/doolta/go-kit/log"
	"git.doolta.com/doolta/go-kit/web"
	"github.com/joho/godotenv"
)

// @title           BiomassX API
// @version         1.0
// @description     An API for BiomassX
// @contact.name    BiomassX team
// @contact.email   dev@biomassx.com
// @host            localhost:12345
// @BasePath        /v1
func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	l := log.GetLogger("debug")
	db, err := gormdb.GetPostgresqlDB(l, os.Getenv("DBHOST"), os.Getenv("DBNAME"), os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBPORT"))
	if err != nil {
		panic("Can't connect to database: " + err.Error())
	}

	memberService := member.NewService(db)
	// fs := http.FileServer(http.Dir("./static"))

	s := web.NewHTTPServer(l)
	s.AddService("member", memberService)
	s.Routes("")
	api.Routes(s)
	// s.Router.Handle("/", http.StripPrefix("/static", fs))
	l.Info("Starting server on port 12345")
	s.Run(":12345")
}
