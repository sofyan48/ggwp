# GGWP TEST

## Getting Started
This support For go version 1.13, ggwp uncle Bob Models

### Environment Setup
For Development Mode Setup dotenv
```
cp .env.example .env
```
Setting up your local configuration see example
```
####################################################################
# SERVER CONFIGURATION
####################################################################
SERVER_ADDRESS=0.0.0.0
SERVER_PORT=3000
SERVER_TIMEZONE=Asia/Jakarta
SECRET_KEY=

####################################################################
# DATABASE CONFIGURATION
####################################################################
DB_MYSQL_USERNAME=root
DB_MYSQL_PASSWORD=password
DB_MYSQL_HOST=localhost
DB_MYSQL_PORT=3306
DB_MYSQL_DATABASE=db


####################################################################
# REDIS CONFIGURATION
####################################################################
REDIS_HOST=localhost
REDIS_PORT=6379

####################################################################
# KAFKA CONFIG
####################################################################
KAFKA_PORT=9092
KAFKA_HOST=localhost
KAFKA_USERNAME=
KAFKA_PASSWORD=

####################################################################
# SWAGGER CONFIGURATION
####################################################################
SWAGGER_SERVER_ADDRESS=http://localhost:3000
```

### Local Development
Fork this repo for your repo then clone in your local
```
git clone https://github.com/sofyan48/boilerplate.git
```
Run Local compose for dependency
```bash
docker-compose -f docker-compose-local.yml
```
Starting Worker 
```
go run worker/main.go
```
Starting GGWP
```
go run src/main.go
```

Building GGWP
```
go build src/main.go
```
Run Building
```
./main
```

#### Live Reload
To activate Live Reload install air 
##### on macOS

```
curl -fLo /usr/local/bin/air \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/darwin/air
chmod +x /usr/local/bin/air
```

##### on Linux

```
curl -fLo /usr/local/bin/air \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air
chmod +x /usr/local/bin/air
```

##### on Windows

```
curl -fLo ~/air.exe \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/windows/air.exe
```

see watcher.conf setting watcher file for air config now

##### Starting Live Reload
go to your project path
```
air -c watcher.conf
```

### Production Mode
#### Building Worker Images
```bash
cd worker && docker build -t ggwp_worker . && cd ..
```
#### Deploy

Building Image

Edit Environment In docker-compose.yml then Run Compose
```
docker-compose up
```
Stop Container
```
docker-compose stop
```
Remove your container
```
docker-compose rm -f
```

## Database Migration
After setup local or production setup database now
### Golang Migrate
Documentation Mode 
[Release Downloads](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)

