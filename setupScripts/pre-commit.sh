#!/bin/sh
echo "pre-commit hook starting"
mysqlcontainerIP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mysqldbsubstance)
pgsqlcontainerIP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' pgsqldbsubstance)
if [ -z "$mysqlcontainerIP" ] || [ -z "$pgsqlcontainerIP" ] ; then
    echo "DB Container not started"
    docker-compose up -d
    sleep 60s
    mysqlcontainerIP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mysqldbsubstance)
    pgsqlcontainerIP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' pgsqldbsubstance)
    docker exec -i mysqldbsubstance bash < $GOPATH/src/github.com/ahmedalhulaibi/substance/setupScripts/setenv_mysql.sh
    docker exec -i pgsqldbsubstance bash < $GOPATH/src/github.com/ahmedalhulaibi/substance/setupScripts/setenv_pgsql.sh
fi
export SUBSTANCE_MYSQL="root@tcp($mysqlcontainerIP:3306)/delivery"
echo $SUBSTANCE_MYSQL
export SUBSTANCE_PGSQL="postgres://travis_test:password@$pgsqlcontainerIP:5432/postgres?sslmode=disable"
echo $SUBSTANCE_PGSQL
if go test ./...; then
    echo "pre commit-hook finished."
    echo "go test return 0"
    exit 0
else
    echo "Aborting commit go test failed"
    exit 1
fi
echo "pre-commit hook finished"