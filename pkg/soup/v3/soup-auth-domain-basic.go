// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
// extern void callbackDelete(gpointer);
// extern gboolean _gotk4_soup3_AuthDomainBasicAuthCallback(SoupAuthDomain*, SoupServerMessage*, char*, char*, gpointer);
import "C"

// GType values.
var (
	GTypeAuthDomainBasic = coreglib.Type(C.soup_auth_domain_basic_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAuthDomainBasic, F: marshalAuthDomainBasic},
	})
}

// AuthDomainBasicAuthCallback: callback used by AuthDomainBasic for
// authentication purposes. The application should verify that username and
// password and valid and return TRUE or FALSE.
//
// If you are maintaining your own password database (rather than using the
// password to authenticate against some other system like PAM or a remote
// server), you should make sure you know what you are doing. In particular,
// don't store cleartext passwords, or easily-computed hashes of cleartext
// passwords, even if you don't care that much about the security of your
// server, because users will frequently use the same password for multiple
// sites, and so compromising any site with a cleartext (or easily-cracked)
// password database may give attackers access to other more-interesting sites
// as well.
type AuthDomainBasicAuthCallback func(domain *AuthDomainBasic, msg *ServerMessage, username, password string) (ok bool)

// AuthDomainBasicOverrides contains methods that are overridable.
type AuthDomainBasicOverrides struct {
}

func defaultAuthDomainBasicOverrides(v *AuthDomainBasic) AuthDomainBasicOverrides {
	return AuthDomainBasicOverrides{}
}

// AuthDomainBasic subclass of AuthDomain for Basic authentication.
type AuthDomainBasic struct {
	_ [0]func() // equal guard
	AuthDomain
}

var (
	_ AuthDomainer = (*AuthDomainBasic)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*AuthDomainBasic, *AuthDomainBasicClass, AuthDomainBasicOverrides](
		GTypeAuthDomainBasic,
		initAuthDomainBasicClass,
		wrapAuthDomainBasic,
		defaultAuthDomainBasicOverrides,
	)
}

func initAuthDomainBasicClass(gclass unsafe.Pointer, overrides AuthDomainBasicOverrides, classInitFunc func(*AuthDomainBasicClass)) {
	if classInitFunc != nil {
		class := (*AuthDomainBasicClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapAuthDomainBasic(obj *coreglib.Object) *AuthDomainBasic {
	return &AuthDomainBasic{
		AuthDomain: AuthDomain{
			Object: obj,
		},
	}
}

func marshalAuthDomainBasic(p uintptr) (interface{}, error) {
	return wrapAuthDomainBasic(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// SetAuthCallback sets the callback that domain will use to authenticate
// incoming requests. For each request containing authorization, domain will
// invoke the callback, and then either accept or reject the request based on
// callback's return value.
//
// You can also set the auth callback by setting the
// SoupAuthDomainBasic:auth-callback and SoupAuthDomainBasic:auth-data
// properties, which can also be used to set the callback at construct time.
//
// The function takes the following parameters:
//
//   - callback: callback.
//
func (domain *AuthDomainBasic) SetAuthCallback(callback AuthDomainBasicAuthCallback) {
	var _arg0 *C.SoupAuthDomain                 // out
	var _arg1 C.SoupAuthDomainBasicAuthCallback // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.SoupAuthDomain)(unsafe.Pointer(coreglib.InternObject(domain).Native()))
	_arg1 = (*[0]byte)(C._gotk4_soup3_AuthDomainBasicAuthCallback)
	_arg2 = C.gpointer(gbox.Assign(callback))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.soup_auth_domain_basic_set_auth_callback(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(callback)
}

// AuthDomainBasicClass: instance of this type is always passed by reference.
type AuthDomainBasicClass struct {
	*authDomainBasicClass
}

// authDomainBasicClass is the struct that's finalized.
type authDomainBasicClass struct {
	native *C.SoupAuthDomainBasicClass
}

func (a *AuthDomainBasicClass) ParentClass() *AuthDomainClass {
	valptr := &a.native.parent_class
	var _v *AuthDomainClass // out
	_v = (*AuthDomainClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
