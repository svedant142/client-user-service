# client-user-service

This is the client side sample logic for the user-grpc-service. The service can be found - https://github.com/svedant142/user-grpc-service

NOTE - This is only sample client for testing purpose, the same gRPC requests can be hit using Postman. Follow Readme in the user-grpc-service repository.

It consists of a call to -
● An endpoint to fetch user details based on user id.
● An endpoint to fetch a list of user details based on a list of ids.

Sample cases are already added.

HOW TO RUN -

1. run the following command on terminal - go run main.go
   it will start the server

2. If using vscode or any IDE -
   run the main.go file in debug mode.
