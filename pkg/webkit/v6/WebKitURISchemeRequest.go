// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/soup/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit/webkit.h>
import "C"

// GType values.
var (
	GTypeURISchemeRequest = coreglib.Type(C.webkit_uri_scheme_request_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeURISchemeRequest, F: marshalURISchemeRequest},
	})
}

// URISchemeRequestOverrides contains methods that are overridable.
type URISchemeRequestOverrides struct {
}

func defaultURISchemeRequestOverrides(v *URISchemeRequest) URISchemeRequestOverrides {
	return URISchemeRequestOverrides{}
}

// URISchemeRequest represents a URI scheme request.
//
// If you register a particular URI scheme in a KitWebContext,
// using webkit_web_context_register_uri_scheme(), you have to provide a
// KitURISchemeRequestCallback. After that, when a URI request is made with
// that particular scheme, your callback will be called. There you will be
// able to access properties such as the scheme, the URI and path, and the
// KitWebView that initiated the request, and also finish the request with
// webkit_uri_scheme_request_finish().
type URISchemeRequest struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*URISchemeRequest)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*URISchemeRequest, *URISchemeRequestClass, URISchemeRequestOverrides](
		GTypeURISchemeRequest,
		initURISchemeRequestClass,
		wrapURISchemeRequest,
		defaultURISchemeRequestOverrides,
	)
}

func initURISchemeRequestClass(gclass unsafe.Pointer, overrides URISchemeRequestOverrides, classInitFunc func(*URISchemeRequestClass)) {
	if classInitFunc != nil {
		class := (*URISchemeRequestClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapURISchemeRequest(obj *coreglib.Object) *URISchemeRequest {
	return &URISchemeRequest{
		Object: obj,
	}
}

func marshalURISchemeRequest(p uintptr) (interface{}, error) {
	return wrapURISchemeRequest(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Finish a KitURISchemeRequest by setting the contents of the request and its
// mime type.
//
// The function takes the following parameters:
//
//   - stream to read the contents of the request.
//   - streamLength: length of the stream or -1 if not known.
//   - contentType (optional): content type of the stream or NULL if not known.
//
func (request *URISchemeRequest) Finish(stream gio.InputStreamer, streamLength int64, contentType string) {
	var _arg0 *C.WebKitURISchemeRequest // out
	var _arg1 *C.GInputStream           // out
	var _arg2 C.gint64                  // out
	var _arg3 *C.gchar                  // out

	_arg0 = (*C.WebKitURISchemeRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))
	_arg1 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	_arg2 = C.gint64(streamLength)
	if contentType != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(contentType)))
		defer C.free(unsafe.Pointer(_arg3))
	}

	C.webkit_uri_scheme_request_finish(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(request)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(streamLength)
	runtime.KeepAlive(contentType)
}

// FinishError: finish a KitURISchemeRequest with a #GError.
//
// The function takes the following parameters:
//
//   - err that will be passed to the KitWebView.
//
func (request *URISchemeRequest) FinishError(err error) {
	var _arg0 *C.WebKitURISchemeRequest // out
	var _arg1 *C.GError                 // out

	_arg0 = (*C.WebKitURISchemeRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))
	if err != nil {
		_arg1 = (*C.GError)(gerror.New(err))
	}

	C.webkit_uri_scheme_request_finish_error(_arg0, _arg1)
	runtime.KeepAlive(request)
	runtime.KeepAlive(err)
}

// FinishWithResponse: finish a KitURISchemeRequest by returning a
// KitURISchemeResponse.
//
// The function takes the following parameters:
//
//   - response: KitURISchemeResponse.
//
func (request *URISchemeRequest) FinishWithResponse(response *URISchemeResponse) {
	var _arg0 *C.WebKitURISchemeRequest  // out
	var _arg1 *C.WebKitURISchemeResponse // out

	_arg0 = (*C.WebKitURISchemeRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))
	_arg1 = (*C.WebKitURISchemeResponse)(unsafe.Pointer(coreglib.InternObject(response).Native()))

	C.webkit_uri_scheme_request_finish_with_response(_arg0, _arg1)
	runtime.KeepAlive(request)
	runtime.KeepAlive(response)
}

