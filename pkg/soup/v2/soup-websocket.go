// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_websocket_close_code_get_type()), F: marshalWebsocketCloseCode},
		{T: externglib.Type(C.soup_websocket_connection_type_get_type()), F: marshalWebsocketConnectionType},
		{T: externglib.Type(C.soup_websocket_data_type_get_type()), F: marshalWebsocketDataType},
		{T: externglib.Type(C.soup_websocket_error_get_type()), F: marshalWebsocketError},
		{T: externglib.Type(C.soup_websocket_state_get_type()), F: marshalWebsocketState},
	})
}

// WebsocketCloseCode: pre-defined close codes that can be passed to
// soup_websocket_connection_close() or received from
// soup_websocket_connection_get_close_code(). (However, other codes are also
// allowed.).
type WebsocketCloseCode C.gint

const (
	// WebsocketCloseNormal: normal, non-error close.
	WebsocketCloseNormal WebsocketCloseCode = 1000
	// WebsocketCloseGoingAway: client/server is going away.
	WebsocketCloseGoingAway WebsocketCloseCode = 1001
	// WebsocketCloseProtocolError: protocol error occurred.
	WebsocketCloseProtocolError WebsocketCloseCode = 1002
	// WebsocketCloseUnsupportedData: endpoint received data of a type that it
	// does not support.
	WebsocketCloseUnsupportedData WebsocketCloseCode = 1003
	// WebsocketCloseNoStatus: reserved value indicating that no close code was
	// present; must not be sent.
	WebsocketCloseNoStatus WebsocketCloseCode = 1005
	// WebsocketCloseAbnormal: reserved value indicating that the connection was
	// closed abnormally; must not be sent.
	WebsocketCloseAbnormal WebsocketCloseCode = 1006
	// WebsocketCloseBadData: endpoint received data that was invalid (eg,
	// non-UTF-8 data in a text message).
	WebsocketCloseBadData WebsocketCloseCode = 1007
	// WebsocketClosePolicyViolation: generic error code indicating some sort of
	// policy violation.
	WebsocketClosePolicyViolation WebsocketCloseCode = 1008
	// WebsocketCloseTooBig: endpoint received a message that is too big to
	// process.
	WebsocketCloseTooBig WebsocketCloseCode = 1009
	// WebsocketCloseNoExtension: client is closing the connection because the
	// server failed to negotiate a required extension.
	WebsocketCloseNoExtension WebsocketCloseCode = 1010
	// WebsocketCloseServerError: server is closing the connection because it
	// was unable to fulfill the request.
	WebsocketCloseServerError WebsocketCloseCode = 1011
	// WebsocketCloseTLSHandshake: reserved value indicating that the TLS
	// handshake failed; must not be sent.
	WebsocketCloseTLSHandshake WebsocketCloseCode = 1015
)

