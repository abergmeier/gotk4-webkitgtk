// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.webkit_user_content_injected_frames_get_type()), F: marshalUserContentInjectedFrames},
		{T: externglib.Type(C.webkit_user_script_injection_time_get_type()), F: marshalUserScriptInjectionTime},
		{T: externglib.Type(C.webkit_user_style_level_get_type()), F: marshalUserStyleLevel},
		{T: externglib.Type(C.webkit_user_content_filter_get_type()), F: marshalUserContentFilter},
		{T: externglib.Type(C.webkit_user_script_get_type()), F: marshalUserScript},
		{T: externglib.Type(C.webkit_user_style_sheet_get_type()), F: marshalUserStyleSheet},
	})
}

// UserContentInjectedFrames specifies in which frames user style sheets are to
// be inserted in.
type UserContentInjectedFrames C.gint

const (
	// UserContentInjectAllFrames: insert the user style sheet in all the frames
	// loaded by the web view, including nested frames. This is the default.
	UserContentInjectAllFrames UserContentInjectedFrames = iota
	// UserContentInjectTopFrame: insert the user style sheet *only* in the
	// top-level frame loaded by the web view, and *not* in the nested frames.
	UserContentInjectTopFrame
)

func marshalUserContentInjectedFrames(p uintptr) (interface{}, error) {
	return UserContentInjectedFrames(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for UserContentInjectedFrames.
func (u UserContentInjectedFrames) String() string {
	switch u {
	case UserContentInjectAllFrames:
		return "AllFrames"
	case UserContentInjectTopFrame:
		return "TopFrame"
	default:
		return fmt.Sprintf("UserContentInjectedFrames(%d)", u)
	}
}

// UserScriptInjectionTime specifies at which place of documents an user script
// will be inserted.
type UserScriptInjectionTime C.gint

const (
	// UserScriptInjectAtDocumentStart: insert the code of the user script at
	// the beginning of loaded documents. This is the default.
	UserScriptInjectAtDocumentStart UserScriptInjectionTime = iota
	// UserScriptInjectAtDocumentEnd: insert the code of the user script at the
	// end of the loaded documents.
	UserScriptInjectAtDocumentEnd
)

func marshalUserScriptInjectionTime(p uintptr) (interface{}, error) {
	return UserScriptInjectionTime(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for UserScriptInjectionTime.
func (u UserScriptInjectionTime) String() string {
	switch u {
	case UserScriptInjectAtDocumentStart:
		return "Start"
	case UserScriptInjectAtDocumentEnd:
		return "End"
	default:
		return fmt.Sprintf("UserScriptInjectionTime(%d)", u)
	}
}

// UserStyleLevel specifies how to treat an user style sheet.
type UserStyleLevel C.gint

const (
	// UserStyleLevelUser: style sheet is an user style sheet, its contents
	// always override other style sheets. This is the default.
	UserStyleLevelUser UserStyleLevel = iota
	// UserStyleLevelAuthor: style sheet will be treated as if it was provided
	// by the loaded documents. That means other user style sheets may still
	// override it.
	UserStyleLevelAuthor
)

func marshalUserStyleLevel(p uintptr) (interface{}, error) {
	return UserStyleLevel(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for UserStyleLevel.
func (u UserStyleLevel) String() string {
	switch u {
	case UserStyleLevelUser:
		return "User"
	case UserStyleLevelAuthor:
		return "Author"
	default:
		return fmt.Sprintf("UserStyleLevel(%d)", u)
	}
}

// UserContentFilter: instance of this type is always passed by reference.
type UserContentFilter struct {
	*userContentFilter
}

// userContentFilter is the struct that's finalized.
type userContentFilter struct {
	native *C.WebKitUserContentFilter
}

func marshalUserContentFilter(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &UserContentFilter{&userContentFilter{(*C.WebKitUserContentFilter)(b)}}, nil
}

// Identifier: obtain the identifier previously used to save the
// user_content_filter in the KitUserContentFilterStore.
func (userContentFilter *UserContentFilter) Identifier() string {
	var _arg0 *C.WebKitUserContentFilter // out
	var _cret *C.char                    // in

	_arg0 = (*C.WebKitUserContentFilter)(gextras.StructNative(unsafe.Pointer(userContentFilter)))

	_cret = C.webkit_user_content_filter_get_identifier(_arg0)
	runtime.KeepAlive(userContentFilter)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UserScript: instance of this type is always passed by reference.
type UserScript struct {
	*userScript
}

// userScript is the struct that's finalized.
type userScript struct {
	native *C.WebKitUserScript
}

func marshalUserScript(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &UserScript{&userScript{(*C.WebKitUserScript)(b)}}, nil
}

// NewUserScript constructs a struct UserScript.
func NewUserScript(source string, injectedFrames UserContentInjectedFrames, injectionTime UserScriptInjectionTime, allowList []string, blockList []string) *UserScript {
	var _arg1 *C.gchar                          // out
	var _arg2 C.WebKitUserContentInjectedFrames // out
	var _arg3 C.WebKitUserScriptInjectionTime   // out
	var _arg4 **C.gchar                         // out
	var _arg5 **C.gchar                         // out
	var _cret *C.WebKitUserScript               // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(source)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.WebKitUserContentInjectedFrames(injectedFrames)
	_arg3 = C.WebKitUserScriptInjectionTime(injectionTime)
	{
		_arg4 = (**C.gchar)(C.malloc(C.size_t(uint((len(allowList) + 1)) * uint(unsafe.Sizeof(uint(0))))))
		defer C.free(unsafe.Pointer(_arg4))
		{
			out := unsafe.Slice(_arg4, len(allowList)+1)
			var zero *C.gchar
			out[len(allowList)] = zero
			for i := range allowList {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(allowList[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	{
		_arg5 = (**C.gchar)(C.malloc(C.size_t(uint((len(blockList) + 1)) * uint(unsafe.Sizeof(uint(0))))))
		defer C.free(unsafe.Pointer(_arg5))
		{
			out := unsafe.Slice(_arg5, len(blockList)+1)
			var zero *C.gchar
			out[len(blockList)] = zero
			for i := range blockList {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(blockList[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.webkit_user_script_new(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(source)
	runtime.KeepAlive(injectedFrames)
	runtime.KeepAlive(injectionTime)
	runtime.KeepAlive(allowList)
	runtime.KeepAlive(blockList)

	var _userScript *UserScript // out

	_userScript = (*UserScript)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_userScript)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_user_script_unref((*C.WebKitUserScript)(intern.C))
		},
	)

	return _userScript
}

// NewUserScriptForWorld constructs a struct UserScript.
func NewUserScriptForWorld(source string, injectedFrames UserContentInjectedFrames, injectionTime UserScriptInjectionTime, worldName string, allowList []string, blockList []string) *UserScript {
	var _arg1 *C.gchar                          // out
	var _arg2 C.WebKitUserContentInjectedFrames // out
	var _arg3 C.WebKitUserScriptInjectionTime   // out
	var _arg4 *C.gchar                          // out
	var _arg5 **C.gchar                         // out
	var _arg6 **C.gchar                         // out
	var _cret *C.WebKitUserScript               // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(source)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.WebKitUserContentInjectedFrames(injectedFrames)
	_arg3 = C.WebKitUserScriptInjectionTime(injectionTime)
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(worldName)))
	defer C.free(unsafe.Pointer(_arg4))
	{
		_arg5 = (**C.gchar)(C.malloc(C.size_t(uint((len(allowList) + 1)) * uint(unsafe.Sizeof(uint(0))))))
		defer C.free(unsafe.Pointer(_arg5))
		{
			out := unsafe.Slice(_arg5, len(allowList)+1)
			var zero *C.gchar
			out[len(allowList)] = zero
			for i := range allowList {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(allowList[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	{
		_arg6 = (**C.gchar)(C.malloc(C.size_t(uint((len(blockList) + 1)) * uint(unsafe.Sizeof(uint(0))))))
		defer C.free(unsafe.Pointer(_arg6))
		{
			out := unsafe.Slice(_arg6, len(blockList)+1)
			var zero *C.gchar
			out[len(blockList)] = zero
			for i := range blockList {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(blockList[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.webkit_user_script_new_for_world(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(source)
	runtime.KeepAlive(injectedFrames)
	runtime.KeepAlive(injectionTime)
	runtime.KeepAlive(worldName)
	runtime.KeepAlive(allowList)
	runtime.KeepAlive(blockList)

	var _userScript *UserScript // out

	_userScript = (*UserScript)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_userScript)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_user_script_unref((*C.WebKitUserScript)(intern.C))
		},
	)

	return _userScript
}

// UserStyleSheet: instance of this type is always passed by reference.
type UserStyleSheet struct {
	*userStyleSheet
}

// userStyleSheet is the struct that's finalized.
type userStyleSheet struct {
	native *C.WebKitUserStyleSheet
}

func marshalUserStyleSheet(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &UserStyleSheet{&userStyleSheet{(*C.WebKitUserStyleSheet)(b)}}, nil
}

// NewUserStyleSheet constructs a struct UserStyleSheet.
func NewUserStyleSheet(source string, injectedFrames UserContentInjectedFrames, level UserStyleLevel, allowList []string, blockList []string) *UserStyleSheet {
	var _arg1 *C.gchar                          // out
	var _arg2 C.WebKitUserContentInjectedFrames // out
	var _arg3 C.WebKitUserStyleLevel            // out
	var _arg4 **C.gchar                         // out
	var _arg5 **C.gchar                         // out
	var _cret *C.WebKitUserStyleSheet           // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(source)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.WebKitUserContentInjectedFrames(injectedFrames)
	_arg3 = C.WebKitUserStyleLevel(level)
	{
		_arg4 = (**C.gchar)(C.malloc(C.size_t(uint((len(allowList) + 1)) * uint(unsafe.Sizeof(uint(0))))))
		defer C.free(unsafe.Pointer(_arg4))
		{
			out := unsafe.Slice(_arg4, len(allowList)+1)
			var zero *C.gchar
			out[len(allowList)] = zero
			for i := range allowList {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(allowList[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	{
		_arg5 = (**C.gchar)(C.malloc(C.size_t(uint((len(blockList) + 1)) * uint(unsafe.Sizeof(uint(0))))))
		defer C.free(unsafe.Pointer(_arg5))
		{
			out := unsafe.Slice(_arg5, len(blockList)+1)
			var zero *C.gchar
			out[len(blockList)] = zero
			for i := range blockList {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(blockList[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.webkit_user_style_sheet_new(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(source)
	runtime.KeepAlive(injectedFrames)
	runtime.KeepAlive(level)
	runtime.KeepAlive(allowList)
	runtime.KeepAlive(blockList)

	var _userStyleSheet *UserStyleSheet // out

	_userStyleSheet = (*UserStyleSheet)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_userStyleSheet)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_user_style_sheet_unref((*C.WebKitUserStyleSheet)(intern.C))
		},
	)

	return _userStyleSheet
}

// NewUserStyleSheetForWorld constructs a struct UserStyleSheet.
func NewUserStyleSheetForWorld(source string, injectedFrames UserContentInjectedFrames, level UserStyleLevel, worldName string, allowList []string, blockList []string) *UserStyleSheet {
	var _arg1 *C.gchar                          // out
	var _arg2 C.WebKitUserContentInjectedFrames // out
	var _arg3 C.WebKitUserStyleLevel            // out
	var _arg4 *C.gchar                          // out
	var _arg5 **C.gchar                         // out
	var _arg6 **C.gchar                         // out
	var _cret *C.WebKitUserStyleSheet           // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(source)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.WebKitUserContentInjectedFrames(injectedFrames)
	_arg3 = C.WebKitUserStyleLevel(level)
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(worldName)))
	defer C.free(unsafe.Pointer(_arg4))
	{
		_arg5 = (**C.gchar)(C.malloc(C.size_t(uint((len(allowList) + 1)) * uint(unsafe.Sizeof(uint(0))))))
		defer C.free(unsafe.Pointer(_arg5))
		{
			out := unsafe.Slice(_arg5, len(allowList)+1)
			var zero *C.gchar
			out[len(allowList)] = zero
			for i := range allowList {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(allowList[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	{
		_arg6 = (**C.gchar)(C.malloc(C.size_t(uint((len(blockList) + 1)) * uint(unsafe.Sizeof(uint(0))))))
		defer C.free(unsafe.Pointer(_arg6))
		{
			out := unsafe.Slice(_arg6, len(blockList)+1)
			var zero *C.gchar
			out[len(blockList)] = zero
			for i := range blockList {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(blockList[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.webkit_user_style_sheet_new_for_world(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(source)
	runtime.KeepAlive(injectedFrames)
	runtime.KeepAlive(level)
	runtime.KeepAlive(worldName)
	runtime.KeepAlive(allowList)
	runtime.KeepAlive(blockList)

	var _userStyleSheet *UserStyleSheet // out

	_userStyleSheet = (*UserStyleSheet)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_userStyleSheet)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_user_style_sheet_unref((*C.WebKitUserStyleSheet)(intern.C))
		},
	)

	return _userStyleSheet
}