// HTTPBody: get the request body.
//
// The function returns the following values:
//
//   - inputStream: (nullable): the body of the request.
//
func (request *URISchemeRequest) HTTPBody() gio.InputStreamer {
	var _arg0 *C.WebKitURISchemeRequest // out
	var _cret *C.GInputStream           // in

	_arg0 = (*C.WebKitURISchemeRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_uri_scheme_request_get_http_body(_arg0)
	runtime.KeepAlive(request)

	var _inputStream gio.InputStreamer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.InputStreamer is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gio.InputStreamer)
			return ok
		})
		rv, ok := casted.(gio.InputStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.InputStreamer")
		}
		_inputStream = rv
	}

	return _inputStream
}

// HTTPHeaders: get the MessageHeaders of the request.
//
// The function returns the following values:
//
//   - messageHeaders of the request.
//
func (request *URISchemeRequest) HTTPHeaders() *soup.MessageHeaders {
	var _arg0 *C.WebKitURISchemeRequest // out
	var _cret *C.SoupMessageHeaders     // in

	_arg0 = (*C.WebKitURISchemeRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_uri_scheme_request_get_http_headers(_arg0)
	runtime.KeepAlive(request)

	var _messageHeaders *soup.MessageHeaders // out

	_messageHeaders = (*soup.MessageHeaders)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.soup_message_headers_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_messageHeaders)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_message_headers_unref((*C.SoupMessageHeaders)(intern.C))
		},
	)

	return _messageHeaders
}

// HTTPMethod: get the HTTP method of the request.
//
// The function returns the following values:
//
//   - utf8: HTTP method of the request.
//
func (request *URISchemeRequest) HTTPMethod() string {
	var _arg0 *C.WebKitURISchemeRequest // out
	var _cret *C.gchar                  // in

	_arg0 = (*C.WebKitURISchemeRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_uri_scheme_request_get_http_method(_arg0)
	runtime.KeepAlive(request)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Path: get the URI path of request.
//
// The function returns the following values:
//
//   - utf8: URI path of request.
//
func (request *URISchemeRequest) Path() string {
	var _arg0 *C.WebKitURISchemeRequest // out
	var _cret *C.gchar                  // in

	_arg0 = (*C.WebKitURISchemeRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_uri_scheme_request_get_path(_arg0)
	runtime.KeepAlive(request)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Scheme: get the URI scheme of request.
//
// The function returns the following values:
//
//   - utf8: URI scheme of request.
//
func (request *URISchemeRequest) Scheme() string {
	var _arg0 *C.WebKitURISchemeRequest // out
	var _cret *C.gchar                  // in

	_arg0 = (*C.WebKitURISchemeRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_uri_scheme_request_get_scheme(_arg0)
	runtime.KeepAlive(request)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// URI: get the URI of request.
//
// The function returns the following values:
//
//   - utf8: full URI of request.
//
func (request *URISchemeRequest) URI() string {
	var _arg0 *C.WebKitURISchemeRequest // out
	var _cret *C.gchar                  // in

	_arg0 = (*C.WebKitURISchemeRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_uri_scheme_request_get_uri(_arg0)
	runtime.KeepAlive(request)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// WebView: get the KitWebView that initiated the request.
//
// The function returns the following values:
//
//   - webView that initiated request.
//
func (request *URISchemeRequest) WebView() *WebView {
	var _arg0 *C.WebKitURISchemeRequest // out
	var _cret *C.WebKitWebView          // in

	_arg0 = (*C.WebKitURISchemeRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_uri_scheme_request_get_web_view(_arg0)
	runtime.KeepAlive(request)

	var _webView *WebView // out

	_webView = wrapWebView(coreglib.Take(unsafe.Pointer(_cret)))

	return _webView
}

// URISchemeRequestClass: instance of this type is always passed by reference.
type URISchemeRequestClass struct {
	*uriSchemeRequestClass
}

// uriSchemeRequestClass is the struct that's finalized.
type uriSchemeRequestClass struct {
	native *C.WebKitURISchemeRequestClass
}
