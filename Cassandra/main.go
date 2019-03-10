package Cassandra

import (
	"github.com/gocql/gocql"
	"fmt"
)

// _Session for visibility to outside packages
var Session *gocql.Session

func main() {
	var err error

	// connecting to cassandra
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "winnieapi"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}
