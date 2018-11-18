This projects demonstrastes on how to use cockroachdb with golang using one of the famous ORM package called `gorm`.

## Development [MacOS]
1. Download and install the `go` using this [link](https://golang.org/dl/). Please select the MacOS installer.

2. Install dependency management tool called `dep` using this [link](https://golang.github.io/dep/docs/installation.html#macos).

3. Install binary compresser called upx using.
```
brew install --HEAD upx
```
4. Please make sure you have `go-swagger` installed. Installation instructions can be found from [here](https://goswagger.io/install.html).  

5. Please clone this project in **$GOPATH/src/**. This is very important step. Typically $GOPATH is your ~/go/ directory incase of linux and macOS

## Preparing cockroachdb
1. Install cockroachdb 
```
brew install cockroach
```

2. Start cockroachdb single node server
```
cockroach start --insecure --background
```
3. Open cockroach SQL CLI
```
cockroach sql --insecure
```
4. Create sm database and sm_user user.
```
create database if not exists sm; create user if not exists sm_user; grant all on database sm to sm_user;
```


## Build Project 
1. `make swagger` autogenerates the swagger code.
2. `make binary` compiles the projects and create executable
3. `make all` performs both of these steps in a single make command
4. `make docker` creates the docker image of this project 

## Adding new api endpoints.

1. Edit the yaml to add api end-point.
2. Run `make swagger` command to autogenerate swagger code from api.yaml. 
3. You should only edit `restapi/configure_sm.go` to add new endpoint, the controller code goes to `source/backend` directory.

## Run api-server
1. Do `./bin/sm-server --port 2323`. You can change host and post using flags to run it on whatever socker you want. 
2. Currently container is also set to expose 2323 port but that is also configurable from `Dockerfile`

