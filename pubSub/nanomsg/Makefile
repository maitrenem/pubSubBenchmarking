RM=rm -f
RMFORCE=rm -rf 
GOLDFLAGS=-r .
all: exe
exe: nanoMsgPub nanoMsgSub

nanoMsgPub: nanoMsgPub.go
	go build -o nanoMsgPub -ldflags="$(GOLDFLAGS)" nanoMsgPub.go

nanoMsgSub: nanoMsgSub.go
	go build -o nanoMsgSub -ldflags="$(GOLDFLAGS)" nanoMsgSub.go

install:
	@echo "Test Nanomsg Pub Sub has no files to install"
clean:
	$(RM) nanoMsgPub nanoMsgSub
