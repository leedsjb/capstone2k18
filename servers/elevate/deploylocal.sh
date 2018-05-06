docker rm -f redis
docker rm -f mongo
docker run -d -p 6379:6379 --name redis redis
docker run -d -p 27017:27017 --name mongo mongo

go install && gateway