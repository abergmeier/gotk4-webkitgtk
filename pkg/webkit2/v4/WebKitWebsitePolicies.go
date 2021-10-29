// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"fmt"
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
		{T: externglib.Type(C.webkit_autoplay_policy_get_type()), F: marshalAutoplayPolicy},
		{T: externglib.Type(C.webkit_website_policies_get_type()), F: marshalWebsitePolicieser},
	})
}

// AutoplayPolicy: enum values used to specify autoplay policies.
type AutoplayPolicy C.gint

const (
	// AutoplayAllow: do not restrict autoplay.
	AutoplayAllow AutoplayPolicy = iota
	// AutoplayAllowWithoutSound: allow videos to autoplay if they have no audio
	// track, or if their audio track is muted.
	AutoplayAllowWithoutSound
	// AutoplayDeny: never allow autoplay.
	AutoplayDeny
)

func marshalAutoplayPolicy(p uintptr) (interface{}, error) {
	return AutoplayPolicy(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for AutoplayPolicy.
func (a AutoplayPolicy) String() string {
	switch a {
	case AutoplayAllow:
		return "Allow"
	case AutoplayAllowWithoutSound:
		return "AllowWithoutSound"
	case AutoplayDeny:
		return "Deny"
	default:
		return fmt.Sprintf("AutoplayPolicy(%d)", a)
	}
}

type WebsitePolicies struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*WebsitePolicies)(nil)
)

func wrapWebsitePolicies(obj *externglib.Object) *WebsitePolicies {
	return &WebsitePolicies{
		Object: obj,
	}
}

func marshalWebsitePolicieser(p uintptr) (interface{}, error) {
	return wrapWebsitePolicies(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewWebsitePolicies: create a new KitWebsitePolicies.
func NewWebsitePolicies() *WebsitePolicies {
	var _cret *C.WebKitWebsitePolicies // in

	_cret = C.webkit_website_policies_new()

	var _websitePolicies *WebsitePolicies // out

	_websitePolicies = wrapWebsitePolicies(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _websitePolicies
}

// AutoplayPolicy: get the KitWebsitePolicies:autoplay property.
func (policies *WebsitePolicies) AutoplayPolicy() AutoplayPolicy {
	var _arg0 *C.WebKitWebsitePolicies // out
	var _cret C.WebKitAutoplayPolicy   // in

	_arg0 = (*C.WebKitWebsitePolicies)(unsafe.Pointer(policies.Native()))

	_cret = C.webkit_website_policies_get_autoplay_policy(_arg0)
	runtime.KeepAlive(policies)

	var _autoplayPolicy AutoplayPolicy // out

	_autoplayPolicy = AutoplayPolicy(_cret)

	return _autoplayPolicy
}
