// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: webkit2gtk-4.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_window_properties_get_type()), F: marshalWindowPropertieser},
	})
}

type WindowProperties struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*WindowProperties)(nil)
)

func wrapWindowProperties(obj *externglib.Object) *WindowProperties {
	return &WindowProperties{
		Object: obj,
	}
}

func marshalWindowPropertieser(p uintptr) (interface{}, error) {
	return wrapWindowProperties(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Fullscreen: get whether the window should be shown in fullscreen state or
// not.
func (windowProperties *WindowProperties) Fullscreen() bool {
	var _arg0 *C.WebKitWindowProperties // out
	var _cret C.gboolean                // in

	_arg0 = (*C.WebKitWindowProperties)(unsafe.Pointer(windowProperties.Native()))

	_cret = C.webkit_window_properties_get_fullscreen(_arg0)
	runtime.KeepAlive(windowProperties)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Geometry: get the geometry the window should have on the screen when shown.
func (windowProperties *WindowProperties) Geometry() *gdk.Rectangle {
	var _arg0 *C.WebKitWindowProperties // out
	var _arg1 C.GdkRectangle            // in

	_arg0 = (*C.WebKitWindowProperties)(unsafe.Pointer(windowProperties.Native()))

	C.webkit_window_properties_get_geometry(_arg0, &_arg1)
	runtime.KeepAlive(windowProperties)

	var _geometry *gdk.Rectangle // out

	_geometry = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _geometry
}

// LocationbarVisible: get whether the window should have the locationbar
// visible or not.
func (windowProperties *WindowProperties) LocationbarVisible() bool {
	var _arg0 *C.WebKitWindowProperties // out
	var _cret C.gboolean                // in

	_arg0 = (*C.WebKitWindowProperties)(unsafe.Pointer(windowProperties.Native()))

	_cret = C.webkit_window_properties_get_locationbar_visible(_arg0)
	runtime.KeepAlive(windowProperties)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MenubarVisible: get whether the window should have the menubar visible or
// not.
func (windowProperties *WindowProperties) MenubarVisible() bool {
	var _arg0 *C.WebKitWindowProperties // out
	var _cret C.gboolean                // in

	_arg0 = (*C.WebKitWindowProperties)(unsafe.Pointer(windowProperties.Native()))

	_cret = C.webkit_window_properties_get_menubar_visible(_arg0)
	runtime.KeepAlive(windowProperties)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Resizable: get whether the window should be resizable by the user or not.
func (windowProperties *WindowProperties) Resizable() bool {
	var _arg0 *C.WebKitWindowProperties // out
	var _cret C.gboolean                // in

	_arg0 = (*C.WebKitWindowProperties)(unsafe.Pointer(windowProperties.Native()))

	_cret = C.webkit_window_properties_get_resizable(_arg0)
	runtime.KeepAlive(windowProperties)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ScrollbarsVisible: get whether the window should have the scrollbars visible
// or not.
func (windowProperties *WindowProperties) ScrollbarsVisible() bool {
	var _arg0 *C.WebKitWindowProperties // out
	var _cret C.gboolean                // in

	_arg0 = (*C.WebKitWindowProperties)(unsafe.Pointer(windowProperties.Native()))

	_cret = C.webkit_window_properties_get_scrollbars_visible(_arg0)
	runtime.KeepAlive(windowProperties)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StatusbarVisible: get whether the window should have the statusbar visible or
// not.
func (windowProperties *WindowProperties) StatusbarVisible() bool {
	var _arg0 *C.WebKitWindowProperties // out
	var _cret C.gboolean                // in

	_arg0 = (*C.WebKitWindowProperties)(unsafe.Pointer(windowProperties.Native()))

	_cret = C.webkit_window_properties_get_statusbar_visible(_arg0)
	runtime.KeepAlive(windowProperties)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ToolbarVisible: get whether the window should have the toolbar visible or
// not.
func (windowProperties *WindowProperties) ToolbarVisible() bool {
	var _arg0 *C.WebKitWindowProperties // out
	var _cret C.gboolean                // in

	_arg0 = (*C.WebKitWindowProperties)(unsafe.Pointer(windowProperties.Native()))

	_cret = C.webkit_window_properties_get_toolbar_visible(_arg0)
	runtime.KeepAlive(windowProperties)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
