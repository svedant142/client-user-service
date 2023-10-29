package main

import (
	pb "client-user-service/proto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Unexpected Error", err)
	}
	defer conn.Close()
	client := pb.NewUserHandlerClient(conn)

	//Testing the two endpoints
	GetSingleUser(client,0)
	GetSingleUser(client,-1)
	GetSingleUser(client,2)
	GetSingleUser(client,200000000000)
	GetSingleUser(client,2000)
	GetUserList(client, []int64{0,-1,-2})
	GetUserList(client, []int64{})
	GetUserList(client, []int64{2,1,0})
	GetUserList(client, []int64{1,2})
	GetUserList(client, []int64{1,2,16})
	GetUserList(client, []int64{24,23,25,26,27,28})


}

func GetSingleUser(client pb.UserHandlerClient,id int64) {
		getUserRequest := &pb.GetUserRequest{
			ID: id,
		}
		getUserResponse, err := client.GetUser(context.Background(), getUserRequest)
		if err != nil {
			log.Printf("Error calling GetUser: %v", err)
		} else if getUserResponse != nil {
				if message := getUserResponse.GetMessage(); message != "" {
					fmt.Printf("ID - %d , Response message-%s\n",id, message)
				} else {
					user := getUserResponse.GetUser()
					fmt.Printf("User Details: ID=%d, Fname=%s, City=%s, Height=%f, Married=%v, Phone=%d \n", user.ID, user.Fname, user.City, user.Height, user.Married, user.Phone)
				}
			} else {
				fmt.Println("empty response")
			}
}

func GetUserList(client pb.UserHandlerClient, ids []int64) {
		getUserListRequest := &pb.GetUserListRequest{
			IDs: ids, 
		}
		getUserListResponse, err := client.GetUsersByIDs(context.Background(), getUserListRequest)
		if err != nil {
			log.Printf("Error calling GetUsersByIDs: %v", err)
		} else {
			if getUserListResponse.GetSuccessListResponse() != nil {
				userList := getUserListResponse.GetSuccessListResponse().GetUsers()
				invalidIDs := getUserListResponse.GetSuccessListResponse().GetInvalidIDs()
				fmt.Println("Users:")
				for _, user := range userList {
					fmt.Printf("User Details: ID=%d, Fname=%s, City=%s, Height=%f, Married=%v, Phone=%d \n", user.ID, user.Fname, user.City, user.Height, user.Married, user.Phone)
				}
				fmt.Println("Records not found for the following IDs:", invalidIDs)
			} else if getUserListResponse.GetErrorResponse() != nil {
				errorResponse := getUserListResponse.GetErrorResponse()
				fmt.Printf("Error: %s\n", errorResponse.Error)
			}
		}
}