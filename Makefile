swagger:
	bash ./build/build_swagger.sh
binary:
	bash ./build/build_binary.sh
docker:
	bash ./build/build_docker.sh

all: swagger binary
