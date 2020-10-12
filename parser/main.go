package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/nextmetaphor/cloud-taxonomy/definition"
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

	e := definition.LoadCategories(session, graph.CreateCategory)
	if e != nil {
		fmt.Println(e)
		return
	}
	e = definition.LoadProviders(session, graph.CreateProvider)
	if e != nil {
		fmt.Println(e)
		return
	}

	e = definition.LoadServices(session, graph.CreateService)
	if e != nil {
		fmt.Println(e)
		return
	}

	e = definition.LoadTenancies(session, graph.CreateTenancy)
	if e != nil {
		fmt.Println(e)
		return
	}

	e = definition.LoadAttributes(session, graph.CreateAttribute)
	if e != nil {
		fmt.Println(e)
		return
	}
}
