// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"context"
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
// void _gotk4_soup2_SocketCallback(SoupSocket*, guint, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_socket_io_status_get_type()), F: marshalSocketIOStatus},
		{T: externglib.Type(C.soup_socket_get_type()), F: marshalSocketter},
	})
}

// SOCKET_ASYNC_CONTEXT alias for the Socket:async-context property. (The
// socket's Context.).
const SOCKET_ASYNC_CONTEXT = "async-context"

// SOCKET_FLAG_NONBLOCKING alias for the Socket:non-blocking property. (Whether
// or not the socket uses non-blocking I/O.).
const SOCKET_FLAG_NONBLOCKING = "non-blocking"

// SOCKET_IS_SERVER alias for the Socket:is-server property, qv.
const SOCKET_IS_SERVER = "is-server"

// SOCKET_LOCAL_ADDRESS alias for the Socket:local-address property. (Address of
// local end of socket.).
const SOCKET_LOCAL_ADDRESS = "local-address"

// SOCKET_REMOTE_ADDRESS alias for the Socket:remote-address property. (Address
// of remote end of socket.).
const SOCKET_REMOTE_ADDRESS = "remote-address"

// SOCKET_SSL_CREDENTIALS alias for the Socket:ssl-creds property. (SSL
// credential information.).
const SOCKET_SSL_CREDENTIALS = "ssl-creds"

// SOCKET_SSL_FALLBACK alias for the Socket:ssl-fallback property.
const SOCKET_SSL_FALLBACK = "ssl-fallback"

// SOCKET_SSL_STRICT alias for the Socket:ssl-strict property.
const SOCKET_SSL_STRICT = "ssl-strict"

// SOCKET_TIMEOUT alias for the Socket:timeout property. (The timeout in seconds
// for blocking socket I/O operations.).
const SOCKET_TIMEOUT = "timeout"

// SOCKET_TLS_CERTIFICATE alias for the Socket:tls-certificate property. Note
// that this property's value is only useful if the socket is for a TLS
// connection, and only reliable after some data has been transferred to or from
// it.
const SOCKET_TLS_CERTIFICATE = "tls-certificate"

// SOCKET_TLS_ERRORS alias for the Socket:tls-errors property. Note that this
// property's value is only useful if the socket is for a TLS connection, and
// only reliable after some data has been transferred to or from it.
const SOCKET_TLS_ERRORS = "tls-errors"

// SOCKET_TRUSTED_CERTIFICATE alias for the Socket:trusted-certificate property.
const SOCKET_TRUSTED_CERTIFICATE = "trusted-certificate"

// SOCKET_USE_THREAD_CONTEXT alias for the Socket:use-thread-context property.
// (Use g_main_context_get_thread_default()).
const SOCKET_USE_THREAD_CONTEXT = "use-thread-context"

// SocketIOStatus: return value from the Socket IO methods.
type SocketIOStatus C.gint

const (
	// SocketOK: success.
	SocketOK SocketIOStatus = iota
	// SocketWouldBlock: cannot read/write any more at this time.
	SocketWouldBlock
	// SocketEOF: end of file.
	SocketEOF
	// SocketError: other error.
	SocketError
)

