// Code generated by girgen. DO NOT EDIT.

package webkit

// #include <stdlib.h>
// #include <webkit/webkit.h>
import "C"

// DeviceInfoPermissionRequestClass: instance of this type is always passed by
// reference.
type DeviceInfoPermissionRequestClass struct {
	*deviceInfoPermissionRequestClass
}

// deviceInfoPermissionRequestClass is the struct that's finalized.
type deviceInfoPermissionRequestClass struct {
	native *C.WebKitDeviceInfoPermissionRequestClass
}
