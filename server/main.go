package main

import (
	"flag"
	"fmt"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	p "proto"
	"runtime"
	"storage"
)

type kvServer struct {
	storage storage.Storage
}

func NewKvServer() *kvServer {
	if runtime.NumCPU() > 4 {
		return &kvServer{storage: &storage.SyncMap{}}
	}
	return &kvServer{storage: &storage.RWLockMap{}}
}

func (k *kvServer) Put(ctx context.Context, req *p.PutRequest) (*p.Empty, error) {
	resp := &p.Empty{}
	k.storage.Put(req.Key, req.Value)
	return resp, nil
}

func (k *kvServer) Get(ctx context.Context, req *p.Request) (resp *p.Response, err error) {
	resp = &p.Response{}
	resp.Value, err = k.storage.Get(req.Key)
	return resp, err
}

func (k *kvServer) Delete(ctx context.Context, req *p.Request) (*p.Empty, error) {
	resp := &p.Empty{}
	err := k.storage.Delete(req.Key)
	return resp, err
}

func main() {
	port := flag.Int("port", 8080, "")
	flag.Parse()

	ls, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listener started port: %d", *port)
	grpcServer := grpc.NewServer()
	srv := NewKvServer()
	p.RegisterKvStoreServer(grpcServer, srv)
	grpcServer.Serve(ls)
}
