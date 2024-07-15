protoc --twirp_out=. --go_out=. *.proto;
go mod tidy;