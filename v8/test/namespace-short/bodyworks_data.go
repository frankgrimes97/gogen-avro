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

// Common information related to the event which must be included in any clean event
type BodyworksData struct {
	// Unique identifier for the event used for de-duplication and tracing.
	Uuid *UnionDatatypeUUID `json:"uuid"`
	// Fully qualified name of the host that generated the event that generated the data.
	Hostname *UnionString `json:"hostname"`
	// Trace information not redundant with this object
	Trace *UnionBodyworksTrace `json:"trace"`
}

const BodyworksDataAvroCRC64Fingerprint = "\xa5\xec\x1f\xf5k\x15\xc1!"

func NewBodyworksData() BodyworksData {
	return BodyworksData{}
}

func DeserializeBodyworksData(r io.Reader) (BodyworksData, error) {
	t := NewBodyworksData()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBodyworksDataFromSchema(r io.Reader, schema string) (BodyworksData, error) {
	t := NewBodyworksData()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBodyworksData(r BodyworksData, w io.Writer) error {
	var err error
	err = writeUnionDatatypeUUID(r.Uuid, w)
	if err != nil {
		return err
	}
	err = writeUnionString(r.Hostname, w)
	if err != nil {
		return err
	}
	err = writeUnionBodyworksTrace(r.Trace, w)
	if err != nil {
		return err
	}
	return err
}

func (r BodyworksData) Serialize(w io.Writer) error {
	return writeBodyworksData(r, w)
}

func (r BodyworksData) Schema() string {
	return "{\"doc\":\"Common information related to the event which must be included in any clean event\",\"fields\":[{\"default\":null,\"doc\":\"Unique identifier for the event used for de-duplication and tracing.\",\"name\":\"uuid\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"bodyworks.datatype\",\"type\":\"record\"}]},{\"default\":null,\"doc\":\"Fully qualified name of the host that generated the event that generated the data.\",\"name\":\"hostname\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"Trace information not redundant with this object\",\"name\":\"trace\",\"type\":[\"null\",{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}]}],\"name\":\"Trace\",\"type\":\"record\"}]}],\"name\":\"bodyworks.Data\",\"type\":\"record\"}"
}

func (r BodyworksData) SchemaName() string {
	return "bodyworks.Data"
}

func (_ BodyworksData) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BodyworksData) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BodyworksData) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BodyworksData) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BodyworksData) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BodyworksData) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BodyworksData) SetString(v string)   { panic("Unsupported operation") }
func (_ BodyworksData) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BodyworksData) Get(i int) types.Field {
	switch i {
	case 0:
		r.Uuid = NewUnionDatatypeUUID()

		return r.Uuid
	case 1:
		r.Hostname = NewUnionString()

		return r.Hostname
	case 2:
		r.Trace = NewUnionBodyworksTrace()

		return r.Trace
	}
	panic("Unknown field index")
}

func (r *BodyworksData) SetDefault(i int) {
	switch i {
	case 0:
		r.Uuid = nil
		return
	case 1:
		r.Hostname = nil
		return
	case 2:
		r.Trace = nil
		return
	}
	panic("Unknown field index")
}

func (r *BodyworksData) NullField(i int) {
	switch i {
	case 0:
		r.Uuid = nil
		return
	case 1:
		r.Hostname = nil
		return
	case 2:
		r.Trace = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ BodyworksData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BodyworksData) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BodyworksData) Finalize()                        {}

func (_ BodyworksData) AvroCRC64Fingerprint() []byte {
	return []byte(BodyworksDataAvroCRC64Fingerprint)
}

func (r BodyworksData) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["uuid"], err = json.Marshal(r.Uuid)
	if err != nil {
		return nil, err
	}
	output["hostname"], err = json.Marshal(r.Hostname)
	if err != nil {
		return nil, err
	}
	output["trace"], err = json.Marshal(r.Trace)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BodyworksData) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["uuid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Uuid); err != nil {
			return err
		}
	} else {
		r.Uuid = NewUnionDatatypeUUID()

		r.Uuid = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["hostname"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Hostname); err != nil {
			return err
		}
	} else {
		r.Hostname = NewUnionString()

		r.Hostname = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["trace"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Trace); err != nil {
			return err
		}
	} else {
		r.Trace = NewUnionBodyworksTrace()

		r.Trace = nil
	}
	return nil
}
