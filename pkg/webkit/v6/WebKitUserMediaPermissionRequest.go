// Code generated by girgen. DO NOT EDIT.

package webkit

// #include <stdlib.h>
// #include <webkit/webkit.h>
import "C"

// UserMediaPermissionRequestClass: instance of this type is always passed by
// reference.
type UserMediaPermissionRequestClass struct {
	*userMediaPermissionRequestClass
}

// userMediaPermissionRequestClass is the struct that's finalized.
type userMediaPermissionRequestClass struct {
	native *C.WebKitUserMediaPermissionRequestClass
}
