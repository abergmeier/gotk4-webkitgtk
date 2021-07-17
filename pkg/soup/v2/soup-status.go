// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"fmt"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_status_get_type()), F: marshalStatus},
	})
}

// Status: these represent the known HTTP status code values, plus various
// network and internal errors.
//
// Note that no libsoup functions take or return this type directly; any
// function that works with status codes will accept unrecognized status codes
// as well.
//
// Prior to 2.44 this type was called <literal>SoupKnownStatusCode</literal>,
// but the individual values have always had the names they have now.
type Status int

const (
	// StatusNone: no status available. (Eg, the message has not been sent yet)
	StatusNone Status = 0
	// StatusCancelled: message was cancelled locally
	StatusCancelled Status = 1
	// StatusCantResolve: unable to resolve destination host name
	StatusCantResolve Status = 2
	// StatusCantResolveProxy: unable to resolve proxy host name
	StatusCantResolveProxy Status = 3
	// StatusCantConnect: unable to connect to remote host
	StatusCantConnect Status = 4
	// StatusCantConnectProxy: unable to connect to proxy
	StatusCantConnectProxy Status = 5
	// StatusSSLFailed: SSL/TLS negotiation failed
	StatusSSLFailed Status = 6
	// StatusIOError: network error occurred, or the other end closed the
	// connection unexpectedly
	StatusIOError Status = 7
	// StatusMalformed: malformed data (usually a programmer error)
	StatusMalformed Status = 8
	// StatusTryAgain: used internally
	StatusTryAgain Status = 9
	// StatusTooManyRedirects: there were too many redirections
	StatusTooManyRedirects Status = 10
	// StatusTLSFailed: used internally
	StatusTLSFailed Status = 11
	// StatusContinue: 100 Continue (HTTP)
	StatusContinue Status = 100
	// StatusSwitchingProtocols: 101 Switching Protocols (HTTP)
	StatusSwitchingProtocols Status = 101
	// StatusProcessing: 102 Processing (WebDAV)
	StatusProcessing Status = 102
	// StatusOk: 200 Success (HTTP). Also used by many lower-level soup routines
	// to indicate success.
	StatusOk Status = 200
	// StatusCreated: 201 Created (HTTP)
	StatusCreated Status = 201
	// StatusAccepted: 202 Accepted (HTTP)
	StatusAccepted Status = 202
	// StatusNonAuthoritative: 203 Non-Authoritative Information (HTTP)
	StatusNonAuthoritative Status = 203
	// StatusNoContent: 204 No Content (HTTP)
	StatusNoContent Status = 204
	// StatusResetContent: 205 Reset Content (HTTP)
	StatusResetContent Status = 205
	// StatusPartialContent: 206 Partial Content (HTTP)
	StatusPartialContent Status = 206
	// StatusMultiStatus: 207 Multi-Status (WebDAV)
	StatusMultiStatus Status = 207
	// StatusMultipleChoices: 300 Multiple Choices (HTTP)
	StatusMultipleChoices Status = 300
	// StatusMovedPermanently: 301 Moved Permanently (HTTP)
	StatusMovedPermanently Status = 301
	// StatusFound: 302 Found (HTTP)
	StatusFound Status = 302
	// StatusMovedTemporarily: 302 Moved Temporarily (old name, RFC 2068)
	StatusMovedTemporarily Status = 302
	// StatusSeeOther: 303 See Other (HTTP)
	StatusSeeOther Status = 303
	// StatusNotModified: 304 Not Modified (HTTP)
	StatusNotModified Status = 304
	// StatusUseProxy: 305 Use Proxy (HTTP)
	StatusUseProxy Status = 305
	// StatusNotAppearingInThisProtocol: 306 [Unused] (HTTP)
	StatusNotAppearingInThisProtocol Status = 306
	// StatusTemporaryRedirect: 307 Temporary Redirect (HTTP)
	StatusTemporaryRedirect Status = 307
	StatusPermanentRedirect Status = 308
	// StatusBadRequest: 400 Bad Request (HTTP)
	StatusBadRequest Status = 400
	// StatusUnauthorized: 401 Unauthorized (HTTP)
	StatusUnauthorized Status = 401
	// StatusPaymentRequired: 402 Payment Required (HTTP)
	StatusPaymentRequired Status = 402
	// StatusForbidden: 403 Forbidden (HTTP)
	StatusForbidden Status = 403
	// StatusNotFound: 404 Not Found (HTTP)
	StatusNotFound Status = 404
	// StatusMethodNotAllowed: 405 Method Not Allowed (HTTP)
	StatusMethodNotAllowed Status = 405
	// StatusNotAcceptable: 406 Not Acceptable (HTTP)
	StatusNotAcceptable Status = 406
	// StatusProxyAuthenticationRequired: 407 Proxy Authentication Required
	// (HTTP)
	StatusProxyAuthenticationRequired Status = 407
	// StatusProxyUnauthorized: shorter alias for
	// SOUP_STATUS_PROXY_AUTHENTICATION_REQUIRED
	StatusProxyUnauthorized Status = 407
	// StatusRequestTimeout: 408 Request Timeout (HTTP)
	StatusRequestTimeout Status = 408
	// StatusConflict: 409 Conflict (HTTP)
	StatusConflict Status = 409
	// StatusGone: 410 Gone (HTTP)
	StatusGone Status = 410
	// StatusLengthRequired: 411 Length Required (HTTP)
	StatusLengthRequired Status = 411
	// StatusPreconditionFailed: 412 Precondition Failed (HTTP)
	StatusPreconditionFailed Status = 412
	// StatusRequestEntityTooLarge: 413 Request Entity Too Large (HTTP)
	StatusRequestEntityTooLarge Status = 413
	// StatusRequestURITooLong: 414 Request-URI Too Long (HTTP)
	StatusRequestURITooLong Status = 414
	// StatusUnsupportedMediaType: 415 Unsupported Media Type (HTTP)
	StatusUnsupportedMediaType Status = 415
	// StatusRequestedRangeNotSatisfiable: 416 Requested Range Not Satisfiable
	// (HTTP)
	StatusRequestedRangeNotSatisfiable Status = 416
	// StatusInvalidRange: shorter alias for
	// SOUP_STATUS_REQUESTED_RANGE_NOT_SATISFIABLE
	StatusInvalidRange Status = 416
	// StatusExpectationFailed: 417 Expectation Failed (HTTP)
	StatusExpectationFailed Status = 417
	// StatusUnprocessableEntity: 422 Unprocessable Entity (WebDAV)
	StatusUnprocessableEntity Status = 422
	// StatusLocked: 423 Locked (WebDAV)
	StatusLocked Status = 423
	// StatusFailedDependency: 424 Failed Dependency (WebDAV)
	StatusFailedDependency Status = 424
	// StatusInternalServerError: 500 Internal Server Error (HTTP)
	StatusInternalServerError Status = 500
	// StatusNotImplemented: 501 Not Implemented (HTTP)
	StatusNotImplemented Status = 501
	// StatusBadGateway: 502 Bad Gateway (HTTP)
	StatusBadGateway Status = 502
	// StatusServiceUnavailable: 503 Service Unavailable (HTTP)
	StatusServiceUnavailable Status = 503
	// StatusGatewayTimeout: 504 Gateway Timeout (HTTP)
	StatusGatewayTimeout Status = 504
	// StatusHttpVersionNotSupported: 505 HTTP Version Not Supported (HTTP)
	StatusHttpVersionNotSupported Status = 505
	// StatusInsufficientStorage: 507 Insufficient Storage (WebDAV)
	StatusInsufficientStorage Status = 507
	// StatusNotExtended: 510 Not Extended (RFC 2774)
	StatusNotExtended Status = 510
)

