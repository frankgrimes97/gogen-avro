// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v8/compiler"
	"github.com/actgardner/gogen-avro/v8/vm"
	"github.com/actgardner/gogen-avro/v8/vm/types"
)

var _ = fmt.Printf

type NestedMap struct {
	MapOfMaps map[string]map[string][]string `json:"MapOfMaps"`
}

const NestedMapAvroCRC64Fingerprint = "\xa1\x9e\x89\xd6\xc52@\xf2"

func NewNestedMap() NestedMap {
	return NestedMap{}
}

func DeserializeNestedMap(r io.Reader) (NestedMap, error) {
	t := NewNestedMap()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNestedMapFromSchema(r io.Reader, schema string) (NestedMap, error) {
	t := NewNestedMap()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNestedMap(r NestedMap, w io.Writer) error {
	var err error
	err = writeMapMapArrayString(r.MapOfMaps, w)
	if err != nil {
		return err
	}
	return err
}

func (r NestedMap) Serialize(w io.Writer) error {
	return writeNestedMap(r, w)
}

func (r NestedMap) Schema() string {
	return "{\"fields\":[{\"name\":\"MapOfMaps\",\"type\":{\"type\":\"map\",\"values\":{\"type\":\"map\",\"values\":{\"items\":\"string\",\"type\":\"array\"}}}}],\"name\":\"NestedMap\",\"type\":\"record\"}"
}

func (r NestedMap) SchemaName() string {
	return "NestedMap"
}

func (_ NestedMap) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NestedMap) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NestedMap) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NestedMap) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NestedMap) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NestedMap) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NestedMap) SetString(v string)   { panic("Unsupported operation") }
func (_ NestedMap) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NestedMap) Get(i int) types.Field {
	switch i {
	case 0:
		r.MapOfMaps = make(map[string]map[string][]string)

		return &MapMapArrayStringWrapper{Target: &r.MapOfMaps}
	}
	panic("Unknown field index")
}

func (r *NestedMap) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NestedMap) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ NestedMap) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NestedMap) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NestedMap) Finalize()                        {}

func (_ NestedMap) AvroCRC64Fingerprint() []byte {
	return []byte(NestedMapAvroCRC64Fingerprint)
}

func (r NestedMap) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["MapOfMaps"], err = json.Marshal(r.MapOfMaps)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NestedMap) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["MapOfMaps"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MapOfMaps); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for MapOfMaps")
	}
	return nil
}
