// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
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
		{T: externglib.Type(C.webkit_web_inspector_get_type()), F: marshalWebInspectorrer},
	})
}

type WebInspector struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*WebInspector)(nil)
)

func wrapWebInspector(obj *externglib.Object) *WebInspector {
	return &WebInspector{
		Object: obj,
	}
}

func marshalWebInspectorrer(p uintptr) (interface{}, error) {
	return wrapWebInspector(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Attach: request inspector to be attached. The signal KitWebInspector::attach
// will be emitted. If the inspector is already attached it does nothing.
func (inspector *WebInspector) Attach() {
	var _arg0 *C.WebKitWebInspector // out

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(inspector.Native()))

	C.webkit_web_inspector_attach(_arg0)
	runtime.KeepAlive(inspector)
}

// Close: request inspector to be closed.
func (inspector *WebInspector) Close() {
	var _arg0 *C.WebKitWebInspector // out

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(inspector.Native()))

	C.webkit_web_inspector_close(_arg0)
	runtime.KeepAlive(inspector)
}

// Detach: request inspector to be detached. The signal KitWebInspector::detach
// will be emitted. If the inspector is already detached it does nothing.
func (inspector *WebInspector) Detach() {
	var _arg0 *C.WebKitWebInspector // out

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(inspector.Native()))

	C.webkit_web_inspector_detach(_arg0)
	runtime.KeepAlive(inspector)
}

// AttachedHeight: get the height that the inspector view should have when it's
// attached. If the inspector view is not attached this returns 0.
func (inspector *WebInspector) AttachedHeight() uint {
	var _arg0 *C.WebKitWebInspector // out
	var _cret C.guint               // in

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(inspector.Native()))

	_cret = C.webkit_web_inspector_get_attached_height(_arg0)
	runtime.KeepAlive(inspector)

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
	runtime.KeepAlive(inspector)

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
	runtime.KeepAlive(inspector)

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
	runtime.KeepAlive(inspector)

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
	runtime.KeepAlive(inspector)

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
	runtime.KeepAlive(inspector)
}

// ConnectAttach: emitted when the inspector is requested to be attached to the
// window where the inspected web view is. If this signal is not handled the
// inspector view will be automatically attached to the inspected view, so you
// only need to handle this signal if you want to attach the inspector view
// yourself (for example, to add the inspector view to a browser tab).
//
// To prevent the inspector view from being attached you can connect to this
// signal and simply return TRUE.
func (inspector *WebInspector) ConnectAttach(f func() bool) externglib.SignalHandle {
	return inspector.Connect("attach", f)
}

// ConnectBringToFront: emitted when the inspector should be shown.
//
// If the inspector is not attached the inspector window should be shown on top
// of any other windows. If the inspector is attached the inspector view should
// be made visible. For example, if the inspector view is attached using a tab
// in a browser window, the browser window should be raised and the tab
// containing the inspector view should be the active one. In both cases, if
// this signal is not handled, the default implementation calls
// gtk_window_present() on the current toplevel Window of the inspector view.
func (inspector *WebInspector) ConnectBringToFront(f func() bool) externglib.SignalHandle {
	return inspector.Connect("bring-to-front", f)
}

// ConnectClosed: emitted when the inspector page is closed. If you are using
// your own inspector window, you should connect to this signal and destroy your
// window.
func (inspector *WebInspector) ConnectClosed(f func()) externglib.SignalHandle {
	return inspector.Connect("closed", f)
}

// ConnectDetach: emitted when the inspector is requested to be detached from
// the window it is currently attached to. The inspector is detached when the
// inspector page is about to be closed, and this signal is emitted right before
// KitWebInspector::closed, or when the user clicks on the detach button in the
// inspector view to show the inspector in a separate window. In this case the
// signal KitWebInspector::open-window is emitted after this one.
//
// To prevent the inspector view from being detached you can connect to this
// signal and simply return TRUE.
func (inspector *WebInspector) ConnectDetach(f func() bool) externglib.SignalHandle {
	return inspector.Connect("detach", f)
}

// ConnectOpenWindow: emitted when the inspector is requested to open in a
// separate window. If this signal is not handled, a Window with the inspector
// will be created and shown, so you only need to handle this signal if you want
// to use your own window. This signal is emitted after KitWebInspector::detach
// to show the inspector in a separate window after being detached.
//
// To prevent the inspector from being shown you can connect to this signal and
// simply return TRUE.
func (inspector *WebInspector) ConnectOpenWindow(f func() bool) externglib.SignalHandle {
	return inspector.Connect("open-window", f)
}
