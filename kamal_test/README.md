# Build and run with docker

docker build -t kamal_test .

docker run -d -p 8080:8080 --name kamal_test kamal_test
