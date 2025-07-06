create the proto file
this would contain all the functions that you need and the message structs
(in the proto file if you set
option go_package = "grpc/grpc_server/chat";
this means the generated go files will have package chat)

to run it in codespaces you might need to take a bit of effort

to install protoc in codespace
PROTOC_ZIP=protoc-3.15.8-linux-x86_64.zip
curl -OL https://github.com/google/protobuf/releases/download/v3.15.8/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local include/*
rm -f $PROTOC_ZIP

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

create the directory where you want to store the output files

run the protoc command from root directory of the project

after this is done and you try to run the protoc command
protoc --go_out=./chat --go_opt=paths=source_relative \
  --go-grpc_out=./chat --go-grpc_opt=paths=source_relative \
  chat.proto
you might get the permission issue

give permission to proto using
sudo chmod +x /usr/local/bin/protoc

if you directly try to run the protoc command with sudo the issue is when you use sudo, the command runs as root and doesn't see your user's PATH.
In this case we can run it as
sudo env "PATH=$PATH" protoc --go_out=./chat --go_opt=paths=source_relative \
  --go-grpc_out=./chat --go-grpc_opt=paths=source_relative \
  chat.proto
but this is not recommended

after the proto files are generated create the server.go file this will have a listener and create a grpc connection

after this you can create a new go file in the proto folder or continue in server.go
create a new server struct that will embed UnimplementedCoffeeShopServer(this is an example it would be something starting with Unimplimented)
after this create methods where you actually impliment the functions written in proto file

try running the server it should run without any errors

create a client that would call the methods created

coffee shop example: https://www.youtube.com/watch?v=mPESsBfUKkc