func marshalWebsocketCloseCode(p uintptr) (interface{}, error) {
	return WebsocketCloseCode(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for WebsocketCloseCode.
func (w WebsocketCloseCode) String() string {
	switch w {
	case WebsocketCloseNormal:
		return "Normal"
	case WebsocketCloseGoingAway:
		return "GoingAway"
	case WebsocketCloseProtocolError:
		return "ProtocolError"
	case WebsocketCloseUnsupportedData:
		return "UnsupportedData"
	case WebsocketCloseNoStatus:
		return "NoStatus"
	case WebsocketCloseAbnormal:
		return "Abnormal"
	case WebsocketCloseBadData:
		return "BadData"
	case WebsocketClosePolicyViolation:
		return "PolicyViolation"
	case WebsocketCloseTooBig:
		return "TooBig"
	case WebsocketCloseNoExtension:
		return "NoExtension"
	case WebsocketCloseServerError:
		return "ServerError"
	case WebsocketCloseTLSHandshake:
		return "TLSHandshake"
	default:
		return fmt.Sprintf("WebsocketCloseCode(%d)", w)
	}
}

// WebsocketConnectionType: type of a WebsocketConnection.
type WebsocketConnectionType C.gint

const (
	// WebsocketConnectionUnknown: unknown/invalid connection.
	WebsocketConnectionUnknown WebsocketConnectionType = iota
	// WebsocketConnectionClient: client-side connection.
	WebsocketConnectionClient
	// WebsocketConnectionServer: server-side connection.
	WebsocketConnectionServer
)

func marshalWebsocketConnectionType(p uintptr) (interface{}, error) {
	return WebsocketConnectionType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for WebsocketConnectionType.
func (w WebsocketConnectionType) String() string {
	switch w {
	case WebsocketConnectionUnknown:
		return "Unknown"
	case WebsocketConnectionClient:
		return "Client"
	case WebsocketConnectionServer:
		return "Server"
	default:
		return fmt.Sprintf("WebsocketConnectionType(%d)", w)
	}
}

// WebsocketDataType: type of data contained in a WebsocketConnection::message
// signal.
type WebsocketDataType C.gint

const (
	// WebsocketDataText: UTF-8 text.
	WebsocketDataText WebsocketDataType = 1
	// WebsocketDataBinary: binary data.
	WebsocketDataBinary WebsocketDataType = 2
)

func marshalWebsocketDataType(p uintptr) (interface{}, error) {
	return WebsocketDataType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for WebsocketDataType.
func (w WebsocketDataType) String() string {
	switch w {
	case WebsocketDataText:
		return "Text"
	case WebsocketDataBinary:
		return "Binary"
	default:
		return fmt.Sprintf("WebsocketDataType(%d)", w)
	}
}

// WebsocketError: webSocket-related errors.
type WebsocketError C.gint

const (
	// WebsocketErrorFailed: generic error.
	WebsocketErrorFailed WebsocketError = iota
	// WebsocketErrorNotWebsocket: attempted to handshake with a server that
	// does not appear to understand WebSockets.
	WebsocketErrorNotWebsocket
	// WebsocketErrorBadHandshake: webSocket handshake failed because some
	// detail was invalid (eg, incorrect accept key).
	WebsocketErrorBadHandshake
	// WebsocketErrorBadOrigin: webSocket handshake failed because the "Origin"
	// header was not an allowed value.
	WebsocketErrorBadOrigin
)

func marshalWebsocketError(p uintptr) (interface{}, error) {
	return WebsocketError(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for WebsocketError.
func (w WebsocketError) String() string {
	switch w {
	case WebsocketErrorFailed:
		return "Failed"
	case WebsocketErrorNotWebsocket:
		return "NotWebsocket"
	case WebsocketErrorBadHandshake:
		return "BadHandshake"
	case WebsocketErrorBadOrigin:
		return "BadOrigin"
	default:
		return fmt.Sprintf("WebsocketError(%d)", w)
	}
}

func WebsocketErrorGetQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.soup_websocket_error_get_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

// WebsocketState: state of the WebSocket connection.
type WebsocketState C.gint

const (
	// WebsocketStateOpen: connection is ready to send messages.
	WebsocketStateOpen WebsocketState = 1
	// WebsocketStateClosing: connection is in the process of closing down;
	// messages may be received, but not sent.
	WebsocketStateClosing WebsocketState = 2
	// WebsocketStateClosed: connection is completely closed down.
	WebsocketStateClosed WebsocketState = 3
)

func marshalWebsocketState(p uintptr) (interface{}, error) {
	return WebsocketState(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for WebsocketState.
func (w WebsocketState) String() string {
	switch w {
	case WebsocketStateOpen:
		return "Open"
	case WebsocketStateClosing:
		return "Closing"
	case WebsocketStateClosed:
		return "Closed"
	default:
		return fmt.Sprintf("WebsocketState(%d)", w)
	}
}

// WebsocketClientPrepareHandshake adds the necessary headers to msg to request
// a WebSocket handshake. The message body and non-WebSocket-related headers are
// not modified.
//
// Use soup_websocket_client_prepare_handshake_with_extensions() if you want to
// include "Sec-WebSocket-Extensions" header in the request.
//
// This is a low-level function; if you use
// soup_session_websocket_connect_async() to create a WebSocket connection, it
// will call this for you.
//
// The function takes the following parameters:
//
//    - msg: Message.
//    - origin: "Origin" header to set.
//    - protocols: list of protocols to offer.
//
func WebsocketClientPrepareHandshake(msg *Message, origin string, protocols []string) {
	var _arg1 *C.SoupMessage // out
	var _arg2 *C.char        // out
	var _arg3 **C.char       // out

	_arg1 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))
	if origin != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(origin)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	{
		_arg3 = (**C.char)(C.malloc(C.size_t(uint((len(protocols) + 1)) * uint(unsafe.Sizeof(uint(0))))))
		defer C.free(unsafe.Pointer(_arg3))
		{
			out := unsafe.Slice(_arg3, len(protocols)+1)
			var zero *C.char
			out[len(protocols)] = zero
			for i := range protocols {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(protocols[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.soup_websocket_client_prepare_handshake(_arg1, _arg2, _arg3)
	runtime.KeepAlive(msg)
	runtime.KeepAlive(origin)
	runtime.KeepAlive(protocols)
}

// WebsocketClientVerifyHandshake looks at the response status code and headers
// in msg and determines if they contain a valid WebSocket handshake response
// (given the handshake request in msg's request headers).
//
// If the response contains the "Sec-WebSocket-Extensions" header, the handshake
// will be considered invalid. You need to use
// soup_websocket_client_verify_handshake_with_extensions() to handle responses
// with extensions.
//
// This is a low-level function; if you use
// soup_session_websocket_connect_async() to create a WebSocket connection, it
// will call this for you.
//
// The function takes the following parameters:
//
//    - msg containing both client and server sides of a WebSocket handshake.
//
func WebsocketClientVerifyHandshake(msg *Message) error {
	var _arg1 *C.SoupMessage // out
	var _cerr *C.GError      // in

	_arg1 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))

	C.soup_websocket_client_verify_handshake(_arg1, &_cerr)
	runtime.KeepAlive(msg)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// WebsocketServerCheckHandshake examines the method and request headers in msg
// and determines whether msg contains a valid handshake request.
//
// If origin is non-NULL, then only requests containing a matching "Origin"
// header will be accepted. If protocols is non-NULL, then only requests
// containing a compatible "Sec-WebSocket-Protocols" header will be accepted.
//
// Requests containing "Sec-WebSocket-Extensions" header will be accepted even
// if the header is not valid. To check a request with extensions you need to
// use soup_websocket_server_check_handshake_with_extensions() and provide the
// list of supported extension types.
//
// Normally soup_websocket_server_process_handshake() will take care of this for
// you, and if you use soup_server_add_websocket_handler() to handle accepting
// WebSocket connections, it will call that for you. However, this function may
// be useful if you need to perform more complicated validation; eg, accepting
// multiple different Origins, or handling different protocols depending on the
// path.
//
// The function takes the following parameters:
//
//    - msg containing the client side of a WebSocket handshake.
//    - origin: expected Origin header.
//    - protocols: allowed WebSocket protocols.
//
func WebsocketServerCheckHandshake(msg *Message, origin string, protocols []string) error {
	var _arg1 *C.SoupMessage // out
	var _arg2 *C.char        // out
	var _arg3 **C.char       // out
	var _cerr *C.GError      // in

	_arg1 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))
	if origin != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(origin)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	{
		_arg3 = (**C.char)(C.malloc(C.size_t(uint((len(protocols) + 1)) * uint(unsafe.Sizeof(uint(0))))))
		defer C.free(unsafe.Pointer(_arg3))
		{
			out := unsafe.Slice(_arg3, len(protocols)+1)
			var zero *C.char
			out[len(protocols)] = zero
			for i := range protocols {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(protocols[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.soup_websocket_server_check_handshake(_arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(msg)
	runtime.KeepAlive(origin)
	runtime.KeepAlive(protocols)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// WebsocketServerProcessHandshake examines the method and request headers in
// msg and (assuming msg contains a valid handshake request), fills in the
// handshake response.
//
// If expected_origin is non-NULL, then only requests containing a matching
// "Origin" header will be accepted. If protocols is non-NULL, then only
// requests containing a compatible "Sec-WebSocket-Protocols" header will be
// accepted.
//
// Requests containing "Sec-WebSocket-Extensions" header will be accepted even
// if the header is not valid. To process a request with extensions you need to
// use soup_websocket_server_process_handshake_with_extensions() and provide the
// list of supported extension types.
//
// This is a low-level function; if you use soup_server_add_websocket_handler()
// to handle accepting WebSocket connections, it will call this for you.
//
// The function takes the following parameters:
//
//    - msg containing the client side of a WebSocket handshake.
//    - expectedOrigin: expected Origin header.
//    - protocols: allowed WebSocket protocols.
//
func WebsocketServerProcessHandshake(msg *Message, expectedOrigin string, protocols []string) bool {
	var _arg1 *C.SoupMessage // out
	var _arg2 *C.char        // out
	var _arg3 **C.char       // out
	var _cret C.gboolean     // in

	_arg1 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))
	if expectedOrigin != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(expectedOrigin)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	{
		_arg3 = (**C.char)(C.malloc(C.size_t(uint((len(protocols) + 1)) * uint(unsafe.Sizeof(uint(0))))))
		defer C.free(unsafe.Pointer(_arg3))
		{
			out := unsafe.Slice(_arg3, len(protocols)+1)
			var zero *C.char
			out[len(protocols)] = zero
			for i := range protocols {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(protocols[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.soup_websocket_server_process_handshake(_arg1, _arg2, _arg3)
	runtime.KeepAlive(msg)
	runtime.KeepAlive(expectedOrigin)
	runtime.KeepAlive(protocols)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
