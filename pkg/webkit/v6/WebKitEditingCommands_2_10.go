// Code generated by girgen. DO NOT EDIT.

package webkit

// #include <stdlib.h>
// #include <webkit/webkit.h>
import "C"

// EDITING_COMMAND_CREATE_LINK: create link command. Creates a link element
// that is inserted at the current cursor position. If there's a selection,
// the selected text will be used as the link text, otherwise the URL itself
// will be used. It receives the link URL as argument. This command should be
// executed with webkit_web_view_execute_editing_command_with_argument().
const EDITING_COMMAND_CREATE_LINK = "CreateLink"

// EDITING_COMMAND_INSERT_IMAGE: insert image command. Creates an image element
// that is inserted at the current cursor position. It receives an URI as
// argument, that is used as the image source. This command should be executed
// with webkit_web_view_execute_editing_command_with_argument().
const EDITING_COMMAND_INSERT_IMAGE = "InsertImage"
