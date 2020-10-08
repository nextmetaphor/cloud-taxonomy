package graph

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/nextmetaphor/cloud-taxonomy/model"
)

const (
	providerNode     = "Provider"
	providerRootNode = "ProviderRoot"

	createProviderCypher = `
		MERGE (p:` + providerNode + `{id:$ID})
		SET p.title=$Title, p.description=$Description`

	createProviderRootCypher = `
		CREATE (pr:` + providerRootNode + ` {id:'provider-root', title: 'Cloud Service Provider'})
		WITH pr
		MATCH (p:Provider)
		MERGE (pr)<-[:IS_A]-(p)`
)

func CreateProviderRoot(session neo4j.Session) error {
	return executeCypher(session, createProviderRootCypher, nil)
}

func CreateProvider(session neo4j.Session, p model.Provider) error {
	return executeCypher(session, createProviderCypher, p)
}
