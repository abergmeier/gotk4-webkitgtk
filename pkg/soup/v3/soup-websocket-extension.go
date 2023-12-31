// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
// extern gboolean _gotk4_soup3_WebsocketExtensionClass_configure(SoupWebsocketExtension*, SoupWebsocketConnectionType, GHashTable*, GError**);
// extern char* _gotk4_soup3_WebsocketExtensionClass_get_response_params(SoupWebsocketExtension*);
// extern char* _gotk4_soup3_WebsocketExtensionClass_get_request_params(SoupWebsocketExtension*);
// char* _gotk4_soup3_WebsocketExtension_virtual_get_request_params(void* fnptr, SoupWebsocketExtension* arg0) {
//   return ((char* (*)(SoupWebsocketExtension*))(fnptr))(arg0);
// };
// char* _gotk4_soup3_WebsocketExtension_virtual_get_response_params(void* fnptr, SoupWebsocketExtension* arg0) {
//   return ((char* (*)(SoupWebsocketExtension*))(fnptr))(arg0);
// };
// gboolean _gotk4_soup3_WebsocketExtension_virtual_configure(void* fnptr, SoupWebsocketExtension* arg0, SoupWebsocketConnectionType arg1, GHashTable* arg2, GError** arg3) {
//   return ((gboolean (*)(SoupWebsocketExtension*, SoupWebsocketConnectionType, GHashTable*, GError**))(fnptr))(arg0, arg1, arg2, arg3);
// };
import "C"

// GType values.
var (
	GTypeWebsocketExtension = coreglib.Type(C.soup_websocket_extension_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWebsocketExtension, F: marshalWebsocketExtension},
	})
}

// WebsocketExtensionOverrides contains methods that are overridable.
type WebsocketExtensionOverrides struct {
	// Configure configures extension with the given params.
	//
	// The function takes the following parameters:
	//
	//   - connectionType: either SOUP_WEBSOCKET_CONNECTION_CLIENT or
	//     SOUP_WEBSOCKET_CONNECTION_SERVER.
	//   - params (optional): parameters, or NULL.
	//
	Configure func(connectionType WebsocketConnectionType, params map[unsafe.Pointer]unsafe.Pointer) error
	// RequestParams: get the parameters strings to be included in the request
	// header. If the extension doesn't include any parameter in the request,
	// this function returns NULL.
	//
	// The function returns the following values:
	//
	//   - utf8 (optional): new allocated string with the parameters.
	//
	RequestParams func() string
	// ResponseParams: get the parameters strings to be included in the response
	// header. If the extension doesn't include any parameter in the response,
	// this function returns NULL.
	//
	// The function returns the following values:
	//
	//   - utf8 (optional): new allocated string with the parameters.
	//
	ResponseParams func() string
}

func defaultWebsocketExtensionOverrides(v *WebsocketExtension) WebsocketExtensionOverrides {
	return WebsocketExtensionOverrides{
		Configure:      v.configure,
		RequestParams:  v.requestParams,
		ResponseParams: v.responseParams,
	}
}

// WebsocketExtension class for impelementing websocket extensions.
type WebsocketExtension struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*WebsocketExtension)(nil)
)

// WebsocketExtensioner describes types inherited from class WebsocketExtension.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type WebsocketExtensioner interface {
	coreglib.Objector
	baseWebsocketExtension() *WebsocketExtension
}

var _ WebsocketExtensioner = (*WebsocketExtension)(nil)

func init() {
	coreglib.RegisterClassInfo[*WebsocketExtension, *WebsocketExtensionClass, WebsocketExtensionOverrides](
		GTypeWebsocketExtension,
		initWebsocketExtensionClass,
		wrapWebsocketExtension,
		defaultWebsocketExtensionOverrides,
	)
}

