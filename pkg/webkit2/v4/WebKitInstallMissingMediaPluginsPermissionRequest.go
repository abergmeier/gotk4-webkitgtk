// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: webkit2gtk-4.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_install_missing_media_plugins_permission_request_get_type()), F: marshalInstallMissingMediaPluginsPermissionRequester},
	})
}

type InstallMissingMediaPluginsPermissionRequest struct {
	*externglib.Object

	PermissionRequest
}

var _ gextras.Nativer = (*InstallMissingMediaPluginsPermissionRequest)(nil)

func wrapInstallMissingMediaPluginsPermissionRequest(obj *externglib.Object) *InstallMissingMediaPluginsPermissionRequest {
	return &InstallMissingMediaPluginsPermissionRequest{
		Object: obj,
		PermissionRequest: PermissionRequest{
			Object: obj,
		},
	}
}

func marshalInstallMissingMediaPluginsPermissionRequester(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapInstallMissingMediaPluginsPermissionRequest(obj), nil
}

// Description gets the description about the missing plugins provided by the
// media backend when a media couldn't be played.
func (request *InstallMissingMediaPluginsPermissionRequest) Description() string {
	var _arg0 *C.WebKitInstallMissingMediaPluginsPermissionRequest // out
	var _cret *C.gchar                                             // in

	_arg0 = (*C.WebKitInstallMissingMediaPluginsPermissionRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_install_missing_media_plugins_permission_request_get_description(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
