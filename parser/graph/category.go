package graph

const (
	categoryNode = "Category"

	createCategoryCypher = `
		MERGE (c:` + categoryNode + `{id:$ID})
		SET c.title=$Title, c.description=$Description`
)