func marshalStatus(p uintptr) (interface{}, error) {
	return Status(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for Status.
func (s Status) String() string {
	switch s {
	case StatusNone:
		return "None"
	case StatusCancelled:
		return "Cancelled"
	case StatusCantResolve:
		return "CantResolve"
	case StatusCantResolveProxy:
		return "CantResolveProxy"
	case StatusCantConnect:
		return "CantConnect"
	case StatusCantConnectProxy:
		return "CantConnectProxy"
	case StatusSSLFailed:
		return "SSLFailed"
	case StatusIOError:
		return "IOError"
	case StatusMalformed:
		return "Malformed"
	case StatusTryAgain:
		return "TryAgain"
	case StatusTooManyRedirects:
		return "TooManyRedirects"
	case StatusTLSFailed:
		return "TLSFailed"
	case StatusContinue:
		return "Continue"
	case StatusSwitchingProtocols:
		return "SwitchingProtocols"
	case StatusProcessing:
		return "Processing"
	case StatusOk:
		return "Ok"
	case StatusCreated:
		return "Created"
	case StatusAccepted:
		return "Accepted"
	case StatusNonAuthoritative:
		return "NonAuthoritative"
	case StatusNoContent:
		return "NoContent"
	case StatusResetContent:
		return "ResetContent"
	case StatusPartialContent:
		return "PartialContent"
	case StatusMultiStatus:
		return "MultiStatus"
	case StatusMultipleChoices:
		return "MultipleChoices"
	case StatusMovedPermanently:
		return "MovedPermanently"
	case StatusFound:
		return "Found"
	case StatusSeeOther:
		return "SeeOther"
	case StatusNotModified:
		return "NotModified"
	case StatusUseProxy:
		return "UseProxy"
	case StatusNotAppearingInThisProtocol:
		return "NotAppearingInThisProtocol"
	case StatusTemporaryRedirect:
		return "TemporaryRedirect"
	case StatusPermanentRedirect:
		return "PermanentRedirect"
	case StatusBadRequest:
		return "BadRequest"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusPaymentRequired:
		return "PaymentRequired"
	case StatusForbidden:
		return "Forbidden"
	case StatusNotFound:
		return "NotFound"
	case StatusMethodNotAllowed:
		return "MethodNotAllowed"
	case StatusNotAcceptable:
		return "NotAcceptable"
	case StatusProxyAuthenticationRequired:
		return "ProxyAuthenticationRequired"
	case StatusRequestTimeout:
		return "RequestTimeout"
	case StatusConflict:
		return "Conflict"
	case StatusGone:
		return "Gone"
	case StatusLengthRequired:
		return "LengthRequired"
	case StatusPreconditionFailed:
		return "PreconditionFailed"
	case StatusRequestEntityTooLarge:
		return "RequestEntityTooLarge"
	case StatusRequestURITooLong:
		return "RequestURITooLong"
	case StatusUnsupportedMediaType:
		return "UnsupportedMediaType"
	case StatusRequestedRangeNotSatisfiable:
		return "RequestedRangeNotSatisfiable"
	case StatusExpectationFailed:
		return "ExpectationFailed"
	case StatusUnprocessableEntity:
		return "UnprocessableEntity"
	case StatusLocked:
		return "Locked"
	case StatusFailedDependency:
		return "FailedDependency"
	case StatusInternalServerError:
		return "InternalServerError"
	case StatusNotImplemented:
		return "NotImplemented"
	case StatusBadGateway:
		return "BadGateway"
	case StatusServiceUnavailable:
		return "ServiceUnavailable"
	case StatusGatewayTimeout:
		return "GatewayTimeout"
	case StatusHttpVersionNotSupported:
		return "HttpVersionNotSupported"
	case StatusInsufficientStorage:
		return "InsufficientStorage"
	case StatusNotExtended:
		return "NotExtended"
	default:
		return fmt.Sprintf("Status(%d)", s)
	}
}

// StatusGetPhrase looks up the stock HTTP description of status_code. This is
// used by soup_message_set_status() to get the correct text to go with a given
// status code.
//
// <emphasis>There is no reason for you to ever use this function.</emphasis> If
// you wanted the textual description for the Message:status_code of a given
// Message, you should just look at the message's Message:reason_phrase.
// However, you should only do that for use in debugging messages; HTTP reason
// phrases are not localized, and are not generally very descriptive anyway, and
// so they should never be presented to the user directly. Instead, you should
// create you own error messages based on the status code, and on what you were
// trying to do.
func StatusGetPhrase(statusCode uint) string {
	var _arg1 C.guint // out
	var _cret *C.char // in

	_arg1 = C.guint(statusCode)

	_cret = C.soup_status_get_phrase(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// StatusProxify turns SOUP_STATUS_CANT_RESOLVE into
// SOUP_STATUS_CANT_RESOLVE_PROXY and SOUP_STATUS_CANT_CONNECT into
// SOUP_STATUS_CANT_CONNECT_PROXY. Other status codes are passed through
// unchanged.
func StatusProxify(statusCode uint) uint {
	var _arg1 C.guint // out
	var _cret C.guint // in

	_arg1 = C.guint(statusCode)

	_cret = C.soup_status_proxify(_arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}
