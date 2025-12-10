package shared_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alaa-aqeel/looply-app/src/shared"
)

func TestOptional(t *testing.T) {
	o := shared.SetValue("hello")
	if !reflect.DeepEqual(o, shared.Optional[string]{Value: "hello", IsSet: true}) {
		t.Fatal("Failed to create optional value")
	}

	o = shared.SetValueWithDefault("hello", "world")
	if !reflect.DeepEqual(o, shared.Optional[string]{Value: "hello", IsSet: true, Default: "world"}) {
		t.Fatal("Failed to create optional value with default value")
	}

	o = shared.SetValueWithDefault("hello", "")
	if !reflect.DeepEqual(o, shared.Optional[string]{Value: "hello", IsSet: true, Default: ""}) {
		t.Fatal("Failed to create optional value with default value")
	}
}

func TestOptionalWithJson(t *testing.T) {
	jsonInput := []byte(`{"name": "test", "active": true}`)
	var args struct {
		Name   shared.Optional[string] `json:"name"`
		Active shared.Optional[bool]   `json:"active"`
	}
	json.Unmarshal(jsonInput, &args)

	if !reflect.DeepEqual(args.Name, shared.SetValue("test")) {
		t.Fatal("Failed to unmarshal optional value")
	}
	if !reflect.DeepEqual(args.Active, shared.SetValue(true)) {
		t.Fatal("Failed to unmarshal optional value")
	}
}

func TestOptionalGetDefaultValue(t *testing.T) {
	o := shared.SetValueWithDefault[int64](0, 10)

	if !reflect.DeepEqual(o, shared.Optional[int64]{Value: 0, IsSet: true, Default: 10}) {
		t.Fatal("Failed to create optional value with default value")
	}

	jsonInput := []byte(`{"name": "hello world"}`)
	var args struct {
		Name   shared.Optional[string] `json:"name"`
		Active shared.Optional[bool]   `json:"active"`
		Age    shared.Optional[int64]  `json:"age"`
	}
	json.Unmarshal(jsonInput, &args)

	if v := args.Name.ValueOrDefault("hi world"); v != "hello world" {
		t.Fatal("Failed to unmarshal optional value : " + v)
	}

	if args.Active.ValueOrDefault(true) != true {
		t.Fatal("Failed to unmarshal optional value")
	}

	if args.Age.ValueOrDefault(20) != 20 {
		t.Fatal("Failed to unmarshal optional value")
	}
}
