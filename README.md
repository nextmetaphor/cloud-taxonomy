# cloud-taxonomy
Project to model the Cloud Computing taxonomy in a graph representation using [yaml-graph](https://github.com/nextmetaphor/yaml-graph) definitions.

![Azure Virtual Machine](terraform-azure-vm.jpg)

## Installation
### Install `yaml-graph`
First, follow the [installation instructions](https://github.com/nextmetaphor/yaml-graph/blob/main/README.md) to build `yaml-graph`. 
Once this is complete, at the root of the `cloud-taxonomy` repository execute the following to create a `yaml-graph`
container:
```bash
docker run -it -p7474:7474 -p7687:7687 -v $(PWD)/taxonomy:/home/ymlgraph/definition -v $(PWD)/report:/home/ymlgraph/report nextmetaphor/yaml-graph
```
## Usage
All of the following commands should be executed from within the `yaml-graph` container.
### Validate the Cloud Taxonomy Definitions
```bash
yaml-graph validate -f definition/definition-format.yml
```
### Building the Cloud Taxonomy Graph
```bash
yaml-graph load
```
Once this is complete, the graph can be examined in the local browser at http://localhost:7474/browser/ using the
following cypher:
```sql
match (n) return n
```

### Building a Static HTML Report
```bash
yaml-graph report -f report/template-format.yaml -t report/output-template.gohtml > report/cloud-taxonomy.html
```

### Building a Table
```bash
 yaml-graph report -f report/table-report/fields.yaml -t report/table-report/template.gohtml > report/table-report/table-report.html
 ```

## Licence
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This project is licenced under the terms of the [Apache 2.0 License](LICENCE.md) licence.