func initWebsocketExtensionClass(gclass unsafe.Pointer, overrides WebsocketExtensionOverrides, classInitFunc func(*WebsocketExtensionClass)) {
	pclass := (*C.SoupWebsocketExtensionClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeWebsocketExtension))))

	if overrides.Configure != nil {
		pclass.configure = (*[0]byte)(C._gotk4_soup3_WebsocketExtensionClass_configure)
	}

	if overrides.RequestParams != nil {
		pclass.get_request_params = (*[0]byte)(C._gotk4_soup3_WebsocketExtensionClass_get_request_params)
	}

	if overrides.ResponseParams != nil {
		pclass.get_response_params = (*[0]byte)(C._gotk4_soup3_WebsocketExtensionClass_get_response_params)
	}

	if classInitFunc != nil {
		class := (*WebsocketExtensionClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapWebsocketExtension(obj *coreglib.Object) *WebsocketExtension {
	return &WebsocketExtension{
		Object: obj,
	}
}

func marshalWebsocketExtension(p uintptr) (interface{}, error) {
	return wrapWebsocketExtension(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (extension *WebsocketExtension) baseWebsocketExtension() *WebsocketExtension {
	return extension
}

// BaseWebsocketExtension returns the underlying base object.
func BaseWebsocketExtension(obj WebsocketExtensioner) *WebsocketExtension {
	return obj.baseWebsocketExtension()
}

// Configure configures extension with the given params.
//
// The function takes the following parameters:
//
//   - connectionType: either SOUP_WEBSOCKET_CONNECTION_CLIENT or
//     SOUP_WEBSOCKET_CONNECTION_SERVER.
//   - params (optional): parameters, or NULL.
//
func (extension *WebsocketExtension) Configure(connectionType WebsocketConnectionType, params map[unsafe.Pointer]unsafe.Pointer) error {
	var _arg0 *C.SoupWebsocketExtension     // out
	var _arg1 C.SoupWebsocketConnectionType // out
	var _arg2 *C.GHashTable                 // out
	var _cerr *C.GError                     // in

	_arg0 = (*C.SoupWebsocketExtension)(unsafe.Pointer(coreglib.InternObject(extension).Native()))
	_arg1 = C.SoupWebsocketConnectionType(connectionType)
	if params != nil {
		_arg2 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
		for ksrc, vsrc := range params {
			var kdst *C.gpointer // out
			var vdst *C.gpointer // out
			kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
			vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
			C.g_hash_table_insert(_arg2, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
		}
		defer C.g_hash_table_unref(_arg2)
	}

	C.soup_websocket_extension_configure(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(extension)
	runtime.KeepAlive(connectionType)
	runtime.KeepAlive(params)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// RequestParams: get the parameters strings to be included in the request
// header. If the extension doesn't include any parameter in the request,
// this function returns NULL.
//
// The function returns the following values:
//
//   - utf8 (optional): new allocated string with the parameters.
//
func (extension *WebsocketExtension) RequestParams() string {
	var _arg0 *C.SoupWebsocketExtension // out
	var _cret *C.char                   // in

	_arg0 = (*C.SoupWebsocketExtension)(unsafe.Pointer(coreglib.InternObject(extension).Native()))

	_cret = C.soup_websocket_extension_get_request_params(_arg0)
	runtime.KeepAlive(extension)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// ResponseParams: get the parameters strings to be included in the response
// header. If the extension doesn't include any parameter in the response,
// this function returns NULL.
//
// The function returns the following values:
//
//   - utf8 (optional): new allocated string with the parameters.
//
func (extension *WebsocketExtension) ResponseParams() string {
	var _arg0 *C.SoupWebsocketExtension // out
	var _cret *C.char                   // in

	_arg0 = (*C.SoupWebsocketExtension)(unsafe.Pointer(coreglib.InternObject(extension).Native()))

	_cret = C.soup_websocket_extension_get_response_params(_arg0)
	runtime.KeepAlive(extension)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Configure configures extension with the given params.
//
// The function takes the following parameters:
//
//   - connectionType: either SOUP_WEBSOCKET_CONNECTION_CLIENT or
//     SOUP_WEBSOCKET_CONNECTION_SERVER.
//   - params (optional): parameters, or NULL.
//
func (extension *WebsocketExtension) configure(connectionType WebsocketConnectionType, params map[unsafe.Pointer]unsafe.Pointer) error {
	gclass := (*C.SoupWebsocketExtensionClass)(coreglib.PeekParentClass(extension))
	fnarg := gclass.configure

	var _arg0 *C.SoupWebsocketExtension     // out
	var _arg1 C.SoupWebsocketConnectionType // out
	var _arg2 *C.GHashTable                 // out
	var _cerr *C.GError                     // in

	_arg0 = (*C.SoupWebsocketExtension)(unsafe.Pointer(coreglib.InternObject(extension).Native()))
	_arg1 = C.SoupWebsocketConnectionType(connectionType)
	if params != nil {
		_arg2 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
		for ksrc, vsrc := range params {
			var kdst *C.gpointer // out
			var vdst *C.gpointer // out
			kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
			vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
			C.g_hash_table_insert(_arg2, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
		}
		defer C.g_hash_table_unref(_arg2)
	}

	C._gotk4_soup3_WebsocketExtension_virtual_configure(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(extension)
	runtime.KeepAlive(connectionType)
	runtime.KeepAlive(params)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// requestParams: get the parameters strings to be included in the request
// header. If the extension doesn't include any parameter in the request,
// this function returns NULL.
//
// The function returns the following values:
//
//   - utf8 (optional): new allocated string with the parameters.
//
func (extension *WebsocketExtension) requestParams() string {
	gclass := (*C.SoupWebsocketExtensionClass)(coreglib.PeekParentClass(extension))
	fnarg := gclass.get_request_params

	var _arg0 *C.SoupWebsocketExtension // out
	var _cret *C.char                   // in

	_arg0 = (*C.SoupWebsocketExtension)(unsafe.Pointer(coreglib.InternObject(extension).Native()))

	_cret = C._gotk4_soup3_WebsocketExtension_virtual_get_request_params(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(extension)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// responseParams: get the parameters strings to be included in the response
// header. If the extension doesn't include any parameter in the response,
// this function returns NULL.
//
// The function returns the following values:
//
//   - utf8 (optional): new allocated string with the parameters.
//
func (extension *WebsocketExtension) responseParams() string {
	gclass := (*C.SoupWebsocketExtensionClass)(coreglib.PeekParentClass(extension))
	fnarg := gclass.get_response_params

	var _arg0 *C.SoupWebsocketExtension // out
	var _cret *C.char                   // in

	_arg0 = (*C.SoupWebsocketExtension)(unsafe.Pointer(coreglib.InternObject(extension).Native()))

	_cret = C._gotk4_soup3_WebsocketExtension_virtual_get_response_params(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(extension)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// WebsocketExtensionClass class structure for the SoupWebsocketExtension.
//
// An instance of this type is always passed by reference.
type WebsocketExtensionClass struct {
	*websocketExtensionClass
}

// websocketExtensionClass is the struct that's finalized.
type websocketExtensionClass struct {
	native *C.SoupWebsocketExtensionClass
}

// Name: name of the extension.
func (w *WebsocketExtensionClass) Name() string {
	valptr := &w.native.name
	var _v string // out
	_v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return _v
}
