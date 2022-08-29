// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

// GoGen test
type ComAvroTestSample struct {
	// Core data information required for any event
	Header *HeaderworksData `json:"header"`
	// Core data information required for any event
	Body *BodyworksData `json:"body"`
}

const ComAvroTestSampleAvroCRC64Fingerprint = "\xdf}\x93 \x19f\x18\n"

func NewComAvroTestSample() ComAvroTestSample {
	r := ComAvroTestSample{}
	return r
}

func DeserializeComAvroTestSample(r io.Reader) (ComAvroTestSample, error) {
	t := NewComAvroTestSample()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeComAvroTestSampleFromSchema(r io.Reader, schema string) (ComAvroTestSample, error) {
	t := NewComAvroTestSample()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeComAvroTestSample(r ComAvroTestSample, w io.Writer) error {
	var err error
	if r.Header == nil {
		err = vm.WriteLong(0, w)
		if err != nil {
			return err
		}
	} else {
		err = vm.WriteLong(int64(1), w)
		if err != nil {
			return err
		}

		err = writeHeaderworksData(*r.Header, w)
	}
	if r.Body == nil {
		err = vm.WriteLong(0, w)
		if err != nil {
			return err
		}
	} else {
		err = vm.WriteLong(int64(1), w)
		if err != nil {
			return err
		}

		err = writeBodyworksData(*r.Body, w)
	}
	return err
}

func (r ComAvroTestSample) Serialize(w io.Writer) error {
	return writeComAvroTestSample(r, w)
}

func (r ComAvroTestSample) Schema() string {
	return "{\"doc\":\"GoGen test\",\"fields\":[{\"default\":null,\"doc\":\"Core data information required for any event\",\"name\":\"header\",\"type\":[\"null\",{\"doc\":\"Common information related to the event which must be included in any clean event\",\"fields\":[{\"default\":null,\"doc\":\"Unique identifier for the event used for de-duplication and tracing.\",\"name\":\"uuid\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}]},{\"default\":null,\"doc\":\"Fully qualified name of the host that generated the event that generated the data.\",\"name\":\"hostname\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"Trace information not redundant with this object\",\"name\":\"trace\",\"type\":[\"null\",{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",\"headerworks.datatype.UUID\"]}],\"name\":\"Trace\",\"type\":\"record\"}]}],\"name\":\"Data\",\"namespace\":\"headerworks\",\"type\":\"record\"}]},{\"default\":null,\"doc\":\"Core data information required for any event\",\"name\":\"body\",\"type\":[\"null\",{\"doc\":\"Common information related to the event which must be included in any clean event\",\"fields\":[{\"default\":null,\"doc\":\"Unique identifier for the event used for de-duplication and tracing.\",\"name\":\"uuid\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"bodyworks.datatype\",\"type\":\"record\"}]},{\"default\":null,\"doc\":\"Fully qualified name of the host that generated the event that generated the data.\",\"name\":\"hostname\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"Trace information not redundant with this object\",\"name\":\"trace\",\"type\":[\"null\",{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",\"headerworks.datatype.UUID\"]}],\"name\":\"Trace\",\"type\":\"record\"}]}],\"name\":\"Data\",\"namespace\":\"bodyworks\",\"type\":\"record\"}]}],\"name\":\"com.avro.test.sample\",\"type\":\"record\"}"
}

func (r ComAvroTestSample) SchemaName() string {
	return "com.avro.test.sample"
}

func (_ ComAvroTestSample) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ComAvroTestSample) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ComAvroTestSample) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ComAvroTestSample) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ComAvroTestSample) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ComAvroTestSample) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ComAvroTestSample) SetString(v string)   { panic("Unsupported operation") }
func (_ ComAvroTestSample) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ComAvroTestSample) Get(i int) types.Field {
	switch i {
	case 0:
		if r.Header == nil {
			var Header = new(HeaderworksData)
			r.Header = Header
		}
		w := r.Header

		return w

	case 1:
		if r.Body == nil {
			var Body = new(BodyworksData)
			r.Body = Body
		}
		w := r.Body

		return w

	}
	panic("Unknown field index")
}

func (r *ComAvroTestSample) SetDefault(i int) {
	switch i {
	case 0:
		r.Header = nil
		return
	case 1:
		r.Body = nil
		return
	}
	panic("Unknown field index")
}

func (r *ComAvroTestSample) NullField(i int) {
	switch i {
	case 0:
		r.Header = nil
		return
	case 1:
		r.Body = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ComAvroTestSample) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ComAvroTestSample) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ComAvroTestSample) HintSize(int)                     { panic("Unsupported operation") }
func (_ ComAvroTestSample) Finalize()                        {}

func (_ ComAvroTestSample) AvroCRC64Fingerprint() []byte {
	return []byte(ComAvroTestSampleAvroCRC64Fingerprint)
}

func (r ComAvroTestSample) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	if r.Header == nil {
		output["header"], err = []byte("null"), nil
	} else {
		output["header"], err = json.Marshal(map[string]interface{}{
			"headerworks.Data": r.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	if r.Body == nil {
		output["body"], err = []byte("null"), nil
	} else {
		output["body"], err = json.Marshal(map[string]interface{}{
			"bodyworks.Data": r.Body,
		})
	}
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ComAvroTestSample) UnmarshalheaderJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}

	if v, ok := fields["headerworks.Data"]; ok {
		r.Header = new(HeaderworksData)
		json.Unmarshal(v, r.Header)
	}

	return nil
}
func (r *ComAvroTestSample) UnmarshalbodyJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}

	if v, ok := fields["bodyworks.Data"]; ok {
		r.Body = new(BodyworksData)
		json.Unmarshal(v, r.Body)
	}

	return nil
}

func (r *ComAvroTestSample) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["header"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := r.UnmarshalheaderJSON(val); err != nil {
			return err
		}
	} else {
		r.Header = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["body"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := r.UnmarshalbodyJSON(val); err != nil {
			return err
		}
	} else {
		r.Body = nil
	}
	return nil
}
