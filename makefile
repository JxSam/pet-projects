ROOT_DIR=$(shell pwd)

united_projects:
	cd golang/united_projects/main && go run main.go

read_file:
	cd golang/file_reader/read_file && go run main.go

read_string:
	cd golang/file_reader/read_string && go run main.go

write_file:
	cd golang/file_reader/write_file && go run main.go