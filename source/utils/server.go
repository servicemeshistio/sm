package utils

import (
	"flag"
	"fmt"
	"sm/source/utils/dbschema"

	"github.com/jinzhu/gorm"
)

var (
	addr = flag.String("addr", "postgresql://sm_user@localhost:26257/sm?sslmode=disable", "the address of the database")
	//ServerInfo is the cockroachdb connection info object
	ServerInfo *Server
)

func init() {
	flag.Parse()
	db := setupDB(*addr)
	defer db.Close()
	ServerInfo = NewServer(db)
}

// Server is an http server that handles REST requests.
type Server struct {
	db *gorm.DB
}

func setupDB(addr string) *gorm.DB {
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	// Migrate the schema
	db.AutoMigrate(&dbschema.Patient{})
	return db
}

// NewServer creates a new instance of a Server.
func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}
