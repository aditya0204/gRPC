package main

import (
	"context"
	"fmt"
	numberspb "gRPC/SumAPI/numbers"
	"google.golang.org/grpc"
	"io"
)


func main(){
	fmt.Println("Welcome to client...")

   cc,err:=	grpc.Dial("localhost:5000",grpc.WithInsecure())
   if err!=nil{
   	fmt.Println(err)
	   return
   }

	c:= numberspb.NewSumServiceClient(cc)

	//doUnary(c)
	doServerStreaming(c)
}

func doServerStreaming(c numberspb.SumServiceClient) {
	fmt.Println("Starting a server streaming rpc request...")
	var num1 int32

	fmt.Println("Enter number ...")
	fmt.Scanf("%d",&num1)

	req := &numberspb.PrimeRequest{Num:num1}

	res ,err := c.Prime(context.Background(),req)
	if err!=nil{
		fmt.Println(err)
		return
	}

    for {
    	msg,err:= res.Recv()
    	if err==io.EOF{
    		fmt.Println("End of result.")
    		break
		}

		if err!=nil{
			fmt.Println(err)
		}

		fmt.Println(msg.GetNum())
	}
}

func doUnary(c numberspb.SumServiceClient){
   fmt.Println("Starting a unary rpc request...")
   var num1,num2 int32

   fmt.Println("Enter 1st number ...")
   fmt.Scanf("%d",&num1)

   fmt.Println("Enter 2nd number ...")
   fmt.Scanf("%d",&num2)

	req := &numberspb.NumRequest{
	   Num1: num1,
	   Num2: num2,
    }

   res,err:= c.Numbers(context.Background(),req)
	if err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println("result of sum is ",res.Sum)


}