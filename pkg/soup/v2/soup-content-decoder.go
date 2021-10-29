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
		{T: externglib.Type(C.soup_content_decoder_get_type()), F: marshalContentDecoderer},
	})
}

type ContentDecoder struct {
	*externglib.Object

	SessionFeature
}

var (
	_ externglib.Objector = (*ContentDecoder)(nil)
)

func wrapContentDecoder(obj *externglib.Object) *ContentDecoder {
	return &ContentDecoder{
		Object: obj,
		SessionFeature: SessionFeature{
			Object: obj,
		},
	}
}

func marshalContentDecoderer(p uintptr) (interface{}, error) {
	return wrapContentDecoder(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
