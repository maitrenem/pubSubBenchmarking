RM=rm -f
RMFORCE=rm -rf 
SRCS=redisPub.go redisSub.go
COMP_NAME=
all: exe
exe: redisPub redisSub

redisPub: redisPub.go
	go build -o redisPub redisPub.go

redisSub: redisSub.go
	go build -o redisSub redisSub.go

install:
	@echo "Test Redis Pub Sub has no files to install"
clean:
	$(RM) redisPub redisSub
