# go-echo-mongo
Simple Golang REST application with Echo Framework & MongoDB

# How to run
These are the steps to run this app: 
1. Make sure Golang,Glide(Go Package Manager), and MongoDB are installed
2. Clone this repository to your dir, eg. ```$GOPATH/src/github.com/<your_dir>```
3. Go to project root directory (```$GOPATH/src/github.com/<your_dir>/go-echo-mongo```)
4. Populate the config file ```./config/config.example.json``` with your own configuration and copy to ```./config/config.json```
5. Run command ```glide install``` to install the dependencies
6. Start the app with command ```go run main.go```

# Application

>The request header should contain:
```Content-Type: "application/json"```

>The error response should be:
```json
{
  "error":true,
  "message":"ERROR MESSAGE"
}
```

>The success response should be:
```json
{
  "data": "DATA<could be string/object/array>",
  "error":false,
  "message":"ERROR MESSAGE"
}
```

These are the list of endpoint:

Method       | URI              | Description
------------ | ---------------- | -------------
POST         | /users           | Create new user.
GET          | /users/<:userID> | Get user by ID.
GET          | /users           | Get list of user.

# References
- [YouTube] (https://www.youtube.com/playlist?list=PLMrwI6jIZn-3a4Hjn-GoYihbMBAzZ6Ae3)
- [GitHub] (https://github.com/bxcodec/go-clean-arch)
