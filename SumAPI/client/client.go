package main

import (
	"context"
	"fmt"
	numberspb "gRPC/SumAPI/numbers"
	"google.golang.org/grpc"
	"io"
	"time"
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
	//doServerStreaming(c)
	//doClientStreaming(c)
	doBiDiStreaming(c)
}

func doBiDiStreaming(c numberspb.SumServiceClient) {
	arr := [][]int32{
		{0, 1, 2, 3} ,   /*  initializers for row indexed by 0 */
		{4, 52, 6, 7} ,   /*  initializers for row indexed by 1 */
		{18, 9, 10, 11} ,  /*  initializers for row indexed by 2 */
	}
	stream, err := c.MaxNumber(context.Background())
	if err!=nil{
		fmt.Println(err)
	}
	ch := make(chan bool)

	go func() {
		for _, el := range arr{
			req:=&numberspb.MaxRequest{Numbers:el}
			err=stream.Send(req)
			if err!=nil{
				fmt.Println(err)
			}
            time.Sleep(1*time.Second)
		}
	}()
	go func() {
		for{
			res, err2 := stream.Recv()
			if err==io.EOF{
				fmt.Println("End of response.")
			}

			if err2 !=nil{
				fmt.Println(err2)
			}
			fmt.Println("All response received.")
			fmt.Println("Max Number is ",res.Result)
		}
	}()





	<-ch

}

func doClientStreaming(c numberspb.SumServiceClient) {
	arr := []int32{1,2,3,4,5,6}
	stream,err:=c.ComputeAverage(context.Background())

	for _, el := range  arr{
		req := &numberspb.AverageRequest{Num:el}
		if err!=nil{
			fmt.Println(err)
		}else{
			stream.Send(req)
		}
		time.Sleep(1*time.Second)
	}

	response,err:=stream.CloseAndRecv()
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("Average is ",response.Result)
	}


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