# grpc-student-6609650251
grpc-student-6609650251

Task1:Add ListStudent RPC

1.แก้ไข proto/student.proto

เพิ่ม RPC ใน service และ เพิ่ม messages:

    service StudentService {

      rpc GetStudent (StudentRequest) returns (StudentResponse);
  
      rpc ListStudents (Empty) returns (StudentListResponse); // เพิ่ม
  
    }

    message Empty {}

    message StudentListResponse {

      repeated StudentResponse student = 1;
  
    }

2.Regenerate code

    protoc --go_out=. --go-grpc_out=. proto/student.proto

3.แก้ server/server.go

    func (s *server) ListStudents(ctx context.Context, req *pb.Empty) (*pb.StudentListResponse, error) {

     students := []*pb.StudentResponse{
    
        {Id: 1, Name: "Alice Johnson", Major: "Computer Science", Email: "alice@university.com"},
        
        {Id: 2, Name: "Bob Smith",     Major: "Mathematics",      Email: "bob@university.com"},
        
        {Id: 3, Name: "Charlie Brown", Major: "Physics",          Email: "charlie@university.com"},
        
    }
    
    return &pb.StudentListResponse{Student: students}, nil
    
}

4.แก้ client/client.go

เพิ่มการเรียก ListStudents:

    listRes, err := client.ListStudents(ctx, &pb.Empty{})

    if err != nil {

    log.Fatalf("Error calling ListStudents: %v", err)
    
    }

    log.Printf("=== ListStudents ===")

    for _, s := range listRes.Student {

    log.Printf("ID: %d | Name: %s | Major: %s | Email: %s", s.Id, s.Name, s.Major, s.Email)
    
    }

Task2: Add Phone Field

1.แก้ proto/student.proto

เพิ่ม field phone ใน StudentResponse:

    message StudentResponse {

      int32  id    = 1;
  
      string name  = 2;
  
      string major = 3;
      
      string email = 4;
      
      string phone = 5; // เพิ่ม
  
    }

2.Regenerate code

    protoc --go_out=. --go-grpc_out=. proto/student.proto

3.แก้ server/server.go

เพิ่ม Phone ในทุก response:

// GetStudent

    return &pb.StudentResponse{

    Id:    req.Id,
    
    Name:  "Alice Johnson",
    
    Major: "Computer Science",
    
    Email: "alice@university.com",
    
    Phone: "081-111-1111", // เพิ่ม
    
    }, nil

// ListStudents

    {Id: 1, Name: "Alice Johnson", Major: "Computer Science", Email: "alice@university.com", Phone: "087-111-1111"},

    {Id: 2, Name: "Bob Smith",     Major: "Mathematics",      Email: "bob@university.com",   Phone: "089-111-2222"},

    {Id: 3, Name: "Charlie Brown", Major: "Physics",          Email: "charlie@university.com", Phone: "088-111-3333"},


4.แก้ client/client.go

เพิ่ม print Phone:

    // GetStudent

    log.Printf("Phone: %s", res.Phone)

    // ListStudents


    log.Printf("ID: %d | Name: %s | Major: %s | Email: %s | Phone: %s",

    s.Id, s.Name, s.Major, s.Email, s.Phone)
  
แล้วก็รันเช็คโดย

Terminal 1 -server

    go run server/server.go

เปิดอีก terminal

Terminal 2 - client

    go run client/client.go

ตัวอย่างหลังจากใช้  go run client/client.go

PS D:\grpc-student-6609650251\grpc-student-6609650251\grpc-student> go run client/client.go

2026/03/30 14:10:33 === GetStudent ===

2026/03/30 14:10:33 ID: 101

2026/03/30 14:10:33 Name: Alice Johnson

2026/03/30 14:10:33 Major: Computer Science

2026/03/30 14:10:33 Email: alice@university.com

2026/03/30 14:10:33 Phone: 087-111-1111

2026/03/30 14:10:33 === ListStudents ===

2026/03/30 14:10:33 ID: 1 | Name: Alice Johnson | Major: Computer Science | Email: alice@university.com | Phone: 087-111-1111

2026/03/30 14:10:33 ID: 2 | Name: Bob Smith | Major: Mathematics | Email: bob@university.com | Phone: 089-111-2222

2026/03/30 14:10:33 ID: 3 | Name: Charlie Brown | Major: Physics | Email: charlie@university.com | Phone: 088-111-3333
