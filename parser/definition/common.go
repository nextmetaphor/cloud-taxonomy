package definition

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/nextmetaphor/cloud-taxonomy/graph"
	"github.com/nextmetaphor/cloud-taxonomy/model"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	taxonomyRoot = "../taxonomy"
	categoryPath = "category"
	providerPath = "provider"
)

func loadCategory(path string) (c model.Category, err error) {
	if yamlFile, err := ioutil.ReadFile(path); err == nil {
		err = yaml.Unmarshal(yamlFile, &c)
	}

	return
}

func loadProvider(path string) (p model.Provider, err error) {
	if yamlFile, err := ioutil.ReadFile(path); err == nil {
		err = yaml.Unmarshal(yamlFile, &p)
	}

	return
}

func LoadCategories(session neo4j.Session, createNode func(session neo4j.Session, c model.Category) error) error {
	err := filepath.Walk(taxonomyRoot+string(filepath.Separator)+categoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		c, e := loadCategory(path)
		if e != nil {
			return e
		}

		e = createNode(session, c)
		return e
	})

	if err == nil {
		err = graph.CreateCategoryRoot(session)
	}

	return err
}

func LoadProviders(session neo4j.Session, createNode func(session neo4j.Session, p model.Provider) error) error {
	err := filepath.Walk(taxonomyRoot+string(filepath.Separator)+providerPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		c, e := loadProvider(path)
		if e != nil {
			return e
		}

		e = createNode(session, c)
		return e
	})

	if err == nil {
		err = graph.CreateProviderRoot(session)
	}

	return err
}
