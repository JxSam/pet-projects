ROOT_DIR=$(shell pwd)

united_projects:
	cd golang/united_projects/main && go run main.go

file_reader:
	cd golang/file_reader && go run main.go