// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

// ValueHashInsertValue inserts value into hash. (Unlike with
// g_hash_table_insert(), both the key and the value are copied).
//
// Deprecated: Use #GVariant API instead.
func ValueHashInsertValue(hash *glib.HashTable, key string, value *externglib.Value) {
	var _arg1 *C.GHashTable // out
	var _arg2 *C.char       // out
	var _arg3 *C.GValue     // out

	_arg1 = (*C.GHashTable)(gextras.StructNative(unsafe.Pointer(hash)))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(key)))
	_arg3 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	C.soup_value_hash_insert_value(_arg1, _arg2, _arg3)
}

// NewValueHash creates a Table whose keys are strings and whose values are
// #GValue.
//
// Deprecated: Use #GVariant API instead.
func NewValueHash() *glib.HashTable {
	var _cret *C.GHashTable // in

	_cret = C.soup_value_hash_new()

	var _hashTable *glib.HashTable // out

	_hashTable = (*glib.HashTable)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_hashTable, func(v *glib.HashTable) {
		C.free(gextras.StructNative(unsafe.Pointer(v)))
	})

	return _hashTable
}
