
```shell
# 制作私钥
openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key

# 自签名公钥 ，设置有效时间365 天
openssl req -nodes -new -x509 -sha256 -days 365 -config ./cert.conf -extensions 'req_ext' -key server.key -out server.crt

# 查看证书内容
openssl x509 -in server.pem -text
```

```shell
go install google.golang.org/protobuf/cmd/protobuf@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
protoc --version
# libprotoc 3.21.7

protoc-gen-go --version
# protoc-gen-go v1.28.1
```
proto 版本会造成报错：missing mustEmbedUnimplemented*Server method； protobuf
server端添加--go-grpc_out=require_unimplemented_servers=false
```shell
protoc --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=require_unimplemented_servers=false:. \
  --go-grpc_opt=paths=source_relative \
  ./grpc/server/myserver/protoc/hi/hi.proto
```
 