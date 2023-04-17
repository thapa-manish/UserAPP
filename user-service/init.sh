if [ ! -f env/.env ]
then
 set -o allexport; source env/.env; set +o allexport
fi
go build -o user-api cmd/main.go 
./user-api
 