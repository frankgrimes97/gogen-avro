// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
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

type UnionIPAddressEventTypeEnum int

const (
	UnionIPAddressEventTypeEnumIPAddress UnionIPAddressEventTypeEnum = 0

	UnionIPAddressEventTypeEnumEvent UnionIPAddressEventTypeEnum = 1
)

type UnionIPAddressEvent struct {
	IPAddress IPAddress
	Event     Event
	UnionType UnionIPAddressEventTypeEnum
}

func writeUnionIPAddressEvent(r UnionIPAddressEvent, w io.Writer) error {

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionIPAddressEventTypeEnumIPAddress:
		return writeIPAddress(r.IPAddress, w)
	case UnionIPAddressEventTypeEnumEvent:
		return writeEvent(r.Event, w)
	}
	return fmt.Errorf("invalid value for UnionIPAddressEvent")
}

func NewUnionIPAddressEvent() UnionIPAddressEvent {
	return UnionIPAddressEvent{}
}

func (r UnionIPAddressEvent) Serialize(w io.Writer) error {
	return writeUnionIPAddressEvent(r, w)
}

func DeserializeUnionIPAddressEvent(r io.Reader) (UnionIPAddressEvent, error) {
	t := NewUnionIPAddressEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionIPAddressEventFromSchema(r io.Reader, schema string) (UnionIPAddressEvent, error) {
	t := NewUnionIPAddressEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r UnionIPAddressEvent) Schema() string {
	return "[{\"aliases\":[\"ip_address\"],\"name\":\"IPAddress\",\"size\":16,\"type\":\"fixed\"},{\"doc\":\"The test record\",\"fields\":[{\"doc\":\"Unique ID for this event.\",\"name\":\"id\",\"type\":\"string\"},{\"doc\":\"Start IP of this observation's IP range.\",\"name\":\"start_ip\",\"type\":\"IPAddress\"},{\"doc\":\"End IP of this observation's IP range.\",\"name\":\"end_ip\",\"type\":\"IPAddress\"}],\"name\":\"event\",\"subject\":\"event\",\"type\":\"record\"}]"
}

func (_ UnionIPAddressEvent) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ UnionIPAddressEvent) SetInt(v int32)      { panic("Unsupported operation") }
func (_ UnionIPAddressEvent) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ UnionIPAddressEvent) SetDouble(v float64) { panic("Unsupported operation") }
func (_ UnionIPAddressEvent) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ UnionIPAddressEvent) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionIPAddressEvent) SetLong(v int64) {

	r.UnionType = (UnionIPAddressEventTypeEnum)(v)
}

func (r *UnionIPAddressEvent) Get(i int) types.Field {

	switch i {
	case 0:
		return &IPAddressWrapper{Target: (&r.IPAddress)}
	case 1:
		r.Event = NewEvent()
		return &types.Record{Target: (&r.Event)}
	}
	panic("Unknown field index")
}
func (_ UnionIPAddressEvent) NullField(i int)                  { panic("Unsupported operation") }
func (_ UnionIPAddressEvent) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ UnionIPAddressEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ UnionIPAddressEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ UnionIPAddressEvent) Finalize()                        {}

func (r UnionIPAddressEvent) MarshalJSON() ([]byte, error) {

	switch r.UnionType {
	case UnionIPAddressEventTypeEnumIPAddress:
		return json.Marshal(map[string]interface{}{"IPAddress": r.IPAddress})
	case UnionIPAddressEventTypeEnumEvent:
		return json.Marshal(map[string]interface{}{"event": r.Event})
	}
	return nil, fmt.Errorf("invalid value for UnionIPAddressEvent")
}

func (r *UnionIPAddressEvent) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["IPAddress"]; ok {
		r.UnionType = 0
		return json.Unmarshal([]byte(value), &r.IPAddress)
	}
	if value, ok := fields["event"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Event)
	}
	return fmt.Errorf("invalid value for UnionIPAddressEvent")
}
