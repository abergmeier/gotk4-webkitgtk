// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: webkit2gtk-4.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_user_content_filter_store_get_type()), F: marshalUserContentFilterStorer},
	})
}

type UserContentFilterStore struct {
	*externglib.Object
}

var _ gextras.Nativer = (*UserContentFilterStore)(nil)

func wrapUserContentFilterStore(obj *externglib.Object) *UserContentFilterStore {
	return &UserContentFilterStore{
		Object: obj,
	}
}

func marshalUserContentFilterStorer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapUserContentFilterStore(obj), nil
}

// NewUserContentFilterStore: create a new KitUserContentFilterStore to
// manipulate filters stored at storage_path. The path must point to a local
// filesystem, and will be created if needed.
func NewUserContentFilterStore(storagePath string) *UserContentFilterStore {
	var _arg1 *C.gchar                        // out
	var _cret *C.WebKitUserContentFilterStore // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(storagePath)))

	_cret = C.webkit_user_content_filter_store_new(_arg1)

	var _userContentFilterStore *UserContentFilterStore // out

	_userContentFilterStore = wrapUserContentFilterStore(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _userContentFilterStore
}

// FetchIdentifiers: asynchronously retrieve a list of the identifiers for all
// the stored filters.
//
// When the operation is finished, callback will be invoked, which then can use
// webkit_user_content_filter_store_fetch_identifiers_finish() to obtain the
// list of filter identifiers.
func (store *UserContentFilterStore) FetchIdentifiers(cancellable *gio.Cancellable, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitUserContentFilterStore // out
	var _arg1 *C.GCancellable                 // out
	var _arg2 C.GAsyncReadyCallback           // out
	var _arg3 C.gpointer

	_arg0 = (*C.WebKitUserContentFilterStore)(unsafe.Pointer(store.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg3 = C.gpointer(gbox.AssignOnce(callback))

	C.webkit_user_content_filter_store_fetch_identifiers(_arg0, _arg1, _arg2, _arg3)
}

// FetchIdentifiersFinish finishes an asynchronous fetch of the list of
// identifiers for the stored filters previously started with
// webkit_user_content_filter_store_fetch_identifiers().
func (store *UserContentFilterStore) FetchIdentifiersFinish(result gio.AsyncResulter) []string {
	var _arg0 *C.WebKitUserContentFilterStore // out
	var _arg1 *C.GAsyncResult                 // out
	var _cret **C.gchar

	_arg0 = (*C.WebKitUserContentFilterStore)(unsafe.Pointer(store.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.webkit_user_content_filter_store_fetch_identifiers_finish(_arg0, _arg1)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

func (store *UserContentFilterStore) Path() string {
	var _arg0 *C.WebKitUserContentFilterStore // out
	var _cret *C.gchar                        // in

	_arg0 = (*C.WebKitUserContentFilterStore)(unsafe.Pointer(store.Native()))

	_cret = C.webkit_user_content_filter_store_get_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Load: asynchronously load a content filter given its identifier. The filter
// must have been previously stored using
// webkit_user_content_filter_store_save().
//
// When the operation is finished, callback will be invoked, which then can use
// webkit_user_content_filter_store_load_finish() to obtain the resulting
// filter.
func (store *UserContentFilterStore) Load(identifier string, cancellable *gio.Cancellable, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitUserContentFilterStore // out
	var _arg1 *C.gchar                        // out
	var _arg2 *C.GCancellable                 // out
	var _arg3 C.GAsyncReadyCallback           // out
	var _arg4 C.gpointer

	_arg0 = (*C.WebKitUserContentFilterStore)(unsafe.Pointer(store.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(identifier)))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg4 = C.gpointer(gbox.AssignOnce(callback))

	C.webkit_user_content_filter_store_load(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// LoadFinish finishes an asynchronous filter load previously started with
// webkit_user_content_filter_store_load().
func (store *UserContentFilterStore) LoadFinish(result gio.AsyncResulter) (*UserContentFilter, error) {
	var _arg0 *C.WebKitUserContentFilterStore // out
	var _arg1 *C.GAsyncResult                 // out
	var _cret *C.WebKitUserContentFilter      // in
	var _cerr *C.GError                       // in

	_arg0 = (*C.WebKitUserContentFilterStore)(unsafe.Pointer(store.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.webkit_user_content_filter_store_load_finish(_arg0, _arg1, &_cerr)

	var _userContentFilter *UserContentFilter // out
	var _goerr error                          // out

	_userContentFilter = (*UserContentFilter)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.webkit_user_content_filter_ref(_cret)
	runtime.SetFinalizer(_userContentFilter, func(v *UserContentFilter) {
		C.webkit_user_content_filter_unref((*C.WebKitUserContentFilter)(gextras.StructNative(unsafe.Pointer(v))))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _userContentFilter, _goerr
}

// Remove: asynchronously remove a content filter given its identifier.
//
// When the operation is finished, callback will be invoked, which then can use
// webkit_user_content_filter_store_remove_finish() to check whether the removal
// was successful.
func (store *UserContentFilterStore) Remove(identifier string, cancellable *gio.Cancellable, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitUserContentFilterStore // out
	var _arg1 *C.gchar                        // out
	var _arg2 *C.GCancellable                 // out
	var _arg3 C.GAsyncReadyCallback           // out
	var _arg4 C.gpointer

	_arg0 = (*C.WebKitUserContentFilterStore)(unsafe.Pointer(store.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(identifier)))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg4 = C.gpointer(gbox.AssignOnce(callback))

	C.webkit_user_content_filter_store_remove(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// RemoveFinish finishes an asynchronous filter removal previously started with
// webkit_user_content_filter_store_remove().
func (store *UserContentFilterStore) RemoveFinish(result gio.AsyncResulter) error {
	var _arg0 *C.WebKitUserContentFilterStore // out
	var _arg1 *C.GAsyncResult                 // out
	var _cerr *C.GError                       // in

	_arg0 = (*C.WebKitUserContentFilterStore)(unsafe.Pointer(store.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	C.webkit_user_content_filter_store_remove_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SaveFinish finishes an asynchronous filter save previously started with
// webkit_user_content_filter_store_save().
func (store *UserContentFilterStore) SaveFinish(result gio.AsyncResulter) (*UserContentFilter, error) {
	var _arg0 *C.WebKitUserContentFilterStore // out
	var _arg1 *C.GAsyncResult                 // out
	var _cret *C.WebKitUserContentFilter      // in
	var _cerr *C.GError                       // in

	_arg0 = (*C.WebKitUserContentFilterStore)(unsafe.Pointer(store.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.webkit_user_content_filter_store_save_finish(_arg0, _arg1, &_cerr)

	var _userContentFilter *UserContentFilter // out
	var _goerr error                          // out

	_userContentFilter = (*UserContentFilter)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.webkit_user_content_filter_ref(_cret)
	runtime.SetFinalizer(_userContentFilter, func(v *UserContentFilter) {
		C.webkit_user_content_filter_unref((*C.WebKitUserContentFilter)(gextras.StructNative(unsafe.Pointer(v))))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _userContentFilter, _goerr
}

// SaveFromFile: asynchronously save a content filter from the contents of a
// file, which must be native to the platform, as checked by g_file_is_native().
// See webkit_user_content_filter_store_save() for more details.
//
// When the operation is finished, callback will be invoked, which then can use
// webkit_user_content_filter_store_save_finish() to obtain the resulting
// filter.
func (store *UserContentFilterStore) SaveFromFile(identifier string, file gio.Filer, cancellable *gio.Cancellable, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitUserContentFilterStore // out
	var _arg1 *C.gchar                        // out
	var _arg2 *C.GFile                        // out
	var _arg3 *C.GCancellable                 // out
	var _arg4 C.GAsyncReadyCallback           // out
	var _arg5 C.gpointer

	_arg0 = (*C.WebKitUserContentFilterStore)(unsafe.Pointer(store.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(identifier)))
	_arg2 = (*C.GFile)(unsafe.Pointer((file).(gextras.Nativer).Native()))
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg5 = C.gpointer(gbox.AssignOnce(callback))

	C.webkit_user_content_filter_store_save_from_file(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// SaveFromFileFinish finishes and asynchronous filter save previously started
// with webkit_user_content_filter_store_save_from_file().
func (store *UserContentFilterStore) SaveFromFileFinish(result gio.AsyncResulter) (*UserContentFilter, error) {
	var _arg0 *C.WebKitUserContentFilterStore // out
	var _arg1 *C.GAsyncResult                 // out
	var _cret *C.WebKitUserContentFilter      // in
	var _cerr *C.GError                       // in

	_arg0 = (*C.WebKitUserContentFilterStore)(unsafe.Pointer(store.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.webkit_user_content_filter_store_save_from_file_finish(_arg0, _arg1, &_cerr)

	var _userContentFilter *UserContentFilter // out
	var _goerr error                          // out

	_userContentFilter = (*UserContentFilter)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.webkit_user_content_filter_ref(_cret)
	runtime.SetFinalizer(_userContentFilter, func(v *UserContentFilter) {
		C.webkit_user_content_filter_unref((*C.WebKitUserContentFilter)(gextras.StructNative(unsafe.Pointer(v))))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _userContentFilter, _goerr
}
