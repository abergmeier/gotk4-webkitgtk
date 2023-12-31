// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit/webkit.h>
import "C"

// GType values.
var (
	GTypeContextMenuItem = coreglib.Type(C.webkit_context_menu_item_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeContextMenuItem, F: marshalContextMenuItem},
	})
}

// ContextMenuItemOverrides contains methods that are overridable.
type ContextMenuItemOverrides struct {
}

func defaultContextMenuItemOverrides(v *ContextMenuItem) ContextMenuItemOverrides {
	return ContextMenuItemOverrides{}
}

// ContextMenuItem: one item of a KitContextMenu.
//
// The KitContextMenu is composed of KitContextMenuItem<!-- -->s. These
// items can be created from a Action, from a KitContextMenuAction or from a
// KitContextMenuAction and a label. These KitContextMenuAction<!-- -->s denote
// stock actions for the items. You can also create separators and submenus.
type ContextMenuItem struct {
	_ [0]func() // equal guard
	coreglib.InitiallyUnowned
}

var ()

func init() {
	coreglib.RegisterClassInfo[*ContextMenuItem, *ContextMenuItemClass, ContextMenuItemOverrides](
		GTypeContextMenuItem,
		initContextMenuItemClass,
		wrapContextMenuItem,
		defaultContextMenuItemOverrides,
	)
}

func initContextMenuItemClass(gclass unsafe.Pointer, overrides ContextMenuItemOverrides, classInitFunc func(*ContextMenuItemClass)) {
	if classInitFunc != nil {
		class := (*ContextMenuItemClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapContextMenuItem(obj *coreglib.Object) *ContextMenuItem {
	return &ContextMenuItem{
		InitiallyUnowned: coreglib.InitiallyUnowned{
			Object: obj,
		},
	}
}

func marshalContextMenuItem(p uintptr) (interface{}, error) {
	return wrapContextMenuItem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewContextMenuItemFromGaction creates a new KitContextMenuItem for the given
// action and label.
//
// On activation target will be passed as parameter to the callback.
//
// The function takes the following parameters:
//
//   - action: #GAction.
//   - label: menu item label text.
//   - target (optional) to use as the action target.
//
// The function returns the following values:
//
//   - contextMenuItem: newly created KitContextMenuItem object.
//
func NewContextMenuItemFromGaction(action gio.Actioner, label string, target *glib.Variant) *ContextMenuItem {
	var _arg1 *C.GAction               // out
	var _arg2 *C.gchar                 // out
	var _arg3 *C.GVariant              // out
	var _cret *C.WebKitContextMenuItem // in

	_arg1 = (*C.GAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))
	if target != nil {
		_arg3 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(target)))
	}

	_cret = C.webkit_context_menu_item_new_from_gaction(_arg1, _arg2, _arg3)
	runtime.KeepAlive(action)
	runtime.KeepAlive(label)
	runtime.KeepAlive(target)

	var _contextMenuItem *ContextMenuItem // out

	_contextMenuItem = wrapContextMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _contextMenuItem
}

// NewContextMenuItemFromStockAction creates a new KitContextMenuItem for the
// given stock action.
//
// Stock actions are handled automatically by WebKit so that, for example, when
// a menu item created with a WEBKIT_CONTEXT_MENU_ACTION_STOP is activated the
// action associated will be handled by WebKit and the current load operation
// will be stopped. You can get the #GAction of a KitContextMenuItem created
// with a KitContextMenuAction with webkit_context_menu_item_get_gaction()
// and connect to the Action::activate signal to be notified when the item is
// activated, but you can't prevent the associated action from being performed.
//
// The function takes the following parameters:
//
//   - action stock action.
//
// The function returns the following values:
//
//   - contextMenuItem: newly created KitContextMenuItem object.
//
func NewContextMenuItemFromStockAction(action ContextMenuAction) *ContextMenuItem {
	var _arg1 C.WebKitContextMenuAction // out
	var _cret *C.WebKitContextMenuItem  // in

	_arg1 = C.WebKitContextMenuAction(action)

	_cret = C.webkit_context_menu_item_new_from_stock_action(_arg1)
	runtime.KeepAlive(action)

	var _contextMenuItem *ContextMenuItem // out

	_contextMenuItem = wrapContextMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _contextMenuItem
}

// NewContextMenuItemFromStockActionWithLabel creates a new KitContextMenuItem
// for the given stock action using the given label.
//
// Stock actions have a predefined label, this method can be used to create a
// KitContextMenuItem for a KitContextMenuAction but using a custom label.
//
// The function takes the following parameters:
//
//   - action stock action.
//   - label: custom label text to use instead of the predefined one.
//
// The function returns the following values:
//
//   - contextMenuItem: newly created KitContextMenuItem object.
//
func NewContextMenuItemFromStockActionWithLabel(action ContextMenuAction, label string) *ContextMenuItem {
	var _arg1 C.WebKitContextMenuAction // out
	var _arg2 *C.gchar                  // out
	var _cret *C.WebKitContextMenuItem  // in

	_arg1 = C.WebKitContextMenuAction(action)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.webkit_context_menu_item_new_from_stock_action_with_label(_arg1, _arg2)
	runtime.KeepAlive(action)
	runtime.KeepAlive(label)

	var _contextMenuItem *ContextMenuItem // out

	_contextMenuItem = wrapContextMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _contextMenuItem
}

// NewContextMenuItemSeparator creates a new KitContextMenuItem representing a
// separator.
//
// The function returns the following values:
//
//   - contextMenuItem: newly created KitContextMenuItem object.
//
func NewContextMenuItemSeparator() *ContextMenuItem {
	var _cret *C.WebKitContextMenuItem // in

	_cret = C.webkit_context_menu_item_new_separator()

	var _contextMenuItem *ContextMenuItem // out

	_contextMenuItem = wrapContextMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _contextMenuItem
}

// NewContextMenuItemWithSubmenu creates a new KitContextMenuItem using the
// given label with a submenu.
//
// The function takes the following parameters:
//
//   - label: menu item label text.
//   - submenu to set.
//
// The function returns the following values:
//
//   - contextMenuItem: newly created KitContextMenuItem object.
//
func NewContextMenuItemWithSubmenu(label string, submenu *ContextMenu) *ContextMenuItem {
	var _arg1 *C.gchar                 // out
	var _arg2 *C.WebKitContextMenu     // out
	var _cret *C.WebKitContextMenuItem // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.WebKitContextMenu)(unsafe.Pointer(coreglib.InternObject(submenu).Native()))

	_cret = C.webkit_context_menu_item_new_with_submenu(_arg1, _arg2)
	runtime.KeepAlive(label)
	runtime.KeepAlive(submenu)

	var _contextMenuItem *ContextMenuItem // out

	_contextMenuItem = wrapContextMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _contextMenuItem
}

// Gaction gets the action associated to item as a #GAction.
//
// The function returns the following values:
//
//   - action associated to the KitContextMenuItem, or NULL if item is a
//     separator.
//
func (item *ContextMenuItem) Gaction() *gio.Action {
	var _arg0 *C.WebKitContextMenuItem // out
	var _cret *C.GAction               // in

	_arg0 = (*C.WebKitContextMenuItem)(unsafe.Pointer(coreglib.InternObject(item).Native()))

	_cret = C.webkit_context_menu_item_get_gaction(_arg0)
	runtime.KeepAlive(item)

	var _action *gio.Action // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_action = &gio.Action{
			Object: obj,
		}
	}

	return _action
}

