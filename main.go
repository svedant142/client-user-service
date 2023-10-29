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
	//Records available by mocking database for userIDs -  
	//  1,2,3,4,5,6,7,8,9,10,200000000000
	
	//Testing the two endpoints
	fmt.Println("Endpoint to fetch user details based on user id :")
	fmt.Println("Sample Case 1 - ")
	GetSingleUser(client,0)
	fmt.Println("Sample Case 2 - ")
	GetSingleUser(client,-1)
	fmt.Println("Sample Case 3 - ")
	GetSingleUser(client,2)
	fmt.Println("Sample Case 4 - ")
	GetSingleUser(client,200000000000)
	fmt.Println("Sample Case 5 - ")
	GetSingleUser(client,2000)
	fmt.Println("\nEndpoint to fetch a list of user details based on a list of ids :")
	fmt.Println("Sample Case 1 - ")
	GetUserList(client, []int64{0,-1,-2})
	fmt.Println("Sample Case 2 - ")
	GetUserList(client, []int64{})
	fmt.Println("Sample Case 3 - ")
	GetUserList(client, []int64{2,1,0})
	fmt.Println("Sample Case 4 - ")
	GetUserList(client, []int64{1,2})
	fmt.Println("Sample Case 5 - ")
	GetUserList(client, []int64{1,2,16})
	fmt.Println("Sample Case 6 - ")
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