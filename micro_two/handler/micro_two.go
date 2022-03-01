package handler

import (
	"context"
	"io"
	"time"

	log "go-micro.dev/v4/logger"

	pb "micro_two/proto"
)

type MicroTwo struct {
	microTwoHandler pb.MicroTwoHandler
}

func NewMicroTwo() *MicroTwo {
	return &MicroTwo{}
}

func (e *MicroTwo) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received Micro_two.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *MicroTwo) ClientStream(ctx context.Context, stream pb.MicroTwo_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *MicroTwo) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.MicroTwo_ClientStreamStream) error {
	log.Infof("Received Micro_two.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		log.Infof("Sending %d", i)
		if err := stream.SendMsg(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *MicroTwo) BidiStream(ctx context.Context, stream pb.MicroTwo_ClientStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.SendMsg(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
