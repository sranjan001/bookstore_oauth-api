package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	//connect to cassandra cluster:
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}

	fmt.Println("Connection to cassandra established ...")
}

func GetSession() *gocql.Session {
	return session
}
