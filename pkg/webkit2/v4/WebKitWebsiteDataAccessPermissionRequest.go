// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: webkit2gtk-4.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_website_data_access_permission_request_get_type()), F: marshalWebsiteDataAccessPermissionRequester},
	})
}

type WebsiteDataAccessPermissionRequest struct {
	*externglib.Object

	PermissionRequest
}

var (
	_ externglib.Objector = (*WebsiteDataAccessPermissionRequest)(nil)
)

func wrapWebsiteDataAccessPermissionRequest(obj *externglib.Object) *WebsiteDataAccessPermissionRequest {
	return &WebsiteDataAccessPermissionRequest{
		Object: obj,
		PermissionRequest: PermissionRequest{
			Object: obj,
		},
	}
}

func marshalWebsiteDataAccessPermissionRequester(p uintptr) (interface{}, error) {
	return wrapWebsiteDataAccessPermissionRequest(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CurrentDomain: get the current domain being browsed.
func (request *WebsiteDataAccessPermissionRequest) CurrentDomain() string {
	var _arg0 *C.WebKitWebsiteDataAccessPermissionRequest // out
	var _cret *C.char                                     // in

	_arg0 = (*C.WebKitWebsiteDataAccessPermissionRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_website_data_access_permission_request_get_current_domain(_arg0)
	runtime.KeepAlive(request)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// RequestingDomain: get the domain requesting permission to access its cookies
// while browsing the current domain.
func (request *WebsiteDataAccessPermissionRequest) RequestingDomain() string {
	var _arg0 *C.WebKitWebsiteDataAccessPermissionRequest // out
	var _cret *C.char                                     // in

	_arg0 = (*C.WebKitWebsiteDataAccessPermissionRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_website_data_access_permission_request_get_requesting_domain(_arg0)
	runtime.KeepAlive(request)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
