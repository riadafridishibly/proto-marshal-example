generate:
	cd protoexample && protoc --go_out=paths=source_relative:. -I. example.proto