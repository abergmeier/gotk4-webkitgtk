// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <webkit/webkit.h>
import "C"

// MediaKeySystemPermissionGetName: get the key system for which access
// permission is being requested.
//
// The function takes the following parameters:
//
//   - request: KitMediaKeySystemPermissionRequest.
//
// The function returns the following values:
//
//   - utf8: key system name for request.
//
func MediaKeySystemPermissionGetName(request *MediaKeySystemPermissionRequest) string {
	var _arg1 *C.WebKitMediaKeySystemPermissionRequest // out
	var _cret *C.gchar                                 // in

	_arg1 = (*C.WebKitMediaKeySystemPermissionRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_media_key_system_permission_get_name(_arg1)
	runtime.KeepAlive(request)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
