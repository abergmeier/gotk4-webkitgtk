// Code generated by girgen. DO NOT EDIT.

package javascriptcore

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: javascriptcoregtk-4.0 webkit2gtk-4.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <jsc/jsc.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.jsc_weak_value_get_type()), F: marshalWeakValueer},
	})
}

type WeakValue struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*WeakValue)(nil)
)

func wrapWeakValue(obj *externglib.Object) *WeakValue {
	return &WeakValue{
		Object: obj,
	}
}

func marshalWeakValueer(p uintptr) (interface{}, error) {
	return wrapWeakValue(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewWeakValue: create a new CWeakValue for the JavaScript value referenced by
// value.
//
// The function takes the following parameters:
//
//    - value: CValue.
//
func NewWeakValue(value *Value) *WeakValue {
	var _arg1 *C.JSCValue     // out
	var _cret *C.JSCWeakValue // in

	_arg1 = (*C.JSCValue)(unsafe.Pointer(value.Native()))

	_cret = C.jsc_weak_value_new(_arg1)
	runtime.KeepAlive(value)

	var _weakValue *WeakValue // out

	_weakValue = wrapWeakValue(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _weakValue
}

// Value: get a CValue referencing the JavaScript value of weak_value.
func (weakValue *WeakValue) Value() *Value {
	var _arg0 *C.JSCWeakValue // out
	var _cret *C.JSCValue     // in

	_arg0 = (*C.JSCWeakValue)(unsafe.Pointer(weakValue.Native()))

	_cret = C.jsc_weak_value_get_value(_arg0)
	runtime.KeepAlive(weakValue)

	var _value *Value // out

	_value = wrapValue(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _value
}

// ConnectCleared: this signal is emitted when the JavaScript value is
// destroyed.
func (weakValue *WeakValue) ConnectCleared(f func()) externglib.SignalHandle {
	return weakValue.Connect("cleared", f)
}
