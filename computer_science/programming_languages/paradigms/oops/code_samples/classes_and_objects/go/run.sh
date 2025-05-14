
rm go.mod go.sum
go mod init classes_and_objects_demo
go mod tidy
go build
./classes_and_objects_demo
rm classes_and_objects_demo
go mod tidy