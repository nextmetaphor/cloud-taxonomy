package graph

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/nextmetaphor/cloud-taxonomy/model"
)

const (
	attributeNode     = "Attribute"
	attributeRootNode = "AttributeRoot"

	createAttributeCypher = `
		MERGE (c:` + attributeNode + `{id:$ID})
		SET c.title=$Title, c.description=$Description`

	createAttributeRootCypher = `
		CREATE (ar:` + attributeRootNode + ` {id:'attribute-root', title: 'Attribute'})
		WITH ar
		MATCH (a:Attribute)
		MERGE (ar)<-[:IS_A]-(a)`
)

func CreateAttributeRoot(session neo4j.Session) error {
	return executeCypher(session, createAttributeRootCypher, nil)
}

func CreateAttribute(session neo4j.Session, c model.Attribute) error {
	return executeCypher(session, createAttributeCypher, c)
}
