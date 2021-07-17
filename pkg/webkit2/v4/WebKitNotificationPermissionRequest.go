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
		{T: externglib.Type(C.webkit_notification_permission_request_get_type()), F: marshalNotificationPermissionRequester},
	})
}

type NotificationPermissionRequest struct {
	*externglib.Object

	PermissionRequest
}

var _ gextras.Nativer = (*NotificationPermissionRequest)(nil)

func wrapNotificationPermissionRequest(obj *externglib.Object) *NotificationPermissionRequest {
	return &NotificationPermissionRequest{
		Object: obj,
		PermissionRequest: PermissionRequest{
			Object: obj,
		},
	}
}

func marshalNotificationPermissionRequester(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNotificationPermissionRequest(obj), nil
}

func (*NotificationPermissionRequest) privateNotificationPermissionRequest() {}
