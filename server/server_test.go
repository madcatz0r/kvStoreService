package main_test

import (
	"bytes"
	c "github.com/madcatz0r/kvStoreService/client"
	"encoding/gob"
	"fmt"
	"google.golang.org/grpc"
	"testing"
)

const (
	ip   = "127.0.0.1"
	port = 8080
)

var (
	testStruct = &TypeTest{Field1: 256, Field2: 255}
)

type TypeTest struct {
	Field1 int
	Field2 byte
}

func newConnection() (conn *grpc.ClientConn, err error) {
	conn, err = grpc.Dial(fmt.Sprintf("%v:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func TestKvServer_Put(t *testing.T) {
	conn, err := newConnection()
	if err != nil {
		t.Fatalf("Unable to create connection to server: %v", err)
	}
	defer conn.Close()

	cl := &c.Client{}
	cl.Init(conn)

	testStruct := &TypeTest{Field1: 256, Field2: 255}
	gob.Register(TypeTest{})
	testBytes, _ := Encode(testStruct)
	err = cl.Put("testStruct", testBytes)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

}

func TestKvServer_Get(t *testing.T) {
	conn, err := newConnection()
	if err != nil {
		t.Fatalf("Unable to create connection to server: %v", err)
	}
	defer conn.Close()

	cl := &c.Client{}
	cl.Init(conn)
	resp, err := cl.Get("testStruct")
	if err != nil {
		t.Fatal(err)
	}

	result, err := Decode(resp)
	compareStruct, ok := result.(TypeTest)
	if !ok {
		t.Errorf("cast failed")
	}

	if testStruct.Field1 != compareStruct.Field1 {
		t.Errorf("expected: %v, got: %v", testStruct.Field1, compareStruct.Field1)
	}
	if testStruct.Field2 != compareStruct.Field2 {
		t.Errorf("expected: %v, got: %v", testStruct.Field2, compareStruct.Field2)
	}

}

func TestKvServer_Delete(t *testing.T) {
	conn, err := newConnection()
	if err != nil {
		t.Fatalf("Unable to create connection to server: %v", err)
	}
	defer conn.Close()

	cl := &c.Client{}
	cl.Init(conn)

	err = cl.Delete("testStruct")
	if err != nil {
		t.Fatal(err)
	}
	_, err = cl.Get("testStruct")
	if err != nil {
		t.Log(err)
		return
	}
	t.Fatal("delete failed")
}

func Encode(val interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(&val)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(b []byte) (val interface{}, err error) {
	var buf bytes.Buffer
	buf.Write(b)
	decoder := gob.NewDecoder(&buf)
	err = decoder.Decode(&val)
	return val, err
}
