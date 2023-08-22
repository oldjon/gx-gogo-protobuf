## 原项目地址
基于 https://github.com/gogo/protobuf
版本v1.3.2

## 修改及新增
1. 修改grpc client的Invoke调用方式。由于项目内使用的grpc版本与gogo-protobuf支持的grpc版本，存在一定版本差距，调用写法存在差异，故对grpc的调用方式做了对应兼容修改。
2. 针对数据库存储用的proto message，支持解析注释内容"scan"来确定是否生成Scan及Value函数。改善原来需要手动在model.go文件内对proto message额外包装pbmessage结构并手动实现Scan和Value函数的写法。同时修改Scan和Value函数内序列化和反序列的方式，以提高序列化和反序列化的性能。
   
   使用：
   ```
   message MsgExample {//scan
        ...
   }
   ```
   MsgExample将会生成Scan和Value函数
3. 针对数据库model，支持通过proto定义，并支持通过解析注释内容"dbtag"，来确定是否生成字段的dbtag。

   使用：
   ```
    message MsgExample {//dbtag
      uint32 field_a = 1;
   }
   ```
   MsgExample.FieldA将会生成`db:"field_a"`的tag

5. 在生成的go代码中，"id"将会被生成为"ID"，即不再生成"Id"

## 工具生成

    make

将生成的proto-gen-gofast放到GOBIN下或者PATH下

## 生成pb.go文件

    protoc --gofast_out=plugins=grpc:./ ./*.proto

## 注意
gogo-protobuf工具会为message生成Size和Value函数，所以message内的字段名请不要用size或者value，否则会产生命名冲突。
