package graph

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/nextmetaphor/cloud-taxonomy/model"
)

const (
	tenancyNode     = "Tenancy"
	tenancyRootNode = "TenancyRoot"

	createTenancyCypher = `
		MERGE (c:` + tenancyNode + `{id:$ID})
		SET c.title=$Title, c.description=$Description`

	createTenancyRootCypher = `
		CREATE (tr:` + tenancyRootNode + ` {id:'tenancy-root', title: 'Tenancy'})
		WITH tr
		MATCH (t:` + tenancyNode + `)
		MERGE (tr)<-[:IS_A]-(t)`
)

func CreateTenancyRoot(session neo4j.Session) error {
	return executeCypher(session, createTenancyRootCypher, nil)
}

func CreateTenancy(session neo4j.Session, t model.Tenancy) error {
	return executeCypher(session, createTenancyCypher, t)
}
