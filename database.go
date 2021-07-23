package gobase

// import (
// 	"context"
// 	"log"

// 	"github.com/facebook/ent/examples/start/ent"
// 	_ "github.com/mattn/go-sqlite3"
// )

// // InitDatabaseConnection : init the database connection
// func InitDatabaseConnection() {
// 	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
// 	if err != nil {
// 		log.Fatalf("failed opening connection to sqlite: %v", err)
// 	}
// 	defer client.Close()
// 	// Run the auto migration tool.
// 	if err := client.Schema.Create(context.Background()); err != nil {
// 		log.Fatalf("failed creating schema resources: %v", err)
// 	}
// }
