package graph

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/nextmetaphor/cloud-taxonomy/model"
)

const (
	categoryNode     = "Category"
	categoryRootNode = "CategoryRoot"

	createCategoryCypher = `
		MERGE (c:` + categoryNode + `{id:$ID})
		SET c.title=$Title, c.description=$Description`

	createCategoryRootCypher = `
		CREATE (cr:` + categoryRootNode + ` {id:'category-root', title: 'Category Root Node'})
		WITH cr
		MATCH (c:Category)
		MERGE (cr)<-[:IS_A]-(c)`
)

func createCategoryRoot(session neo4j.Session) error {
	return executeCypher(session, createCategoryRootCypher, nil)
}

func createCategory(session neo4j.Session, c model.Category) error {
	return executeCypher(session, createCategoryCypher, c)
}
