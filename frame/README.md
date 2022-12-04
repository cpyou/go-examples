
上传文件
```shell
curl -X POST http://localhost:8080/upload \
  -F "file=@/Users/cpy/Downloads/test.zip" \
  -H "Content-Type: multipart/form-data"
```

上传多个文件
```shell
curl -X POST http://localhost:8080/multi-upload \
  -F "upload[]=@/Users/cpy/Downloads/test1.zip" \
  -F "upload[]=@/Users/cpy/Downloads/test2.zip" \
  -H "Content-Type: multipart/form-data"
```