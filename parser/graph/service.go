package graph

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/nextmetaphor/cloud-taxonomy/model"
)

const (
	serviceNode     = "Service"
	serviceRootNode = "ServiceRoot"

	createServiceCypher = `
		MATCH (c:` + categoryNode + `{id:$CategoryID})
		MERGE (p:` + serviceNode + `{id:$ID})
		SET p.title=$Title, p.description=$Description
		MERGE (c)<-[:TYPE_OF]-(p)`

	createServiceRootCypher = `
		CREATE (sr:` + serviceRootNode + ` {id:'service-root', title: 'Service'})
		WITH sr
		MATCH (s:Service)
		MERGE (sr)<-[:IS_A]-(s)`
)

func CreateServiceRoot(session neo4j.Session) error {
	return executeCypher(session, createServiceRootCypher, nil)
}

func CreateService(session neo4j.Session, s model.Service) error {
	return executeCypher(session, createServiceCypher, s)
}