#### Installing
##### MAC
```
brew install golang-migrate
```
##### Linux And Windows
```
curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz
```
### Migrating Database
```
migrate -path src/migration/mysql/ -database 'mysql://root:root@tcp(localhost:3306)/boiler_db' up
```
***in this boilerplate migration path : src/migration/mysql/***

## Tree
Tree Project
```
.
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── buildspecs
│   └── docker.sh
├── docker-compose-local.yml
├── docker-compose.yml
├── docs
│   ├── insomnia
│   │   └── api.json
│   └── swagger
│       └── docs
│           ├── docs.go
│           ├── swagger.json
│           └── swagger.yaml
├── go.mod
├── go.sum
├── src
│   ├── config
│   │   └── server_configuration.go
│   ├── controller
│   │   └── v1
│   │       ├── gql
│   │       │   └── GraphQlController.go
│   │       ├── health
│   │       │   └── V1HealthController.go
│   │       ├── producer
│   │       │   └── ProducerController.go
│   │       ├── routes.go
│   │       └── user
│   │           └── UserController.go
│   ├── entity
│   │   └── v1
│   │       ├── db
│   │       │   └── user.go
│   │       ├── graphql
│   │       │   └── user.go
│   │       ├── http
│   │       │   ├── health.go
│   │       │   ├── producer.go
│   │       │   └── user.go
│   │       └── realtime
│   │           └── producer.go
│   ├── main.go
│   ├── migration
│   │   └── mysql
│   │       ├── 20191029093439_users.down.sql
│   │       └── 20191029093439_users.up.sql
│   ├── repository
│   │   └── v1
│   │       └── user
│   │           └── user.go
│   ├── resolver
│   │   └── v1
│   │       └── resolver.go
│   ├── routes
│   │   ├── route.go
│   │   └── router_test.go
│   ├── schema
│   │   └── v1
│   │       └── schema.go
│   ├── service
│   │   └── v1
│   │       ├── graphql
│   │       │   └── user.go
│   │       ├── health
│   │       │   └── health.go
│   │       ├── producer
│   │       │   └── producer.go
│   │       └── user
│   │           └── user.go
│   └── util
│       ├── helper
│       │   ├── crypto
│       │   │   ├── crypto.go
│       │   │   └── crypto_test.go
│       │   ├── mysqlconnection
│       │   │   └── connection.go
│       │   ├── redis
│       │   │   ├── redis.go
│       │   │   └── redis_test.go
│       │   ├── rest
│       │   │   ├── rest.go
│       │   │   └── rest_test.go
│       │   └── str_process
│       │       └── str_process.go
│       ├── kafka
│       │   └── kafka.go
│       └── middleware
│           ├── auth.go
│           ├── cors_middleware.go
│           └── middleware.go
├── watcher.conf
└── worker
    ├── Dockerfile
    ├── consumer
    │   └── consumer.go
    ├── go.mod
    ├── go.sum
    └── main.go
```

## How To Understand this project
See the list below for instructions on how to use this boilerplate:

1. ***Entity*** are a form of modeling of the REST API and Query GraphQL that you will create
2. ***Repositories*** Is a connection model that you will use in a service that will be created later, the entity plays an important role for modeling data that will be stored in a database
3. ***Service*** Is a logic process that you will create services can also be combined with business logic to facilitate you in managing all services so pay attention to the layering version of the service
4. ***Controller*** an intermediary between routing and service controller regulates all input and input formats of the REST API that you will create

6. ***Resolver*** Root Resolver lingking graphql
7. ***Schema*** Root Schema public graphql
8. ***Routing*** Routing is a mapping of the paths of the REST API that you have designed, routing is available on each controller version layering, this routing will be called on the router that has been created

Plugins and utils are in the ***util*** folder all third-party packages that help you should be stored in this folder, you can choose whether the package is a middleware of your REST API or as a pure supporting utility

This ***Worker*** Runnning Worker for realtime Execution



## Documentation Format
### Setup Swagger Docs
See Documentation 
[Swag Docs](https://github.com/swaggo/swag)
### Local Docs
Local Swagger
[Local](http://localhost:3000/swagger/index.html)
### Insomnia
import insomia REST file in ***docs/insomnia/api.json***
### Curl
#### Producer
```bash
curl --request POST \
  --url http://localhost:3000/v1/producer \
  --header 'content-type: application/json' \
  --data '{
	"target":"order",
	"data" : {
		"name": "sofyan",
		"order_code": "123456"
	}
}'
```
![producer](https://github.com/sofyan48/ggwp/blob/master/docs/images/producer.png)
***Consumer***
![consumer](https://github.com/sofyan48/ggwp/blob/master/docs/images/consumer.png)


#### Insert Data
```bash
curl --request POST \
  --url http://localhost:3000/v1/user \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data 'name=iank 2' \
  --data email=iank_stc@yahoo.com \
  --data password=password \
  --data date_of_birth=1992-01-02 \
  --data phone_number=081247930699 \
  --data current_address=Current \
  --data 'city=Jakarta Barat' \
  --data province=Jakarta \
  --data 'district=jakarta barat' \
  --data lat=0 \
  --data lng=0 \
  --data job=Jobs \
  --data image=https://image.com
```
![Insert](https://github.com/sofyan48/ggwp/blob/master/docs/images/insert.png)

#### Update User
```bash
curl --request PUT \
  --url http://localhost:3000/v1/user/4 \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data 'name=iank 6' \
  --data email=iank_stc@yahoo.com \
  --data password=password \
  --data date_of_birth=1992-01-02 \
  --data phone_number=081247930699 \
  --data current_address=Current \
  --data 'city=Jakarta Barat' \
  --data province=Jakarta \
  --data 'district=jakarta barat' \
  --data lat=0 \
  --data lng=0 \
  --data job=Jobs \
  --data image=https://image.com
```
![Update](https://github.com/sofyan48/ggwp/blob/master/docs/images/update.png)

#### Get All User -> graphql
```bash
curl --request POST \
  --url http://localhost:3000/v1/graphql \
  --header 'content-type: application/json' \
  --data '{"query":"query {\n  Users(limit:10, page:0){\n      page\n    results {\n      id\n      name\n      city\n    }\n  }\n}"}'
```
![all_user](https://github.com/sofyan48/ggwp/blob/master/docs/images/all_user.png)

#### Get User By ID -> Graphql
```bash
curl --request POST \
  --url http://localhost:3000/v1/graphql \
  --header 'content-type: application/json' \
  --data '{"query":"query {\n  User(id:1){\n     id\n     name\n     city\n  }\n}"}'
```
![user_id](https://github.com/sofyan48/ggwp/blob/master/docs/images/id_user.png)


## How To Contribute
Please refer to each project's style and contribution guidelines for submitting patches and additions. In general, we follow the "fork-and-pull" Git workflow.
 1. ***Fork*** the repo on GitHub
 2. ***Clone*** the project to your own machine
 3. ***Commit*** changes to your own branch
 4. ***Push*** your work back up to your fork
 5. Submit a ***Pull request*** so that we can review your changes
