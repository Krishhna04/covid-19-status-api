#Understanding the project structure 
###The following are the purposes of the directories in the codebase
|
> This document assume that you have already installed  go version 1.15.7 or above

# Setting up the project
### Clone the repo
  
git clone https://github.com/Krishhna04/covid-19-status-api.git


### Copy .env.example to .env. 

cp .env.example .env

### Make necessary changes as below
|Key|Purpose|
|-|-|
|APP_PORT|The port to run the service|
|APP_ENVIRONMENT|The environment to run the application. Possible values are: `production`, `development`|


# Setting up and starting the API server for the first time
This API is built using docker container. Hence, to start the service, docker image + container needs to be built from the project root. The command is:
## For development environment:
```
$ sudo docker-compose -f docker-compose-dev.yml up -d --build
```
Here, once the container is built, it listens for changes made to the `.go` files. As soon as the change is detected, the code is compiled and the binary is run.

## For production environment:
```
$ sudo docker-compose up -d --build
```
The production environment is a multi-stage docker build where the final container just contains the executable binary.


## Performing tasks on an already existing API server

### Installing a new package
Since this is a dockerized environment, typically a new go package has to be added from inside the running docker container. Hence, on the dev container, you need to :
```
$ docker exec -it dl_ocr bash
```
This would take you inside the docker container. Then you simply need to issue the command(s). e.g.:
```
$ go get -u github.com/gomodule/redigo/redis
```
Once done, the package would be a part of go.mod (on the host machine as well). You should then commit and push the code. Having done it at the dev level, the production container build system will take care of this installation automatically at the right time. 

You see, some old school development workflow is really helpful.

# Understanding the project structure
### The following are the purposes of the directories in the codebase
|Directory  |Purpose|
|-|-|
|configs  |This should contain all project specific configurations. Common examples are : setting up database connection, cache connection and so on |
|controllers  |This should contain all of the controllers defining the endpoints for the service in question. In short, this is the entry point of the incoming HTTP request to the system. The business logic should be written in this directory in such a way that the business logic remains in a completely different file. We need to ensure unit testability of the business logic|
|docs  |This should contain the API documentation. Preferred API Documentation tool is Swagger. GoLang specific plugin is [swaggo](https://github.com/swaggo/swag) |
|helpers  |This should contain the commonly reused code. Examples are: input validation formulae, error logging methods and so on|
|models  |This should contain ALL the database interactions for the project. In some languages, this gets called as Data Access Layer|
|services |This contains business logic implemented in the API  |
### The following are the purposes of the project root level files in the codebase
|File|Purpose|
|-|-|
|routes.go|This lists down ALL the available routes for the service. It can also contain the actions that needs to be commonly taken across multiple routes (middleware)|
|main.go|This is the file which is responsible for running the application server. This should be attached to a monitoring service e.g. [supervisord](http://supervisord.org/) / [mmonit](https://mmonit.com/) to ensure that it recovers from a shutdown|On a typical docker-compose based environment, this would be taken care of by the docker engine|This also lists down ALL the available routes for the service. It can also contain the actions that needs to be commonly taken across multiple routes (middleware)|
|go.sum|This file Lists down the checksum of direct and indirect dependency required by the module|
|covid_status.postman_collection.json| This files contains the Postman Collection of the API,Corresponding environment is to be replaced with the URL variable .|
|README.md|This file which contains the API documentation|