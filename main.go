package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/go-dart-rpc/user-model/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

type server struct {
	proto.UserServiceServer
}

type UserModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string
	LastName  string
	Email     string
	Gender    string
}

func GetgRPCUserModel(data *UserModel) *proto.User {
	return &proto.User{
		Id:        data.ID.Hex(),
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Gender:    data.Gender,
	}
}

func (*server) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	fmt.Println("Create user request")
	user := req.GetUser()

	data := UserModel{
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
		Email:     user.GetEmail(),
		Gender:    user.GetGender(),
	}

	collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to Object ID (OID)"),
		)
	}

	return &proto.CreateUserResponse{
		User: &proto.User{
			Id:        oid.Hex(),
			FirstName: user.GetFirstName(),
			LastName:  user.GetLastName(),
			Email:     user.GetEmail(),
			Gender:    user.GetGender(),
		},
	}, nil

}

func (*server) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	fmt.Println("Delete user request")

	oid, err := primitive.ObjectIDFromHex(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	filter := bson.M{"_id": oid}

	res, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find user in MongoDB: %v", err),
		)
	}

	return &proto.DeleteUserResponse{UserId: req.GetUserId()}, nil
}

func (*server) ListUser(_ *proto.ListUserRequest, stream proto.UserService_ListUserServer) error {
	fmt.Println("List user request")
	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &UserModel{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)

		}
		stream.Send(&proto.ListUserResponse{User: GetgRPCUserModel(data)})
	}
	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return nil
}

func (*server) ReadUser(ctx context.Context, req *proto.ReadUserRequest) (*proto.ReadUserResponse, error) {
	fmt.Println("Read user request")

	userID := req.GetUserId()
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	data := &UserModel{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find user with specified ID: %v", err),
		)
	}

	return &proto.ReadUserResponse{
		User: GetgRPCUserModel(data),
	}, nil
}

func (*server) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	fmt.Println("Update user request")
	user := req.GetUser()

	oid, err := primitive.ObjectIDFromHex(user.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	data := &UserModel{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find user with specified ID: %v", err),
		)
	}

	data.FirstName = user.GetFirstName()
	data.LastName = user.GetFirstName()
	data.Email = user.GetEmail()
	data.Gender = user.GetGender()

	_, updateErr := collection.ReplaceOne(context.Background(), filter, data)
	if updateErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot update object in MongoDB: %v", updateErr),
		)
	}

	return &proto.UpdateUserResponse{
		User: GetgRPCUserModel(data),
	}, nil
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Connecting to MongoDB")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:admin@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User service")
	collection = client.Database("mydb").Collection("user")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	proto.RegisterUserServiceServer(s, &server{})

	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	// First we close the connection with MongoDB:
	fmt.Println("Closing MongoDB Connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error on disconnection with MongoDB : %v", err)
	}

	// Finally, we stop the server
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of Program")
}
