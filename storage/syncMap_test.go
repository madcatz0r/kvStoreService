package storage

import (
	"testing"
	"bytes"
)

func TestSyncMap_Put(t *testing.T) {
	test := &SyncMap{}

	test.Put("testKey", []byte{1})
	testValue, ok := test.Load("testKey")
	if !ok {
		t.Error("put test failed, testKey not found")
	}
	testCastValue, _ := testValue.([]byte)
	if !bytes.Equal(testCastValue, []byte{1}) {
		t.Errorf("put test failed, expected: []byte{0x1}, got: %#v", testCastValue)
	}
}

func TestSyncMap_Get(t *testing.T) {
	test := &SyncMap{}

	test.Store("testKey", []byte{1})
	res, err := test.Get("testKey")
	if err != nil || !bytes.Equal(res, []byte{1}){
		t.Errorf("get test failed, expected: []byte{0x1} got: %#v", res)
	}

	res, err = test.Get("testMissedKey")
	if err == nil || err.Error() != "key testMissedKey not found" {
		t.Errorf("get missed key test failed, expected: key testMissedKey not found, got: %v", err.Error())
	}

	type InvalidStruct struct {
		name string
	}

	test.Store("testInvalidValue", &InvalidStruct{name: "cause cast error"})
	res, err = test.Get("testInvalidValue")
	if err == nil || err.Error() != "unexpected cast error, key: testInvalidValue, value: &storage.InvalidStruct{name:\"cause cast error\"}" {
		t.Errorf("get missed key test failed, got: %v", err.Error())
	}
}

func TestSyncMap_Delete(t *testing.T) {
	test := &SyncMap{}

	test.Store("testKey", []byte{1})
	err := test.Delete("testKey")
	if err != nil {
		t.Errorf("delete test failed: %v", err)
	}
	err = test.Delete("testKey")
	if err.Error() != "key testKey not found" {
		t.Errorf("%v", err)
	}
}