func marshalSocketIOStatus(p uintptr) (interface{}, error) {
	return SocketIOStatus(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SocketIOStatus.
func (s SocketIOStatus) String() string {
	switch s {
	case SocketOK:
		return "OK"
	case SocketWouldBlock:
		return "WouldBlock"
	case SocketEOF:
		return "EOF"
	case SocketError:
		return "Error"
	default:
		return fmt.Sprintf("SocketIOStatus(%d)", s)
	}
}

// SocketCallback: callback function passed to soup_socket_connect_async().
type SocketCallback func(sock *Socket, status uint)

//export _gotk4_soup2_SocketCallback
func _gotk4_soup2_SocketCallback(arg0 *C.SoupSocket, arg1 C.guint, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var sock *Socket // out
	var status uint  // out

	sock = wrapSocket(externglib.Take(unsafe.Pointer(arg0)))
	status = uint(arg1)

	fn := v.(SocketCallback)
	fn(sock, status)
}

// SocketOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SocketOverrider interface {
	Disconnected()
	NewConnection(newSock *Socket)
	Readable()
	Writable()
}

type Socket struct {
	*externglib.Object

	gio.Initable
}

var (
	_ externglib.Objector = (*Socket)(nil)
)

func wrapSocket(obj *externglib.Object) *Socket {
	return &Socket{
		Object: obj,
		Initable: gio.Initable{
			Object: obj,
		},
	}
}

func marshalSocketter(p uintptr) (interface{}, error) {
	return wrapSocket(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectAsync begins asynchronously connecting to sock's remote address. The
// socket will call callback when it succeeds or fails (but not before returning
// from this function).
//
// If cancellable is non-NULL, it can be used to cancel the connection. callback
// will still be invoked in this case, with a status of SOUP_STATUS_CANCELLED.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - callback to call after connecting.
//
func (sock *Socket) ConnectAsync(ctx context.Context, callback SocketCallback) {
	var _arg0 *C.SoupSocket        // out
	var _arg1 *C.GCancellable      // out
	var _arg2 C.SoupSocketCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(sock.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (*[0]byte)(C._gotk4_soup2_SocketCallback)
	_arg3 = C.gpointer(gbox.AssignOnce(callback))

	C.soup_socket_connect_async(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(sock)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// ConnectSync: attempt to synchronously connect sock to its remote address.
//
// If cancellable is non-NULL, it can be used to cancel the connection, in which
// case soup_socket_connect_sync() will return SOUP_STATUS_CANCELLED.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//
func (sock *Socket) ConnectSync(ctx context.Context) uint {
	var _arg0 *C.SoupSocket   // out
	var _arg1 *C.GCancellable // out
	var _cret C.guint         // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(sock.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C.soup_socket_connect_sync(_arg0, _arg1)
	runtime.KeepAlive(sock)
	runtime.KeepAlive(ctx)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Disconnect disconnects sock. Any further read or write attempts on it will
// fail.
func (sock *Socket) Disconnect() {
	var _arg0 *C.SoupSocket // out

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(sock.Native()))

	C.soup_socket_disconnect(_arg0)
	runtime.KeepAlive(sock)
}

// Fd gets sock's underlying file descriptor.
//
// Note that fiddling with the file descriptor may break the Socket.
func (sock *Socket) Fd() int {
	var _arg0 *C.SoupSocket // out
	var _cret C.int         // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(sock.Native()))

	_cret = C.soup_socket_get_fd(_arg0)
	runtime.KeepAlive(sock)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// LocalAddress returns the Address corresponding to the local end of sock.
//
// Calling this method on an unconnected socket is considered to be an error,
// and produces undefined results.
func (sock *Socket) LocalAddress() *Address {
	var _arg0 *C.SoupSocket  // out
	var _cret *C.SoupAddress // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(sock.Native()))

	_cret = C.soup_socket_get_local_address(_arg0)
	runtime.KeepAlive(sock)

	var _address *Address // out

	_address = wrapAddress(externglib.Take(unsafe.Pointer(_cret)))

	return _address
}

// RemoteAddress returns the Address corresponding to the remote end of sock.
//
// Calling this method on an unconnected socket is considered to be an error,
// and produces undefined results.
func (sock *Socket) RemoteAddress() *Address {
	var _arg0 *C.SoupSocket  // out
	var _cret *C.SoupAddress // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(sock.Native()))

	_cret = C.soup_socket_get_remote_address(_arg0)
	runtime.KeepAlive(sock)

	var _address *Address // out

	_address = wrapAddress(externglib.Take(unsafe.Pointer(_cret)))

	return _address
}

// IsConnected tests if sock is connected to another host.
func (sock *Socket) IsConnected() bool {
	var _arg0 *C.SoupSocket // out
	var _cret C.gboolean    // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(sock.Native()))

	_cret = C.soup_socket_is_connected(_arg0)
	runtime.KeepAlive(sock)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSSL tests if sock is doing (or has attempted to do) SSL.
func (sock *Socket) IsSSL() bool {
	var _arg0 *C.SoupSocket // out
	var _cret C.gboolean    // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(sock.Native()))

	_cret = C.soup_socket_is_ssl(_arg0)
	runtime.KeepAlive(sock)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Listen makes sock start listening on its local address. When connections come
// in, sock will emit Socket::new_connection.
func (sock *Socket) Listen() bool {
	var _arg0 *C.SoupSocket // out
	var _cret C.gboolean    // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(sock.Native()))

	_cret = C.soup_socket_listen(_arg0)
	runtime.KeepAlive(sock)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Read attempts to read up to len bytes from sock into buffer. If some data is
// successfully read, soup_socket_read() will return SOUP_SOCKET_OK, and *nread
// will contain the number of bytes actually read (which may be less than len).
//
// If sock is non-blocking, and no data is available, the return value will be
// SOUP_SOCKET_WOULD_BLOCK. In this case, the caller can connect to the
// Socket::readable signal to know when there is more data to read. (NB: You
// MUST read all available data off the socket first. Socket::readable is only
// emitted after soup_socket_read() returns SOUP_SOCKET_WOULD_BLOCK, and it is
// only emitted once. See the documentation for Socket:non-blocking.).
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - buffer to read into.
//
func (sock *Socket) Read(ctx context.Context, buffer []byte) (uint, SocketIOStatus, error) {
	var _arg0 *C.SoupSocket   // out
	var _arg4 *C.GCancellable // out
	var _arg1 C.gpointer      // out
	var _arg2 C.gsize
	var _arg3 C.gsize              // in
	var _cret C.SoupSocketIOStatus // in
	var _cerr *C.GError            // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(sock.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		_arg1 = (C.gpointer)(unsafe.Pointer(&buffer[0]))
	}

	_cret = C.soup_socket_read(_arg0, _arg1, _arg2, &_arg3, _arg4, &_cerr)
	runtime.KeepAlive(sock)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(buffer)

	var _nread uint                    // out
	var _socketIOStatus SocketIOStatus // out
	var _goerr error                   // out

	_nread = uint(_arg3)
	_socketIOStatus = SocketIOStatus(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _nread, _socketIOStatus, _goerr
}

// StartProxySsl starts using SSL on socket, expecting to find a host named
// ssl_host.
//
// The function takes the following parameters:
//
//    - ctx: #GCancellable.
//    - sslHost: hostname of the SSL server.
//
func (sock *Socket) StartProxySsl(ctx context.Context, sslHost string) bool {
	var _arg0 *C.SoupSocket   // out
	var _arg2 *C.GCancellable // out
	var _arg1 *C.char         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(sock.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(sslHost)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.soup_socket_start_proxy_ssl(_arg0, _arg1, _arg2)
	runtime.KeepAlive(sock)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(sslHost)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StartSSL starts using SSL on socket.
//
// The function takes the following parameters:
//
//    - ctx: #GCancellable.
//
func (sock *Socket) StartSSL(ctx context.Context) bool {
	var _arg0 *C.SoupSocket   // out
	var _arg1 *C.GCancellable // out
	var _cret C.gboolean      // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(sock.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C.soup_socket_start_ssl(_arg0, _arg1)
	runtime.KeepAlive(sock)
	runtime.KeepAlive(ctx)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Write attempts to write len bytes from buffer to sock. If some data is
// successfully written, the return status will be SOUP_SOCKET_OK, and *nwrote
// will contain the number of bytes actually written (which may be less than
// len).
//
// If sock is non-blocking, and no data could be written right away, the return
// value will be SOUP_SOCKET_WOULD_BLOCK. In this case, the caller can connect
// to the Socket::writable signal to know when more data can be written. (NB:
// Socket::writable is only emitted after soup_socket_write() returns
// SOUP_SOCKET_WOULD_BLOCK, and it is only emitted once. See the documentation
// for Socket:non-blocking.).
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - buffer: data to write.
//
func (sock *Socket) Write(ctx context.Context, buffer []byte) (uint, SocketIOStatus, error) {
	var _arg0 *C.SoupSocket   // out
	var _arg4 *C.GCancellable // out
	var _arg1 C.gconstpointer // out
	var _arg2 C.gsize
	var _arg3 C.gsize              // in
	var _cret C.SoupSocketIOStatus // in
	var _cerr *C.GError            // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(sock.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		_arg1 = (C.gconstpointer)(unsafe.Pointer(&buffer[0]))
	}

	_cret = C.soup_socket_write(_arg0, _arg1, _arg2, &_arg3, _arg4, &_cerr)
	runtime.KeepAlive(sock)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(buffer)

	var _nwrote uint                   // out
	var _socketIOStatus SocketIOStatus // out
	var _goerr error                   // out

	_nwrote = uint(_arg3)
	_socketIOStatus = SocketIOStatus(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _nwrote, _socketIOStatus, _goerr
}

// ConnectDisconnected: emitted when the socket is disconnected, for whatever
// reason.
func (sock *Socket) ConnectDisconnected(f func()) externglib.SignalHandle {
	return sock.Connect("disconnected", f)
}

// ConnectEvent: emitted when a network-related event occurs. See Client::event
// for more details.
func (sock *Socket) ConnectEvent(f func(event gio.SocketClientEvent, connection gio.IOStreamer)) externglib.SignalHandle {
	return sock.Connect("event", f)
}

// ConnectNewConnection: emitted when a listening socket (set up with
// soup_socket_listen()) receives a new connection.
//
// You must ref the new if you want to keep it; otherwise it will be destroyed
// after the signal is emitted.
func (sock *Socket) ConnectNewConnection(f func(new Socket)) externglib.SignalHandle {
	return sock.Connect("new-connection", f)
}

// ConnectReadable: emitted when an async socket is readable. See
// soup_socket_read(), soup_socket_read_until() and Socket:non-blocking.
func (sock *Socket) ConnectReadable(f func()) externglib.SignalHandle {
	return sock.Connect("readable", f)
}

// ConnectWritable: emitted when an async socket is writable. See
// soup_socket_write() and Socket:non-blocking.
func (sock *Socket) ConnectWritable(f func()) externglib.SignalHandle {
	return sock.Connect("writable", f)
}
