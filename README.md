# cloud-taxonomy
Project to model Cloud Computing Taxonomy definitions within a local Docker graph database.

## Installation
### Start a Local Graph Database
First, from the root of the cloned repository, get the latest version of `apoc-x.x.x.x-all.jar
` from https://github.com/neo4j-contrib/neo4j-apoc-procedures/releases and place into `NEO4J_HOME/plugins`.

Secondly, again from the root of the cloned repository, run the following command to start the local graph Docker container
```bash
docker run \
    --publish=7474:7474 \
    --publish=7687:7687 \
    --volume=`pwd`/NEO4J_HOME/data:/data \
    --volume=`pwd`/NEO4J_HOME/plugins:/plugins \
    --volume=`pwd`:/import \
    --env=NEO4J_AUTH=none \
    --env=NEO4J_apoc_export_file_enabled=true \
    --env=NEO4J_apoc_import_file_enabled=true \
    --env=NEO4J_apoc_import_file_use__neo4j__config=true \
    --env=NEO4J_dbms_security_procedures_unrestricted=apoc.\\\* \
    neo4j
```

## Usage
### Building the Project
```bash
go build -i
./data-graph
```

## Licence
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL%202.0-brightgreen.svg)](https://opensource.org/licenses/MPL-2.0)

This project is licenced under the terms of the [Mozilla Public License Version 2.0](LICENCE.md) licence.