// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"context"
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit/webkit.h>
// extern void _gotk4_webkit6_FaviconDatabase_ConnectFaviconChanged(gpointer, gchar*, gchar*, guintptr);
// extern void _gotk4_webkit6_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// GType values.
var (
	GTypeFaviconDatabaseError = coreglib.Type(C.webkit_favicon_database_error_get_type())
	GTypeFaviconDatabase      = coreglib.Type(C.webkit_favicon_database_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFaviconDatabaseError, F: marshalFaviconDatabaseError},
		coreglib.TypeMarshaler{T: GTypeFaviconDatabase, F: marshalFaviconDatabase},
	})
}

// FaviconDatabaseError: enum values used to denote the various errors related
// to the KitFaviconDatabase.
type FaviconDatabaseError C.gint

const (
	// FaviconDatabaseErrorNotInitialized is closed.
	FaviconDatabaseErrorNotInitialized FaviconDatabaseError = iota
	// FaviconDatabaseErrorFaviconNotFound: there is not an icon available for
	// the requested URL.
	FaviconDatabaseErrorFaviconNotFound
	// FaviconDatabaseErrorFaviconUnknown: there might be an icon for the
	// requested URL, but its data is unknown at the moment.
	FaviconDatabaseErrorFaviconUnknown
)

func marshalFaviconDatabaseError(p uintptr) (interface{}, error) {
	return FaviconDatabaseError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for FaviconDatabaseError.
func (f FaviconDatabaseError) String() string {
	switch f {
	case FaviconDatabaseErrorNotInitialized:
		return "NotInitialized"
	case FaviconDatabaseErrorFaviconNotFound:
		return "FaviconNotFound"
	case FaviconDatabaseErrorFaviconUnknown:
		return "FaviconUnknown"
	default:
		return fmt.Sprintf("FaviconDatabaseError(%d)", f)
	}
}

// FaviconDatabaseErrorQuark gets the quark for the domain of favicon database
// errors.
//
// The function returns the following values:
//
//   - quark: favicon database error domain.
//
func FaviconDatabaseErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.webkit_favicon_database_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)
	type _ = glib.Quark
	type _ = uint32

	return _quark
}

// FaviconDatabaseOverrides contains methods that are overridable.
type FaviconDatabaseOverrides struct {
}

func defaultFaviconDatabaseOverrides(v *FaviconDatabase) FaviconDatabaseOverrides {
	return FaviconDatabaseOverrides{}
}

// FaviconDatabase provides access to the icons associated with web sites.
//
// WebKit will automatically look for available icons in <link> elements on
// opened pages as well as an existing favicon.ico and load the images found
// into a memory cache if possible. That cache is frozen to an on-disk database
// for persistence.
//
// If KitSettings:enable-private-browsing is TRUE, new icons won't be added
// to the on-disk database and no existing icons will be deleted from it.
// Nevertheless, WebKit will still store them in the in-memory cache during the
// current execution.
type FaviconDatabase struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*FaviconDatabase)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*FaviconDatabase, *FaviconDatabaseClass, FaviconDatabaseOverrides](
		GTypeFaviconDatabase,
		initFaviconDatabaseClass,
		wrapFaviconDatabase,
		defaultFaviconDatabaseOverrides,
	)
}

