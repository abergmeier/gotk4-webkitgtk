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
		{T: externglib.Type(C.soup_session_sync_get_type()), F: marshalSessionSyncer},
	})
}

type SessionSync struct {
	Session
}

var (
	_ externglib.Objector = (*SessionSync)(nil)
)

func wrapSessionSync(obj *externglib.Object) *SessionSync {
	return &SessionSync{
		Session: Session{
			Object: obj,
		},
	}
}

func marshalSessionSyncer(p uintptr) (interface{}, error) {
	return wrapSessionSync(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSessionSync creates an synchronous Session with the default options.
//
// Deprecated: SessionSync is deprecated; use a plain Session, created with
// soup_session_new(). See the <link linkend="libsoup-session-porting">porting
// guide</link>.
func NewSessionSync() *SessionSync {
	var _cret *C.SoupSession // in

	_cret = C.soup_session_sync_new()

	var _sessionSync *SessionSync // out

	_sessionSync = wrapSessionSync(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _sessionSync
}
