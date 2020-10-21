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
	taxonomyRoot  = "../taxonomy"
	categoryPath  = "category"
	providerPath  = "provider"
	servicePath   = "capability"
	tenancyPath   = "tenancy"
	attributePath = "attribute"
)

func loadCategory(path string) (category model.Category, err error) {
	if yamlFile, err := ioutil.ReadFile(path); err == nil {
		err = yaml.Unmarshal(yamlFile, &category)
	}

	return
}

func loadProvider(path string) (provider model.Provider, err error) {
	if yamlFile, err := ioutil.ReadFile(path); err == nil {
		err = yaml.Unmarshal(yamlFile, &provider)
	}

	return
}

func loadService(path string) (service model.Capability, err error) {
	if yamlFile, err := ioutil.ReadFile(path); err == nil {
		err = yaml.Unmarshal(yamlFile, &service)
	}

	return
}

func loadTenancy(path string) (tenancy model.Tenancy, err error) {
	if yamlFile, err := ioutil.ReadFile(path); err == nil {
		err = yaml.Unmarshal(yamlFile, &tenancy)
	}

	return
}

func loadAttribute(path string) (attribute model.Attribute, err error) {
	if yamlFile, err := ioutil.ReadFile(path); err == nil {
		err = yaml.Unmarshal(yamlFile, &attribute)
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

		p, e := loadProvider(path)
		if e != nil {
			return e
		}

		e = createNode(session, p)
		return e
	})

	if err == nil {
		err = graph.CreateProviderRoot(session)
	}

	return err
}

func LoadServices(session neo4j.Session, createNode func(session neo4j.Session, s model.Capability) error) error {
	err := filepath.Walk(taxonomyRoot+string(filepath.Separator)+servicePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		s, e := loadService(path)
		if e != nil {
			return e
		}

		e = createNode(session, s)
		return e
	})

	if err == nil {
		err = graph.CreateCapabilityRoot(session)
	}

	return err
}

func LoadTenancies(session neo4j.Session, createNode func(session neo4j.Session, t model.Tenancy) error) error {
	err := filepath.Walk(taxonomyRoot+string(filepath.Separator)+tenancyPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		s, e := loadTenancy(path)
		if e != nil {
			return e
		}

		e = createNode(session, s)
		return e
	})

	if err == nil {
		err = graph.CreateTenancyRoot(session)
	}

	return err
}

func LoadAttributes(session neo4j.Session, createNode func(session neo4j.Session, attribute model.Attribute) error) error {
	err := filepath.Walk(taxonomyRoot+string(filepath.Separator)+attributePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		s, e := loadAttribute(path)
		if e != nil {
			return e
		}

		e = createNode(session, s)
		return e
	})

	if err == nil {
		err = graph.CreateAttributeRoot(session)
	}

	return err
}