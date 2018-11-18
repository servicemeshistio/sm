package utils

import (
	"fmt"
	"sm/source/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	addr = "postgresql://sm_user@localhost:26257/sm?sslmode=disable"

	//ServerInfo is the cockroachdb connection info object
	ServerInfo *Server
)

func init() {
	db := setupDB(addr)
	// defer db.Close()
	ServerInfo = NewServer(db)
}

// Server is an http server that handles REST requests.
type Server struct {
	Db *gorm.DB
}

func setupDB(addr string) *gorm.DB {
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	// Migrate the schema
	db.AutoMigrate(&models.PatientObject{})
	return db
}

// NewServer creates a new instance of a Server.
func NewServer(db *gorm.DB) *Server {
	return &Server{Db: db}
}
