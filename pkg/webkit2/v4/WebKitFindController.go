// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"fmt"
	"runtime"
	"strings"
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
		{T: externglib.Type(C.webkit_find_options_get_type()), F: marshalFindOptions},
		{T: externglib.Type(C.webkit_find_controller_get_type()), F: marshalFindControllerer},
	})
}

// FindOptions: enum values used to specify search options.
type FindOptions C.guint

const (
	// FindOptionsNone: no search flags, this means a case sensitive, no wrap,
	// forward only search.
	FindOptionsNone FindOptions = 0b0
	// FindOptionsCaseInsensitive: case insensitive search.
	FindOptionsCaseInsensitive FindOptions = 0b1
	// FindOptionsAtWordStarts: search text only at the begining of the words.
	FindOptionsAtWordStarts FindOptions = 0b10
	// FindOptionsTreatMedialCapitalAsWordStart: treat capital letters in the
	// middle of words as word start.
	FindOptionsTreatMedialCapitalAsWordStart FindOptions = 0b100
	// FindOptionsBackwards: search backwards.
	FindOptionsBackwards FindOptions = 0b1000
	// FindOptionsWrapAround: if not present search will stop at the end of the
	// document.
	FindOptionsWrapAround FindOptions = 0b10000
)

func marshalFindOptions(p uintptr) (interface{}, error) {
	return FindOptions(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for FindOptions.
func (f FindOptions) String() string {
	if f == 0 {
		return "FindOptions(0)"
	}

	var builder strings.Builder
	builder.Grow(150)

	for f != 0 {
		next := f & (f - 1)
		bit := f - next

		switch bit {
		case FindOptionsNone:
			builder.WriteString("None|")
		case FindOptionsCaseInsensitive:
			builder.WriteString("CaseInsensitive|")
		case FindOptionsAtWordStarts:
			builder.WriteString("AtWordStarts|")
		case FindOptionsTreatMedialCapitalAsWordStart:
			builder.WriteString("TreatMedialCapitalAsWordStart|")
		case FindOptionsBackwards:
			builder.WriteString("Backwards|")
		case FindOptionsWrapAround:
			builder.WriteString("WrapAround|")
		default:
			builder.WriteString(fmt.Sprintf("FindOptions(0b%b)|", bit))
		}

		f = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if f contains other.
func (f FindOptions) Has(other FindOptions) bool {
	return (f & other) == other
}

type FindController struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*FindController)(nil)
)

func wrapFindController(obj *externglib.Object) *FindController {
	return &FindController{
		Object: obj,
	}
}

func marshalFindControllerer(p uintptr) (interface{}, error) {
	return wrapFindController(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CountMatches counts the number of matches for search_text found in the
// KitWebView with the provided find_options. The number of matches will be
// provided by the KitFindController::counted-matches signal.
//
// The function takes the following parameters:
//
//    - searchText: text to look for.
//    - findOptions: bitmask with the KitFindOptions used in the search.
//    - maxMatchCount: maximum number of matches allowed in the search.
//
func (findController *FindController) CountMatches(searchText string, findOptions uint32, maxMatchCount uint) {
	var _arg0 *C.WebKitFindController // out
	var _arg1 *C.gchar                // out
	var _arg2 C.guint32               // out
	var _arg3 C.guint                 // out

	_arg0 = (*C.WebKitFindController)(unsafe.Pointer(findController.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(searchText)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint32(findOptions)
	_arg3 = C.guint(maxMatchCount)

	C.webkit_find_controller_count_matches(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(findController)
	runtime.KeepAlive(searchText)
	runtime.KeepAlive(findOptions)
	runtime.KeepAlive(maxMatchCount)
}

// MaxMatchCount gets the maximum number of matches to report during a text
// lookup. This number is passed as the last argument of
// webkit_find_controller_search() or webkit_find_controller_count_matches().
func (findController *FindController) MaxMatchCount() uint {
	var _arg0 *C.WebKitFindController // out
	var _cret C.guint                 // in

	_arg0 = (*C.WebKitFindController)(unsafe.Pointer(findController.Native()))

	_cret = C.webkit_find_controller_get_max_match_count(_arg0)
	runtime.KeepAlive(findController)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Options gets a bitmask containing the KitFindOptions associated with the
// current search.
func (findController *FindController) Options() uint32 {
	var _arg0 *C.WebKitFindController // out
	var _cret C.guint32               // in

	_arg0 = (*C.WebKitFindController)(unsafe.Pointer(findController.Native()))

	_cret = C.webkit_find_controller_get_options(_arg0)
	runtime.KeepAlive(findController)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// SearchText gets the text that find_controller is currently searching for.
// This text is passed to either webkit_find_controller_search() or
// webkit_find_controller_count_matches().
func (findController *FindController) SearchText() string {
	var _arg0 *C.WebKitFindController // out
	var _cret *C.gchar                // in

	_arg0 = (*C.WebKitFindController)(unsafe.Pointer(findController.Native()))

	_cret = C.webkit_find_controller_get_search_text(_arg0)
	runtime.KeepAlive(findController)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// WebView gets the KitWebView this find controller is associated to. Do not
// dereference the returned instance as it belongs to the KitFindController.
func (findController *FindController) WebView() *WebView {
	var _arg0 *C.WebKitFindController // out
	var _cret *C.WebKitWebView        // in

	_arg0 = (*C.WebKitFindController)(unsafe.Pointer(findController.Native()))

	_cret = C.webkit_find_controller_get_web_view(_arg0)
	runtime.KeepAlive(findController)

	var _webView *WebView // out

	_webView = wrapWebView(externglib.Take(unsafe.Pointer(_cret)))

	return _webView
}

// Search looks for search_text in the KitWebView associated with
// find_controller since the beginning of the document highlighting up to
// max_match_count matches. The outcome of the search will be asynchronously
// provided by the KitFindController::found-text and
// KitFindController::failed-to-find-text signals.
//
// To look for the next or previous occurrences of the same text with the same
// find options use webkit_find_controller_search_next() and/or
// webkit_find_controller_search_previous(). The KitFindController will use the
// same text and options for the following searches unless they are modified by
// another call to this method.
//
// Note that if the number of matches is higher than max_match_count then
// KitFindController::found-text will report G_MAXUINT matches instead of the
// actual number.
//
// Callers should call webkit_find_controller_search_finish() to finish the
// current search operation.
//
// The function takes the following parameters:
//
//    - searchText: text to look for.
//    - findOptions: bitmask with the KitFindOptions used in the search.
//    - maxMatchCount: maximum number of matches allowed in the search.
//
func (findController *FindController) Search(searchText string, findOptions uint32, maxMatchCount uint) {
	var _arg0 *C.WebKitFindController // out
	var _arg1 *C.gchar                // out
	var _arg2 C.guint32               // out
	var _arg3 C.guint                 // out

	_arg0 = (*C.WebKitFindController)(unsafe.Pointer(findController.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(searchText)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint32(findOptions)
	_arg3 = C.guint(maxMatchCount)

	C.webkit_find_controller_search(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(findController)
	runtime.KeepAlive(searchText)
	runtime.KeepAlive(findOptions)
	runtime.KeepAlive(maxMatchCount)
}

// SearchFinish finishes a find operation started by
// webkit_find_controller_search(). It will basically unhighlight every text
// match found.
//
// This method will be typically called when the search UI is closed/hidden by
// the client application.
func (findController *FindController) SearchFinish() {
	var _arg0 *C.WebKitFindController // out

	_arg0 = (*C.WebKitFindController)(unsafe.Pointer(findController.Native()))

	C.webkit_find_controller_search_finish(_arg0)
	runtime.KeepAlive(findController)
}

// SearchNext looks for the next occurrence of the search text.
//
// Calling this method before webkit_find_controller_search() or
// webkit_find_controller_count_matches() is a programming error.
func (findController *FindController) SearchNext() {
	var _arg0 *C.WebKitFindController // out

	_arg0 = (*C.WebKitFindController)(unsafe.Pointer(findController.Native()))

	C.webkit_find_controller_search_next(_arg0)
	runtime.KeepAlive(findController)
}

// SearchPrevious looks for the previous occurrence of the search text.
//
// Calling this method before webkit_find_controller_search() or
// webkit_find_controller_count_matches() is a programming error.
func (findController *FindController) SearchPrevious() {
	var _arg0 *C.WebKitFindController // out

	_arg0 = (*C.WebKitFindController)(unsafe.Pointer(findController.Native()))

	C.webkit_find_controller_search_previous(_arg0)
	runtime.KeepAlive(findController)
}

// ConnectCountedMatches: this signal is emitted when the KitFindController has
// counted the number of matches for a given text after a call to
// webkit_find_controller_count_matches().
func (findController *FindController) ConnectCountedMatches(f func(matchCount uint)) externglib.SignalHandle {
	return findController.Connect("counted-matches", f)
}

// ConnectFailedToFindText: this signal is emitted when a search operation does
// not find any result for the given text. It will be issued if the text is not
// found asynchronously after a call to webkit_find_controller_search(),
// webkit_find_controller_search_next() or
// webkit_find_controller_search_previous().
func (findController *FindController) ConnectFailedToFindText(f func()) externglib.SignalHandle {
	return findController.Connect("failed-to-find-text", f)
}

// ConnectFoundText: this signal is emitted when a given text is found in the
// web page text. It will be issued if the text is found asynchronously after a
// call to webkit_find_controller_search(), webkit_find_controller_search_next()
// or webkit_find_controller_search_previous().
func (findController *FindController) ConnectFoundText(f func(matchCount uint)) externglib.SignalHandle {
	return findController.Connect("found-text", f)
}
