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
		{T: externglib.Type(C.webkit_notification_get_type()), F: marshalNotificationer},
	})
}

type Notification struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*Notification)(nil)
)

func wrapNotification(obj *externglib.Object) *Notification {
	return &Notification{
		Object: obj,
	}
}

func marshalNotificationer(p uintptr) (interface{}, error) {
	return wrapNotification(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Clicked tells WebKit the notification has been clicked. This will emit the
// KitNotification::clicked signal.
func (notification *Notification) Clicked() {
	var _arg0 *C.WebKitNotification // out

	_arg0 = (*C.WebKitNotification)(unsafe.Pointer(notification.Native()))

	C.webkit_notification_clicked(_arg0)
	runtime.KeepAlive(notification)
}

// Close closes the notification.
func (notification *Notification) Close() {
	var _arg0 *C.WebKitNotification // out

	_arg0 = (*C.WebKitNotification)(unsafe.Pointer(notification.Native()))

	C.webkit_notification_close(_arg0)
	runtime.KeepAlive(notification)
}

// Body obtains the body for the notification.
func (notification *Notification) Body() string {
	var _arg0 *C.WebKitNotification // out
	var _cret *C.gchar              // in

	_arg0 = (*C.WebKitNotification)(unsafe.Pointer(notification.Native()))

	_cret = C.webkit_notification_get_body(_arg0)
	runtime.KeepAlive(notification)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ID obtains the unique id for the notification.
func (notification *Notification) ID() uint64 {
	var _arg0 *C.WebKitNotification // out
	var _cret C.guint64             // in

	_arg0 = (*C.WebKitNotification)(unsafe.Pointer(notification.Native()))

	_cret = C.webkit_notification_get_id(_arg0)
	runtime.KeepAlive(notification)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// Tag obtains the tag identifier for the notification.
func (notification *Notification) Tag() string {
	var _arg0 *C.WebKitNotification // out
	var _cret *C.gchar              // in

	_arg0 = (*C.WebKitNotification)(unsafe.Pointer(notification.Native()))

	_cret = C.webkit_notification_get_tag(_arg0)
	runtime.KeepAlive(notification)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Title obtains the title for the notification.
func (notification *Notification) Title() string {
	var _arg0 *C.WebKitNotification // out
	var _cret *C.gchar              // in

	_arg0 = (*C.WebKitNotification)(unsafe.Pointer(notification.Native()))

	_cret = C.webkit_notification_get_title(_arg0)
	runtime.KeepAlive(notification)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ConnectClicked: emitted when a notification has been clicked. See
// webkit_notification_clicked().
func (notification *Notification) ConnectClicked(f func()) externglib.SignalHandle {
	return notification.Connect("clicked", f)
}

// ConnectClosed: emitted when a notification has been withdrawn.
//
// The default handler will close the notification using libnotify, if built
// with support for it.
func (notification *Notification) ConnectClosed(f func()) externglib.SignalHandle {
	return notification.Connect("closed", f)
}
