// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     union.avsc
 */
package avro

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
)

func writeIp_address(r Ip_address, w io.Writer) error {
	_, err := w.Write(r[:])
	return err
}

type Ip_address Ip_addressWrapper
type Ip_addressWrapper [16]byte

func (_ *Ip_addressWrapper) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Ip_addressWrapper) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Ip_addressWrapper) SetLong(v int64) { panic("Unsupported operation") }
func (_ *Ip_addressWrapper) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Ip_addressWrapper) SetDouble(v float64) { panic("Unsupported operation") }
func (r *Ip_addressWrapper) SetBytes(v []byte) { 
	copy((*r)[:], v)
}
func (_ *Ip_addressWrapper) SetString(v string) { panic("Unsupported operation") }
func (_ *Ip_addressWrapper) SetUnionElem(v int64) { panic("Unsupported operation") }
func (_ *Ip_addressWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *Ip_addressWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Ip_addressWrapper) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Ip_addressWrapper) Finalize() { }
func (_ *Ip_addressWrapper) SetDefault(i int) { panic("Unsupported operation") }

