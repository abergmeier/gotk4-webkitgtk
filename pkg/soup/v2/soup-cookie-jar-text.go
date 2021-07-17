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
		{T: externglib.Type(C.soup_cookie_jar_text_get_type()), F: marshalCookieJarTexter},
	})
}

type CookieJarText struct {
	CookieJar
}

var _ gextras.Nativer = (*CookieJarText)(nil)

func wrapCookieJarText(obj *externglib.Object) *CookieJarText {
	return &CookieJarText{
		CookieJar: CookieJar{
			Object: obj,
			SessionFeature: SessionFeature{
				Object: obj,
			},
		},
	}
}

func marshalCookieJarTexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCookieJarText(obj), nil
}

// NewCookieJarText creates a CookieJarText.
//
// filename will be read in at startup to create an initial set of cookies. If
// read_only is FALSE, then the non-session cookies will be written to filename
// when the 'changed' signal is emitted from the jar. (If read_only is TRUE,
// then the cookie jar will only be used for this session, and changes made to
// it will be lost when the jar is destroyed.)
func NewCookieJarText(filename string, readOnly bool) *CookieJarText {
	var _arg1 *C.char          // out
	var _arg2 C.gboolean       // out
	var _cret *C.SoupCookieJar // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	if readOnly {
		_arg2 = C.TRUE
	}

	_cret = C.soup_cookie_jar_text_new(_arg1, _arg2)

	var _cookieJarText *CookieJarText // out

	_cookieJarText = wrapCookieJarText(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cookieJarText
}

func (*CookieJarText) privateCookieJarText() {}
