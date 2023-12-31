// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <webkit/webkit.h>
import "C"

//export _gotk4_webkit6_BackForwardList_ConnectChanged
func _gotk4_webkit6_BackForwardList_ConnectChanged(arg0 C.gpointer, arg1 *C.WebKitBackForwardListItem, arg2 C.gpointer, arg3 C.guintptr) {
	var f func(itemAdded *BackForwardListItem, itemsRemoved unsafe.Pointer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(itemAdded *BackForwardListItem, itemsRemoved unsafe.Pointer))
	}

	var _itemAdded *BackForwardListItem // out
	var _itemsRemoved unsafe.Pointer    // out

	if arg1 != nil {
		_itemAdded = wrapBackForwardListItem(coreglib.Take(unsafe.Pointer(arg1)))
	}
	_itemsRemoved = (unsafe.Pointer)(unsafe.Pointer(arg2))

	f(_itemAdded, _itemsRemoved)
}
