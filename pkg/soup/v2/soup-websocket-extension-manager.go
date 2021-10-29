// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_websocket_extension_manager_get_type()), F: marshalWebsocketExtensionManagerer},
	})
}

type WebsocketExtensionManager struct {
	*externglib.Object

	SessionFeature
}

var (
	_ externglib.Objector = (*WebsocketExtensionManager)(nil)
)

func wrapWebsocketExtensionManager(obj *externglib.Object) *WebsocketExtensionManager {
	return &WebsocketExtensionManager{
		Object: obj,
		SessionFeature: SessionFeature{
			Object: obj,
		},
	}
}

func marshalWebsocketExtensionManagerer(p uintptr) (interface{}, error) {
	return wrapWebsocketExtensionManager(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