// StockAction gets the KitContextMenuAction of item.
//
// If the KitContextMenuItem was not created for a stock action
// WEBKIT_CONTEXT_MENU_ACTION_CUSTOM will be returned. If the KitContextMenuItem
// is a separator WEBKIT_CONTEXT_MENU_ACTION_NO_ACTION will be returned.
//
// The function returns the following values:
//
//   - contextMenuAction of item.
//
func (item *ContextMenuItem) StockAction() ContextMenuAction {
	var _arg0 *C.WebKitContextMenuItem  // out
	var _cret C.WebKitContextMenuAction // in

	_arg0 = (*C.WebKitContextMenuItem)(unsafe.Pointer(coreglib.InternObject(item).Native()))

	_cret = C.webkit_context_menu_item_get_stock_action(_arg0)
	runtime.KeepAlive(item)

	var _contextMenuAction ContextMenuAction // out

	_contextMenuAction = ContextMenuAction(_cret)

	return _contextMenuAction
}

// Submenu gets the submenu of item.
//
// The function returns the following values:
//
//   - contextMenu representing the submenu of item or NULL if item doesn't have
//     a submenu.
//
func (item *ContextMenuItem) Submenu() *ContextMenu {
	var _arg0 *C.WebKitContextMenuItem // out
	var _cret *C.WebKitContextMenu     // in

	_arg0 = (*C.WebKitContextMenuItem)(unsafe.Pointer(coreglib.InternObject(item).Native()))

	_cret = C.webkit_context_menu_item_get_submenu(_arg0)
	runtime.KeepAlive(item)

	var _contextMenu *ContextMenu // out

	_contextMenu = wrapContextMenu(coreglib.Take(unsafe.Pointer(_cret)))

	return _contextMenu
}

// IsSeparator checks whether item is a separator.
//
// The function returns the following values:
//
//   - ok: TRUE is item is a separator or FALSE otherwise.
//
func (item *ContextMenuItem) IsSeparator() bool {
	var _arg0 *C.WebKitContextMenuItem // out
	var _cret C.gboolean               // in

	_arg0 = (*C.WebKitContextMenuItem)(unsafe.Pointer(coreglib.InternObject(item).Native()))

	_cret = C.webkit_context_menu_item_is_separator(_arg0)
	runtime.KeepAlive(item)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetSubmenu sets or replaces the item submenu.
//
// If submenu is NULL the current submenu of item is removed.
//
// The function takes the following parameters:
//
//   - submenu (optional): KitContextMenu.
//
func (item *ContextMenuItem) SetSubmenu(submenu *ContextMenu) {
	var _arg0 *C.WebKitContextMenuItem // out
	var _arg1 *C.WebKitContextMenu     // out

	_arg0 = (*C.WebKitContextMenuItem)(unsafe.Pointer(coreglib.InternObject(item).Native()))
	if submenu != nil {
		_arg1 = (*C.WebKitContextMenu)(unsafe.Pointer(coreglib.InternObject(submenu).Native()))
	}

	C.webkit_context_menu_item_set_submenu(_arg0, _arg1)
	runtime.KeepAlive(item)
	runtime.KeepAlive(submenu)
}

// ContextMenuItemClass: instance of this type is always passed by reference.
type ContextMenuItemClass struct {
	*contextMenuItemClass
}

// contextMenuItemClass is the struct that's finalized.
type contextMenuItemClass struct {
	native *C.WebKitContextMenuItemClass
}
