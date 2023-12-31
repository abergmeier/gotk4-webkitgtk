// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/javascriptcore/v6"
)

// #include <stdlib.h>
// #include <webkit/webkit.h>
import "C"

//export _gotk4_webkit6_UserContentManager_ConnectScriptMessageReceived
func _gotk4_webkit6_UserContentManager_ConnectScriptMessageReceived(arg0 C.gpointer, arg1 *C.JSCValue, arg2 C.guintptr) {
	var f func(value *javascriptcore.Value)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(value *javascriptcore.Value))
	}

	var _value *javascriptcore.Value // out

	{
		obj := coreglib.Take(unsafe.Pointer(arg1))
		_value = &javascriptcore.Value{
			Object: obj,
		}
	}

	f(_value)
}

//export _gotk4_webkit6_UserContentManager_ConnectScriptMessageWithReplyReceived
func _gotk4_webkit6_UserContentManager_ConnectScriptMessageWithReplyReceived(arg0 C.gpointer, arg1 *C.JSCValue, arg2 *C.WebKitScriptMessageReply, arg3 C.guintptr) (cret C.gboolean) {
	var f func(value *javascriptcore.Value, reply *ScriptMessageReply) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(value *javascriptcore.Value, reply *ScriptMessageReply) (ok bool))
	}

	var _value *javascriptcore.Value // out
	var _reply *ScriptMessageReply   // out

	{
		obj := coreglib.Take(unsafe.Pointer(arg1))
		_value = &javascriptcore.Value{
			Object: obj,
		}
	}
	_reply = (*ScriptMessageReply)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	C.webkit_script_message_reply_ref(arg2)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_reply)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_script_message_reply_unref((*C.WebKitScriptMessageReply)(intern.C))
		},
	)

	ok := f(_value, _reply)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}
