# Environment Setup

Install [docker](https://docs.docker.com/install/) and [docker-compose](https://docs.docker.com/compose/install/)

Run the local hooks setup script if you want your code tested locally after each commit. The commits will be aborted if the `go test ./...` fails. This script will also start the containers if required and export the required environment variable:
```bash
. setupScripts/setup-local-hooks.sh
```

Run [start-sql-containers.sh](https://github.com/ahmedalhulaibi/substance/blob/feature/gqlgen/setupScripts/setup-local-hooks.sh). This script will start postgres and mysql containers described in [docker-compose.yml](https://github.com/ahmedalhulaibi/substance/blob/feature/gqlgen/docker-compose.yml) and setup two environment variables
```bash
$ . setupScripts/start-sql-containers.sh
$ echo $SUBSTANCE_MYSQL 
root@tcp(172.19.0.3:3306)/delivery
$ echo $SUBSTANCE_PGSQL
postgres://travis_test:password@172.19.0.2:5432/postgres?sslmode=disable
```

