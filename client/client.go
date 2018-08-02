package client

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	p "proto"
)

// TODO: custom codec && ? || any type
//type Codec struct{}
//
//func (c *Codec) Marshal(i interface{}) ([]byte, error) {
//	return nil, nil
//}
//
//func (c *Codec) Unmarshal(data []byte, i interface{}) error {
//	return nil
//}
//
//func (c *Codec) String() string {
//	return "CustomCodec"
//}

type Client struct {
	base p.KvStoreClient
}

func (c *Client) Init(conn *grpc.ClientConn) {
	c.base = p.NewKvStoreClient(conn)
}

func (c *Client) Get(key string) ([]byte, error) {
	resp, err := c.base.Get(context.Background(), &p.Request{Key: key})
	if err != nil {
		return nil, err
	}
	return resp.Value, nil
}

func (c *Client) Put(key string, value []byte) error {
	_, err := c.base.Put(context.Background(), &p.PutRequest{Key: key, Value: value})
	return err
}

func (c *Client) Delete(key string) error {
	_, err := c.base.Delete(context.Background(), &p.Request{Key: key})
	return err
}
