// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)


// A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014
  
type BodyworksDatatypeUUID struct {

	
	
		Uuid string
	

}

func NewBodyworksDatatypeUUID() (*BodyworksDatatypeUUID) {
	return &BodyworksDatatypeUUID{}
}

func DeserializeBodyworksDatatypeUUID(r io.Reader) (*BodyworksDatatypeUUID, error) {
	t := NewBodyworksDatatypeUUID()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err	
	}
	return t, err
}

func DeserializeBodyworksDatatypeUUIDFromSchema(r io.Reader, schema string) (*BodyworksDatatypeUUID, error) {
	t := NewBodyworksDatatypeUUID()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err	
	}
	return t, err
}

func writeBodyworksDatatypeUUID(r *BodyworksDatatypeUUID, w io.Writer) error {
	var err error
	
	err = vm.WriteString( r.Uuid, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *BodyworksDatatypeUUID) Serialize(w io.Writer) error {
	return writeBodyworksDatatypeUUID(r, w)
}

func (r *BodyworksDatatypeUUID) Schema() string {
	return "{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"bodyworks.datatype\",\"type\":\"record\"}"
}

func (r *BodyworksDatatypeUUID) SchemaName() string {
	return "bodyworks.datatype.UUID"
}

func (_ *BodyworksDatatypeUUID) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *BodyworksDatatypeUUID) SetInt(v int32) { panic("Unsupported operation") }
func (_ *BodyworksDatatypeUUID) SetLong(v int64) { panic("Unsupported operation") }
func (_ *BodyworksDatatypeUUID) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *BodyworksDatatypeUUID) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *BodyworksDatatypeUUID) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *BodyworksDatatypeUUID) SetString(v string) { panic("Unsupported operation") }
func (_ *BodyworksDatatypeUUID) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BodyworksDatatypeUUID) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.String)(&r.Uuid)
		
	
	}
	panic("Unknown field index")
}

func (r *BodyworksDatatypeUUID) SetDefault(i int) {
	switch (i) {
	
        
	case 0:
       	 	r.Uuid = ""
		return
	
	
	}
	panic("Unknown field index")
}

func (_ *BodyworksDatatypeUUID) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *BodyworksDatatypeUUID) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *BodyworksDatatypeUUID) Finalize() { }
