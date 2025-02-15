gin + mongo real-time chat with Tinode

Setup instructions:
# 1. Setup mongo (!IMPORTANT! Due to the requirements of Tinode it is necessary to run Mongo in single-node ReplicaSet mode: https://docs.mongodb.com/manual/administration/replica-set-deployment/)
Depends on your OS

# 1. Create .env config file:
```dotenv
PORT=<app port>
MONGO_URI=<mongo db url>
DB_NAME=<mongo db name>
JWT_SECRET=<JWT secret key to gen tokens>
JWT_EXPIRATION_MINUTES=<JWT token expiration time>
MESSAGES_PAGE_SIZE=<messages page size>
```
All variables have default values (except JWT_SECRET) defined in AppConfig

# 3. Install project dependencies
```shell
go mod tidy
```

# 4. Make tinode db connection:
```shell
$GOPATH/bin/tinode-db -config=./tinode-db/tinode.conf 
```

# 5. Launch tinode server:
```shell
$GOPATH/bin/server
```
If anything goes wrong: https://github.com/tinode/chat/blob/master/INSTALL.md#running-a-standalone-server

# 6. Start the application:
```shell
go run main.go
```

