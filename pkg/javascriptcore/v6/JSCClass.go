// Code generated by girgen. DO NOT EDIT.

package javascriptcore

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <jsc/jsc.h>
import "C"

// GType values.
var (
	GTypeClass = coreglib.Type(C.jsc_class_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeClass, F: marshalClass},
	})
}

// ClassOverrides contains methods that are overridable.
type ClassOverrides struct {
}

func defaultClassOverrides(v *Class) ClassOverrides {
	return ClassOverrides{}
}

// Class represents a custom JavaScript class registered by the user in a
// CContext. It allows to create new JavaScripts objects whose instances are
// created by the user using this API. It's possible to add constructors,
// properties and methods for a JSSClass by providing #GCallback<!-- -->s to
// implement them.
type Class struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Class)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Class, *ClassClass, ClassOverrides](
		GTypeClass,
		initClassClass,
		wrapClass,
		defaultClassOverrides,
	)
}

func initClassClass(gclass unsafe.Pointer, overrides ClassOverrides, classInitFunc func(*ClassClass)) {
	if classInitFunc != nil {
		class := (*ClassClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapClass(obj *coreglib.Object) *Class {
	return &Class{
		Object: obj,
	}
}

func marshalClass(p uintptr) (interface{}, error) {
	return wrapClass(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Name: get the class name of jsc_class.
//
// The function returns the following values:
//
//   - utf8: name of jsc_class.
//
func (jscClass *Class) Name() string {
	var _arg0 *C.JSCClass // out
	var _cret *C.char     // in

	_arg0 = (*C.JSCClass)(unsafe.Pointer(coreglib.InternObject(jscClass).Native()))

	_cret = C.jsc_class_get_name(_arg0)
	runtime.KeepAlive(jscClass)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Parent: get the parent class of jsc_class.
//
// The function returns the following values:
//
//   - class: parent class of jsc_class.
//
func (jscClass *Class) Parent() *Class {
	var _arg0 *C.JSCClass // out
	var _cret *C.JSCClass // in

	_arg0 = (*C.JSCClass)(unsafe.Pointer(coreglib.InternObject(jscClass).Native()))

	_cret = C.jsc_class_get_parent(_arg0)
	runtime.KeepAlive(jscClass)

	var _class *Class // out

	_class = wrapClass(coreglib.Take(unsafe.Pointer(_cret)))

	return _class
}

// ClassClass: instance of this type is always passed by reference.
type ClassClass struct {
	*classClass
}

// classClass is the struct that's finalized.
type classClass struct {
	native *C.JSCClassClass
}

// ClassVTable: virtual table for a JSCClass. This can be optionally used when
// registering a CClass in a CContext to provide a custom implementation for the
// class. All virtual functions are optional and can be set to NULL to fallback
// to the default implementation.
//
// An instance of this type is always passed by reference.
type ClassVTable struct {
	*classVTable
}

// classVTable is the struct that's finalized.
type classVTable struct {
	native *C.JSCClassVTable
}
