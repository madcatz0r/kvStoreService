package storage

import (
	"testing"
	"reflect"
	"bytes"
)

func TestNewRWLockMap(t *testing.T) {
	test := NewRWLockMap()
	testType := reflect.TypeOf(test.internal)
	if testType.String() != "map[string]interface {}" {
		t.Errorf("unexpected strorage type: %#v", test.internal)
	}
	if test.internal == nil {
		t.Error("internal storage is missed")
	}
}

func TestRWLockMap_Put(t *testing.T) {
	test := NewRWLockMap()

	test.Put("testKey", []byte{1})
	testValue := test.internal["testKey"]
	testCastValue, _ := testValue.([]byte)
	if !bytes.Equal(testCastValue, []byte{1}) {
		t.Errorf("put test failed, expected: []byte{0x1}, got: %#v", testCastValue)
	}
}

func TestRWLockMap_Get(t *testing.T) {
	test := NewRWLockMap()

	test.internal["testKey"] = []byte{1}
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

	test.internal["testInvalidValue"] = &InvalidStruct{name: "cause cast error"}
	res, err = test.Get("testInvalidValue")
	if err == nil || err.Error() != "unexpected cast error, key: testInvalidValue, value: &storage.InvalidStruct{name:\"cause cast error\"}" {
		t.Errorf("get missed key test failed, got: %v", err.Error())
	}
}

func TestRWLockMap_Delete(t *testing.T) {
	test := NewRWLockMap()

	test.internal["testKey"] = []byte{1}
	err := test.Delete("testKey")
	if err != nil {
		t.Errorf("delete test failed: %v", err)
	}
	err = test.Delete("testKey")
	if err.Error() != "key testKey not found" {
		t.Errorf("%v", err)
	}
}