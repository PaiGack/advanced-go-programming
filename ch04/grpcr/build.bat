@REM 
@REM build *.proto
@REM 
@REM 
@REM 
@REM p1 指定源 proto 文件和依赖的 proto 文件的目录路径（通过 --
@REM proto_path 或 -I 命令行标记来指定）。如果不指定该值，则将
@REM 使用当前目录作为源目录。在这个目录下，需要根据包名来存放依
@REM 赖的 proto 文件

@REM p2 指定希望编译的 proto 文件路径。编译器将阅读该文件并生成输
@REM 出的 Go 文件。

@REM p3 指定生成的代码要存放的目标目录
@REM 
@REM 
@REM 


cd proto

set p1=.
set p2=hello.proto
set p3=../..
set p4=%GOPATH%/src

protoc -I %p4% -I %p1% %p2% --go-grpc_out=%p3% --go_out=%p3%

cd ..


go build ./service
go build ./client