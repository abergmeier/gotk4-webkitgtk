// Code generated by girgen. DO NOT EDIT.

package soupgnome

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: libsoup-gnome-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <libsoup/soup-gnome.h>
import "C"

func GnomeFeatures226GetType() externglib.Type {
	var _cret C.GType // in

	_cret = C.soup_gnome_features_2_26_get_type()

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}
