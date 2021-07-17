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
		{T: externglib.Type(C.webkit_web_inspector_get_type()), F: marshalWebInspectorrer},
	})
}

type WebInspector struct {
	*externglib.Object
}

var _ gextras.Nativer = (*WebInspector)(nil)

func wrapWebInspector(obj *externglib.Object) *WebInspector {
	return &WebInspector{
		Object: obj,
	}
}

func marshalWebInspectorrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapWebInspector(obj), nil
}

// Attach: request inspector to be attached. The signal KitWebInspector::attach
// will be emitted. If the inspector is already attached it does nothing.
func (inspector *WebInspector) Attach() {
	var _arg0 *C.WebKitWebInspector // out

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(inspector.Native()))

	C.webkit_web_inspector_attach(_arg0)
}

// Close: request inspector to be closed.
func (inspector *WebInspector) Close() {
	var _arg0 *C.WebKitWebInspector // out

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(inspector.Native()))

	C.webkit_web_inspector_close(_arg0)
}

// Detach: request inspector to be detached. The signal KitWebInspector::detach
// will be emitted. If the inspector is already detached it does nothing.
func (inspector *WebInspector) Detach() {
	var _arg0 *C.WebKitWebInspector // out

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(inspector.Native()))

	C.webkit_web_inspector_detach(_arg0)
}

// AttachedHeight: get the height that the inspector view should have when it's
// attached. If the inspector view is not attached this returns 0.
func (inspector *WebInspector) AttachedHeight() uint {
	var _arg0 *C.WebKitWebInspector // out
	var _cret C.guint               // in

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(inspector.Native()))

	_cret = C.webkit_web_inspector_get_attached_height(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// CanAttach: whether the inspector can be attached to the same window that
// contains the inspected view.
func (inspector *WebInspector) CanAttach() bool {
	var _arg0 *C.WebKitWebInspector // out
	var _cret C.gboolean            // in

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(inspector.Native()))

	_cret = C.webkit_web_inspector_get_can_attach(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InspectedURI: get the URI that is currently being inspected. This can be NULL
// if nothing has been loaded yet in the inspected view, if the inspector has
// been closed or when inspected view was loaded from a HTML string instead of a
// URI.
func (inspector *WebInspector) InspectedURI() string {
	var _arg0 *C.WebKitWebInspector // out
	var _cret *C.char               // in

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(inspector.Native()))

	_cret = C.webkit_web_inspector_get_inspected_uri(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// WebView: get the KitWebViewBase used to display the inspector. This might be
// NULL if the inspector hasn't been loaded yet, or it has been closed.
func (inspector *WebInspector) WebView() *WebViewBase {
	var _arg0 *C.WebKitWebInspector // out
	var _cret *C.WebKitWebViewBase  // in

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(inspector.Native()))

	_cret = C.webkit_web_inspector_get_web_view(_arg0)

	var _webViewBase *WebViewBase // out

	_webViewBase = wrapWebViewBase(externglib.Take(unsafe.Pointer(_cret)))

	return _webViewBase
}

// IsAttached: whether the inspector view is currently attached to the same
// window that contains the inspected view.
func (inspector *WebInspector) IsAttached() bool {
	var _arg0 *C.WebKitWebInspector // out
	var _cret C.gboolean            // in

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(inspector.Native()))

	_cret = C.webkit_web_inspector_is_attached(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Show: request inspector to be shown.
func (inspector *WebInspector) Show() {
	var _arg0 *C.WebKitWebInspector // out

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(inspector.Native()))

	C.webkit_web_inspector_show(_arg0)
}
