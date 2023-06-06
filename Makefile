nginx-start:
	cd nginx
	docker build --tag my-nginx .
	docker run --detach --publish 8080:80 --name nginx my-nginx

go-build:
	docker build --tag kvs .
	docker run --detach --publish 8080:8080 kvs

protoc-gen:
	protoc --go_out=paths=source_relative:. keyvalue.proto