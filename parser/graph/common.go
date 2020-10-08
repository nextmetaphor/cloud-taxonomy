package graph

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/neo4j/neo4j-go-driver/neo4j"
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

//func CreateCategories(session neo4j.Session, cm map[string]model.Category) (err error) {
//	for _, e := range cm {
//		if err = CreateCategory(session, e); err != nil {
//			return err
//		}
//	}
//
//	err = createCategoryRoot(session)
//	return err
//}
