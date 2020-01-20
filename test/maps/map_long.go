// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     maps.avsc
 */
package avro

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
)

func writeMapLong(r *MapLong, w io.Writer) error {
	err := vm.WriteLong(int64(len(r.M)), w)
	if err != nil || len(r.M) == 0 {
		return err
	}
	for k, e := range r.M {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = vm.WriteLong(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapLong struct {
	keys []string
	values []int64
	M map[string]int64
}

func NewMapLong() *MapLong{
	return &MapLong {
		keys: make([]string, 0),
		values: make([]int64, 0),
		M: make(map[string]int64),
	}
}

func (_ *MapLong) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *MapLong) SetInt(v int32) { panic("Unsupported operation") }
func (_ *MapLong) SetLong(v int64) { panic("Unsupported operation") }
func (_ *MapLong) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *MapLong) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *MapLong) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *MapLong) SetString(v string) { panic("Unsupported operation") }
func (_ *MapLong) SetUnionElem(v int64) { panic("Unsupported operation") }
func (_ *MapLong) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapLong) SetDefault(i int) { panic("Unsupported operation") }
func (r *MapLong) Finalize() { 
	for i := range r.keys {
		r.M[r.keys[i]] = r.values[i]
	}
	r.keys = nil
	r.values = nil
}

func (r *MapLong) AppendMap(key string) types.Field { 
	r.keys = append(r.keys, key)
	var v int64
	
	r.values = append(r.values, v)
	
	return (*types.Long)(&r.values[len(r.values)-1])
	
}

func (_ *MapLong) AppendArray() types.Field { panic("Unsupported operation") }

