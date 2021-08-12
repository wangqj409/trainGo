package service

import (
	"context"
	"log"
	"time"

	kafka "github.com/segmentio/kafka-go"

	pb "shopp/api/shop/v1"
)

type MqProducerService struct {
	pb.UnimplementedMqProducerServer
}

var (
	kafkaConn *kafka.Conn
	err       error
)

func NewMqProducerService() *MqProducerService {
	topic := "goods"
	partition := 0

	log.Printf("start producer\n")
	kafkaConn, err = kafka.DialLeader(context.Background(), "tcp", "localhost:9093", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	kafkaConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	go func() {
		tick1 := time.NewTicker(time.Second * 5)
		for _ = range tick1.C {
			log.Printf("produce msg\n")
			_, err = kafkaConn.WriteMessages(
				kafka.Message{Value: []byte("one!1")},
				kafka.Message{Value: []byte("two!2")},
				kafka.Message{Value: []byte("three!3")},
			)

		}
	}()

	return &MqProducerService{}
}

func (s *MqProducerService) CreateMqproducer(ctx context.Context, req *pb.CreateMqproducerRequest) (*pb.CreateMqproducerReply, error) {
	return &pb.CreateMqproducerReply{}, nil
}
func (s *MqProducerService) UpdateMqproducer(ctx context.Context, req *pb.UpdateMqproducerRequest) (*pb.UpdateMqproducerReply, error) {
	return &pb.UpdateMqproducerReply{}, nil
}
func (s *MqProducerService) DeleteMqproducer(ctx context.Context, req *pb.DeleteMqproducerRequest) (*pb.DeleteMqproducerReply, error) {
	return &pb.DeleteMqproducerReply{}, nil
}
func (s *MqProducerService) GetMqproducer(ctx context.Context, req *pb.GetMqproducerRequest) (*pb.GetMqproducerReply, error) {
	return &pb.GetMqproducerReply{}, nil
}
func (s *MqProducerService) ListMqproducer(ctx context.Context, req *pb.ListMqproducerRequest) (*pb.ListMqproducerReply, error) {
	return &pb.ListMqproducerReply{}, nil
}
