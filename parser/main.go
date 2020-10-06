package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/nextmetaphor/cloud-taxonomy/graph"
)

func main() {
	configForNeo4j40 := func(conf *neo4j.Config) { conf.Encrypted = false }

	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("username", "password", ""), configForNeo4j40)
	if err != nil {
		// TODO
		fmt.Println(err)
		return
	}
	// handle driver lifetime based on your application lifetime requirements
	// driver's lifetime is usually bound by the application lifetime, which usually implies one driver instance per application
	defer driver.Close()

	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		// TODO
		fmt.Println(err)
		return
	}
	defer session.Close()

	graph.DeleteAll(session)

	//result, err := session.Run("CREATE (n:Item { id: $id, name: $name }) RETURN n.id, n.name", map[string]interface{}{
	//	"id":   1,
	//	"name": "Item 1",
	//})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//for result.Next() {
	//	fmt.Printf("Created Item with Id = '%d' and Name = '%s'\n", result.Record().GetByIndex(0).(int64), result.Record().GetByIndex(1).(string))
	//}
}