func initFaviconDatabaseClass(gclass unsafe.Pointer, overrides FaviconDatabaseOverrides, classInitFunc func(*FaviconDatabaseClass)) {
	if classInitFunc != nil {
		class := (*FaviconDatabaseClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapFaviconDatabase(obj *coreglib.Object) *FaviconDatabase {
	return &FaviconDatabase{
		Object: obj,
	}
}

func marshalFaviconDatabase(p uintptr) (interface{}, error) {
	return wrapFaviconDatabase(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectFaviconChanged: this signal is emitted when the favicon URI of
// page_uri has been changed to favicon_uri in the database. You can connect
// to this signal and call webkit_favicon_database_get_favicon() to get the
// favicon. If you are interested in the favicon of a KitWebView it's easier to
// use the KitWebView:favicon property. See webkit_web_view_get_favicon() for
// more details.
func (database *FaviconDatabase) ConnectFaviconChanged(f func(pageUri, faviconUri string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(database, "favicon-changed", false, unsafe.Pointer(C._gotk4_webkit6_FaviconDatabase_ConnectFaviconChanged), f)
}

// Clear clears all icons from the database.
func (database *FaviconDatabase) Clear() {
	var _arg0 *C.WebKitFaviconDatabase // out

	_arg0 = (*C.WebKitFaviconDatabase)(unsafe.Pointer(coreglib.InternObject(database).Native()))

	C.webkit_favicon_database_clear(_arg0)
	runtime.KeepAlive(database)
}

// Favicon: asynchronously obtains a favicon image.
//
// Asynchronously obtains an image of the favicon for the given page URI.
// It returns the cached icon if it's in the database asynchronously waiting for
// the icon to be read from the database.
//
// This is an asynchronous method. When the operation is finished, callback will
// be invoked. You can then call webkit_favicon_database_get_favicon_finish() to
// get the result of the operation.
//
// The function takes the following parameters:
//
//   - ctx (optional) or NULL.
//   - pageUri: URI of the page for which we want to retrieve the favicon.
//   - callback (optional) to call when the request is satisfied or NULL if you
//     don't care about the result.
//
func (database *FaviconDatabase) Favicon(ctx context.Context, pageUri string, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitFaviconDatabase // out
	var _arg2 *C.GCancellable          // out
	var _arg1 *C.gchar                 // out
	var _arg3 C.GAsyncReadyCallback    // out
	var _arg4 C.gpointer

	_arg0 = (*C.WebKitFaviconDatabase)(unsafe.Pointer(coreglib.InternObject(database).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(pageUri)))
	defer C.free(unsafe.Pointer(_arg1))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.webkit_favicon_database_get_favicon(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(database)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(pageUri)
	runtime.KeepAlive(callback)
}

// FaviconFinish finishes an operation started with
// webkit_favicon_database_get_favicon().
//
// The function takes the following parameters:
//
//   - result obtained from the ReadyCallback passed to
//     webkit_favicon_database_get_favicon().
//
// The function returns the following values:
//
//   - texture: new favicon image, or NULL in case of error.
//
func (database *FaviconDatabase) FaviconFinish(result gio.AsyncResulter) (gdk.Texturer, error) {
	var _arg0 *C.WebKitFaviconDatabase // out
	var _arg1 *C.GAsyncResult          // out
	var _cret *C.GdkTexture            // in
	var _cerr *C.GError                // in

	_arg0 = (*C.WebKitFaviconDatabase)(unsafe.Pointer(coreglib.InternObject(database).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_cret = C.webkit_favicon_database_get_favicon_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(database)
	runtime.KeepAlive(result)

	var _texture gdk.Texturer // out
	var _goerr error          // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Texturer is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.Texturer)
			return ok
		})
		rv, ok := casted.(gdk.Texturer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Texturer")
		}
		_texture = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _texture, _goerr
}

// FaviconURI obtains the URI of the favicon for the given page_uri.
//
// The function takes the following parameters:
//
//   - pageUri: URI of the page containing the icon.
//
// The function returns the following values:
//
//   - utf8: newly allocated URI for the favicon, or NULL if the database
//     doesn't have a favicon for page_uri.
//
func (database *FaviconDatabase) FaviconURI(pageUri string) string {
	var _arg0 *C.WebKitFaviconDatabase // out
	var _arg1 *C.gchar                 // out
	var _cret *C.gchar                 // in

	_arg0 = (*C.WebKitFaviconDatabase)(unsafe.Pointer(coreglib.InternObject(database).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(pageUri)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.webkit_favicon_database_get_favicon_uri(_arg0, _arg1)
	runtime.KeepAlive(database)
	runtime.KeepAlive(pageUri)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FaviconDatabaseClass: instance of this type is always passed by reference.
type FaviconDatabaseClass struct {
	*faviconDatabaseClass
}

// faviconDatabaseClass is the struct that's finalized.
type faviconDatabaseClass struct {
	native *C.WebKitFaviconDatabaseClass
}
