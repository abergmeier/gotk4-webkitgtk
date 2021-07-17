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
		{T: externglib.Type(C.webkit_response_policy_decision_get_type()), F: marshalResponsePolicyDecisioner},
	})
}

type ResponsePolicyDecision struct {
	PolicyDecision
}

var _ gextras.Nativer = (*ResponsePolicyDecision)(nil)

func wrapResponsePolicyDecision(obj *externglib.Object) *ResponsePolicyDecision {
	return &ResponsePolicyDecision{
		PolicyDecision: PolicyDecision{
			Object: obj,
		},
	}
}

func marshalResponsePolicyDecisioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapResponsePolicyDecision(obj), nil
}

// Request: return the KitURIRequest associated with the response decision.
// Modifications to the returned object are <emphasis>not</emphasis> taken into
// account when the request is sent over the network, and is intended only to
// aid in evaluating whether a response decision should be taken or not. To
// modify requests before they are sent over the network the
// KitPage::send-request signal can be used instead.
func (decision *ResponsePolicyDecision) Request() *URIRequest {
	var _arg0 *C.WebKitResponsePolicyDecision // out
	var _cret *C.WebKitURIRequest             // in

	_arg0 = (*C.WebKitResponsePolicyDecision)(unsafe.Pointer(decision.Native()))

	_cret = C.webkit_response_policy_decision_get_request(_arg0)

	var _uriRequest *URIRequest // out

	_uriRequest = wrapURIRequest(externglib.Take(unsafe.Pointer(_cret)))

	return _uriRequest
}

// Response gets the value of the KitResponsePolicyDecision:response property.
func (decision *ResponsePolicyDecision) Response() *URIResponse {
	var _arg0 *C.WebKitResponsePolicyDecision // out
	var _cret *C.WebKitURIResponse            // in

	_arg0 = (*C.WebKitResponsePolicyDecision)(unsafe.Pointer(decision.Native()))

	_cret = C.webkit_response_policy_decision_get_response(_arg0)

	var _uriResponse *URIResponse // out

	_uriResponse = wrapURIResponse(externglib.Take(unsafe.Pointer(_cret)))

	return _uriResponse
}

// IsMIMETypeSupported gets whether the MIME type of the response can be
// displayed in the KitWebView that triggered this policy decision request. See
// also webkit_web_view_can_show_mime_type().
func (decision *ResponsePolicyDecision) IsMIMETypeSupported() bool {
	var _arg0 *C.WebKitResponsePolicyDecision // out
	var _cret C.gboolean                      // in

	_arg0 = (*C.WebKitResponsePolicyDecision)(unsafe.Pointer(decision.Native()))

	_cret = C.webkit_response_policy_decision_is_mime_type_supported(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
