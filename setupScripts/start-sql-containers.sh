#!bin/bash
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
msqlcontainerPort=$(docker inspect --format='{{(index (index .NetworkSettings.Ports "3306/tcp") 0).HostPort}}' mysqldbsubstance)
export SUBSTANCE_MYSQL="root@tcp($mysqlcontainerIP:$msqlcontainerPort)/delivery"
echo $SUBSTANCE_MYSQL
pgqlcontainerPort=$(docker inspect --format='{{(index (index .NetworkSettings.Ports "5432/tcp") 0).HostPort}}' pgsqldbsubstance)
export SUBSTANCE_PGSQL="postgres://travis_test:password@$pgsqlcontainerIP:$pgqlcontainerPort/postgres?sslmode=disable"
echo $SUBSTANCE_PGSQL