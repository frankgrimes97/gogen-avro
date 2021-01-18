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

// Trace
type HeaderworksTrace struct {
	// Trace Identifier
	TraceId *UnionDatatypeUUID `json:"traceId"`
}

const HeaderworksTraceAvroCRC64Fingerprint = "\x8a\xdfu\xe7˻\xa6\xbc"

func NewHeaderworksTrace() HeaderworksTrace {
	return HeaderworksTrace{}
}

func DeserializeHeaderworksTrace(r io.Reader) (HeaderworksTrace, error) {
	t := NewHeaderworksTrace()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeHeaderworksTraceFromSchema(r io.Reader, schema string) (HeaderworksTrace, error) {
	t := NewHeaderworksTrace()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeHeaderworksTrace(r HeaderworksTrace, w io.Writer) error {
	var err error
	err = writeUnionDatatypeUUID(r.TraceId, w)
	if err != nil {
		return err
	}
	return err
}

func (r HeaderworksTrace) Serialize(w io.Writer) error {
	return writeHeaderworksTrace(r, w)
}

func (r HeaderworksTrace) Schema() string {
	return "{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}]}],\"name\":\"headerworks.Trace\",\"type\":\"record\"}"
}

func (r HeaderworksTrace) SchemaName() string {
	return "headerworks.Trace"
}

func (_ HeaderworksTrace) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ HeaderworksTrace) SetInt(v int32)       { panic("Unsupported operation") }
func (_ HeaderworksTrace) SetLong(v int64)      { panic("Unsupported operation") }
func (_ HeaderworksTrace) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ HeaderworksTrace) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ HeaderworksTrace) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ HeaderworksTrace) SetString(v string)   { panic("Unsupported operation") }
func (_ HeaderworksTrace) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *HeaderworksTrace) Get(i int) types.Field {
	switch i {
	case 0:
		r.TraceId = NewUnionDatatypeUUID()

		return r.TraceId
	}
	panic("Unknown field index")
}

func (r *HeaderworksTrace) SetDefault(i int) {
	switch i {
	case 0:
		r.TraceId = nil
		return
	}
	panic("Unknown field index")
}

func (r *HeaderworksTrace) NullField(i int) {
	switch i {
	case 0:
		r.TraceId = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ HeaderworksTrace) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ HeaderworksTrace) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ HeaderworksTrace) Finalize()                        {}

func (_ HeaderworksTrace) AvroCRC64Fingerprint() []byte {
	return []byte(HeaderworksTraceAvroCRC64Fingerprint)
}

func (r HeaderworksTrace) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traceId"], err = json.Marshal(r.TraceId)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *HeaderworksTrace) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["traceId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TraceId); err != nil {
			return err
		}
	} else {
		r.TraceId = NewUnionDatatypeUUID()

		r.TraceId = nil
	}
	return nil
}
