package main

import (
	"context"
	"fmt"
	numberspb "gRPC/SumAPI/numbers"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type server struct {

}

func getMax(arr []int32) int32{
	n:=len(arr)
	var res int32
	res=0
	for i:=0 ; i<n;i++{
		if arr[i]>res{
			res=arr[i]
		}
	}

	return res
}


func (s *server) MaxNumber(stream numberspb.SumService_MaxNumberServer) error {
   fmt.Println("Request received in Max Number api...")

	for  {
		req,err:=stream.Recv()
		if err==io.EOF{
			fmt.Println("Request ended.")
			return nil
		}

		if err!=nil{
			return err
		}

		numbers := req.GetNumbers()
		mx:= getMax(numbers)

		response:=&numberspb.MaxResponse{Result:mx}
		stream.Send(response)
		fmt.Println("Response sent.")
	}
}

func (s *server) ComputeAverage(stream numberspb.SumService_ComputeAverageServer) error {
	fmt.Println("Compute average request received.")
    var sum int32
    var cnt int32
	for {
		req ,err := stream.Recv()
		if err==io.EOF{
			res := &numberspb.AverageResponse{Result:sum/cnt}
			stream.SendAndClose(res)
			fmt.Println("Response sent.")
			return nil
		}
		if err!=nil{
			fmt.Println(err)
		}else{
			n:= req.GetNum()
			fmt.Println("Number received is ",n)
			sum = sum+ n
			cnt++
		}
	}

}

func (s *server) Prime(req *numberspb.PrimeRequest, stream numberspb.SumService_PrimeServer) error {
   num:= req.GetNum()
   fmt.Println("Prime request received with number ",num)
   var k int32
   k=2

   for {
   	  if num<=1{
   	  	break
	  }
   	  if num%k ==int32(0){
   	  	  res :=&numberspb.PrimeResponse{Num:k}
   	  	  time.Sleep(500*time.Millisecond)
   	  	  stream.Send(res)
   	  	  num= num/k
	  }else {
	  	k++
	  }
   }
	return nil
}

func (s *server) Numbers(ctx context.Context, req *numberspb.NumRequest) (*numberspb.NumResponse, error) {
	n1:=req.GetNum1()
	n2:= req.GetNum2()
	fmt.Println("Received request with numbers ",n1," , ",n2)
	sum := n1+n2
    return &numberspb.NumResponse{Sum:sum}, nil
}

func main(){
	fmt.Println("Welcome to gRPC server...")

	l,err:=net.Listen("tcp","localhost:5000")
	if err!=nil{
		fmt.Println(err)
		return
	}

	s:= grpc.NewServer()

	numberspb.RegisterSumServiceServer(s,&server{})

   if err=s.Serve(l); err!=nil{
   	log.Fatal(err)
   }

}
