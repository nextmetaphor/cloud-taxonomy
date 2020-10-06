package graph

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/fatih/structs"
)

const (
	deleteAllCypher = `
		MATCH (n) DETACH DELETE(n);`
)

func executeCypher(session neo4j.Session, cypher string, param interface{}) error {
	var res neo4j.Result
	var err error
	var s map[string]interface{}

	if param != nil {
		s = structs.Map(param)
	}

	if res, err = session.Run(cypher, s); err != nil {
		fmt.Println(err)

		return err
	}

	return res.Err()
}

func DeleteAll(session neo4j.Session) error {
	return executeCypher(session, deleteAllCypher, nil)
}
