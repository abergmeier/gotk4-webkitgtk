// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_auth_manager_get_type()), F: marshalAuthManagerer},
	})
}

// AuthManagerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type AuthManagerOverrider interface {
	Authenticate(msg *Message, auth Auther, retrying bool)
}

type AuthManager struct {
	*externglib.Object

	SessionFeature
}

var _ gextras.Nativer = (*AuthManager)(nil)

func wrapAuthManager(obj *externglib.Object) *AuthManager {
	return &AuthManager{
		Object: obj,
		SessionFeature: SessionFeature{
			Object: obj,
		},
	}
}

func marshalAuthManagerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAuthManager(obj), nil
}

// ClearCachedCredentials: clear all credentials cached by manager
func (manager *AuthManager) ClearCachedCredentials() {
	var _arg0 *C.SoupAuthManager // out

	_arg0 = (*C.SoupAuthManager)(unsafe.Pointer(manager.Native()))

	C.soup_auth_manager_clear_cached_credentials(_arg0)
}

// UseAuth records that auth is to be used under uri, as though a
// WWW-Authenticate header had been received at that URI. This can be used to
// "preload" manager's auth cache, to avoid an extra HTTP round trip in the case
// where you know ahead of time that a 401 response will be returned.
//
// This is only useful for authentication types where the initial Authorization
// header does not depend on any additional information from the server. (Eg,
// Basic or NTLM, but not Digest.)
func (manager *AuthManager) UseAuth(uri *URI, auth Auther) {
	var _arg0 *C.SoupAuthManager // out
	var _arg1 *C.SoupURI         // out
	var _arg2 *C.SoupAuth        // out

	_arg0 = (*C.SoupAuthManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))
	_arg2 = (*C.SoupAuth)(unsafe.Pointer((auth).(gextras.Nativer).Native()))

	C.soup_auth_manager_use_auth(_arg0, _arg1, _arg2)
}
