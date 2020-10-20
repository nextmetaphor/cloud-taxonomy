package graph

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/nextmetaphor/cloud-taxonomy/model"
)

const (
	capabilityNode          = "Capability"
	capabilityRootNode      = "CapabilityRoot"
	capabilityRootNodeID    = "capability-root"
	capabilityRootNodeTitle = "capability"

	createCapabilityCypher = `
		MATCH (c:` + categoryNode + `{id:$CategoryID})
		MERGE (p:` + capabilityNode + `{id:$ID})
		SET p.title=$Title, p.description=$Description
		MERGE (c)<-[:TYPE_OF]-(p)`

	createCapabilityRootCypher = `
		CREATE (sr:` + capabilityRootNode + ` {id:'` + capabilityRootNodeID + `', title: '` + capabilityRootNodeTitle + `'})
		WITH sr
		MATCH (s:Capability)
		MERGE (sr)<-[:IS_A]-(s)`
)

func CreateCapabilityRoot(session neo4j.Session) error {
	return executeCypher(session, createCapabilityRootCypher, nil)
}

func CreateCapability(session neo4j.Session, s model.Capability) error {
	return executeCypher(session, createCapabilityCypher, s)
}
