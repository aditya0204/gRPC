package main

import (
	"context"
	"fmt"
	numberspb "gRPC/SumAPI/numbers"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {

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
