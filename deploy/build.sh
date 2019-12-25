mkdir dist

go build -o dist/auth ../apps/auth/auth.go

go build -o dist/data ../apps/data/data.go


docker build -f Dockerfile-auth -t mengwei2275/datasource-auth:v1 .
docker build -f Dockerfile-data -t mengwei2275/datasource-data:v1 .

docker push mengwei2275/datasource-auth:v1
docker push mengwei2275/datasource-data:v1
docker push mengwei2275/datasource-data:v1