package main

import (
	"context"
	"log"
	"time"

	pb "grpc-student/studentpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// ── GetStudent ──────────────────────────────────────
	res, err := client.GetStudent(ctx, &pb.StudentRequest{Id: 101})
	if err != nil {
		log.Fatalf("Error calling GetStudent: %v", err)
	}

	log.Printf("=== GetStudent ===")
	log.Printf("ID: %d", res.Id)
	log.Printf("Name: %s", res.Name)
	log.Printf("Major: %s", res.Major)
	log.Printf("Email: %s", res.Email)
	log.Printf("Phone: %s", res.Phone)

	listRes, err := client.ListStudents(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error calling ListStudents: %v", err)
	}

	log.Printf("=== ListStudents ===")
	for _, s := range listRes.Student {
		log.Printf("ID: %d | Name: %s | Major: %s | Email: %s | Phone: %s", s.Id, s.Name, s.Major, s.Email, s.Phone)
	}
}
