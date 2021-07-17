// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_tld_error_get_type()), F: marshalTLDError},
	})
}

// TLDError: error codes for SOUP_TLD_ERROR.
type TLDError int

const (
	// TldErrorInvalidHostname: hostname was syntactically invalid.
	TldErrorInvalidHostname TLDError = iota
	// TldErrorIsIPAddress: passed-in "hostname" was actually an IP address (and
	// thus has no base domain or public suffix).
	TldErrorIsIPAddress
	// TldErrorNotEnoughDomains: passed-in hostname did not have enough
	// components. Eg, calling soup_tld_get_base_domain() on
	// <literal>"co.uk"</literal>.
	TldErrorNotEnoughDomains
	// TldErrorNoBaseDomain: passed-in hostname has no recognized public suffix.
	TldErrorNoBaseDomain
	TldErrorNoPslData
)

func marshalTLDError(p uintptr) (interface{}, error) {
	return TLDError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for TLDError.
func (t TLDError) String() string {
	switch t {
	case TldErrorInvalidHostname:
		return "InvalidHostname"
	case TldErrorIsIPAddress:
		return "IsIPAddress"
	case TldErrorNotEnoughDomains:
		return "NotEnoughDomains"
	case TldErrorNoBaseDomain:
		return "NoBaseDomain"
	case TldErrorNoPslData:
		return "NoPslData"
	default:
		return fmt.Sprintf("TLDError(%d)", t)
	}
}

// TldDomainIsPublicSuffix looks whether the domain passed as argument is a
// public domain suffix (.org, .com, .co.uk, etc) or not.
//
// Prior to libsoup 2.46, this function required that domain be in UTF-8 if it
// was an IDN. From 2.46 on, the name can be in either UTF-8 or ASCII format.
func TldDomainIsPublicSuffix(domain string) bool {
	var _arg1 *C.char    // out
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(domain)))

	_cret = C.soup_tld_domain_is_public_suffix(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TldGetBaseDomain finds the base domain for a given hostname. The base domain
// is composed by the top level domain (such as .org, .com, .co.uk, etc) plus
// the second level domain, for example for myhost.mydomain.com it will return
// mydomain.com.
//
// Note that NULL will be returned for private URLs (those not ending with any
// well known TLD) because choosing a base domain for them would be totally
// arbitrary.
//
// Prior to libsoup 2.46, this function required that hostname be in UTF-8 if it
// was an IDN. From 2.46 on, the name can be in either UTF-8 or ASCII format
// (and the return value will be in the same format).
func TldGetBaseDomain(hostname string) (string, error) {
	var _arg1 *C.char   // out
	var _cret *C.char   // in
	var _cerr *C.GError // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(hostname)))

	_cret = C.soup_tld_get_base_domain(_arg1, &_cerr)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8, _goerr
}
