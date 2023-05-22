// Copyright 2023 Jigsaw Operations LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// go:build windows

package connectivity

import (
	"fmt"
	"golang.org/x/sys/windows"
	"syscall"
)

func errnoName(errno syscall.Errno) string {
	// List from https://cs.opensource.google/go/x/sys/+/master:windows/zerrors_windows.go,
	// restricted to WSAE errors.
	switch errno {
	case windows.WSAEINTR:
		return "EINTR"
	case windows.WSAEBADF:
		return "EBADF"
	case windows.WSAEACCES:
		return "EACCES"
	case windows.WSAEFAULT:
		return "EFAULT"
	case windows.WSAEINVAL:
		return "EINVAL"
	case windows.WSAEMFILE:
		return "EMFILE"
	case windows.WSAEWOULDBLOCK:
		return "EWOULDBLOCK"
	case windows.WSAEINPROGRESS:
		return "EINPROGRESS"
	case windows.WSAEALREADY:
		return "EALREADY"
	case windows.WSAENOTSOCK:
		return "ENOTSOCK"
	case windows.WSAEDESTADDRREQ:
		return "EDESTADDRREQ"
	case windows.WSAEMSGSIZE:
		return "EMSGSIZE"
	case windows.WSAEPROTOTYPE:
		return "EPROTOTYPE"
	case windows.WSAENOPROTOOPT:
		return "ENOPROTOOPT"
	case windows.WSAEPROTONOSUPPORT:
		return "EPROTONOSUPPORT"
	case windows.WSAESOCKTNOSUPPORT:
		return "ESOCKTNOSUPPORT"
	case windows.WSAEOPNOTSUPP:
		return "EOPNOTSUPP"
	case windows.WSAEPFNOSUPPORT:
		return "EPFNOSUPPORT"
	case windows.WSAEAFNOSUPPORT:
		return "EAFNOSUPPORT"
	case windows.WSAEADDRINUSE:
		return "EADDRINUSE"
	case windows.WSAEADDRNOTAVAIL:
		return "EADDRNOTAVAIL"
	case windows.WSAENETDOWN:
		return "ENETDOWN"
	case windows.WSAENETUNREACH:
		return "ENETUNREACH"
	case windows.WSAENETRESET:
		return "ENETRESET"
	case windows.WSAECONNABORTED:
		return "ECONNABORTED"
	case windows.WSAECONNRESET:
		return "ECONNRESET"
	case windows.WSAENOBUFS:
		return "ENOBUFS"
	case windows.WSAEISCONN:
		return "EISCONN"
	case windows.WSAENOTCONN:
		return "ENOTCONN"
	case windows.WSAESHUTDOWN:
		return "ESHUTDOWN"
	case windows.WSAETOOMANYREFS:
		return "ETOOMANYREFS"
	case windows.WSAETIMEDOUT:
		return "ETIMEDOUT"
	case windows.WSAECONNREFUSED:
		return "ECONNREFUSED"
	case windows.WSAELOOP:
		return "ELOOP"
	case windows.WSAENAMETOOLONG:
		return "ENAMETOOLONG"
	case windows.WSAEHOSTDOWN:
		return "EHOSTDOWN"
	case windows.WSAEHOSTUNREACH:
		return "EHOSTUNREACH"
	case windows.WSAENOTEMPTY:
		return "ENOTEMPTY"
	case windows.WSAEPROCLIM:
		return "EPROCLIM"
	case windows.WSAEUSERS:
		return "EUSERS"
	case windows.WSAEDQUOT:
		return "EDQUOT"
	case windows.WSAESTALE:
		return "ESTALE"
	case windows.WSAEREMOTE:
		return "EREMOTE"
	case windows.WSAEDISCON:
		return "EDISCON"
	case windows.WSAENOMORE:
		return "ENOMORE"
	case windows.WSAECANCELLED:
		return "ECANCELLED"
	case windows.WSAEINVALIDPROCTABLE:
		return "EINVALIDPROCTABLE"
	case windows.WSAEINVALIDPROVIDER:
		return "EINVALIDPROVIDER"
	case windows.WSAEPROVIDERFAILEDINIT:
		return "EPROVIDERFAILEDINIT"
	case windows.WSAEREFUSED:
		return "EREFUSED"
	default:
		return fmt.Sprintf("Error %d (0x%x)", int(errno), int(errno))
	}
}