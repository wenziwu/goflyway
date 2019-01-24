package xl

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

// Refer to https://msdn.microsoft.com/en-us/library/windows/desktop/ms740668(v=vs.85).aspx

var WSAErrno = map[int]string{
	10004: "Interrupted function call",
	10009: "File handle is not valid",
	10013: "Permission denied",
	10014: "Bad address",
	10022: "Invalid argument",
	10024: "Too many open files",
	10035: "Resource temporarily unavailable",
	10036: "Operation now in progress",
	10037: "Operation already in progress",
	10038: "Socket operation on nonsocket",
	10039: "Destination address required",
	10040: "Message too long",
	10041: "Protocol wrong type for socket",
	10042: "Bad protocol option",
	10043: "Protocol not supported",
	10044: "Socket type not supported",
	10045: "Operation not supported",
	10046: "Protocol family not supported",
	10047: "Address family not supported by protocol family",
	10048: "Address already in use",
	10049: "Cannot assign requested address",
	10050: "Network is down",
	10051: "Network is unreachable",
	10052: "Network dropped connection on reset",
	10053: "Software caused connection abort",
	10054: "Connection reset by peer",
	10055: "No buffer space available",
	10056: "Socket is already connected",
	10057: "Socket is not connected",
	10058: "Cannot send after socket shutdown",
	10059: "Too many references",
	10060: "Connection timed out",
	10061: "Connection refused",
	10062: "Cannot translate name",
	10063: "Name too long",
	10064: "Host is down",
	10065: "No route to host",
	10066: "Directory not empty",
	10067: "Too many processes",
	10068: "User quota exceeded",
	10069: "Disk quota exceeded",
	10070: "Stale file handle reference",
	10071: "Item is remote",
	10091: "Network subsystem is unavailable",
	10092: "Winsock.dll version out of range",
	10093: "Successful WSAStartup not yet performed",
	10101: "Graceful shutdown in progress",
	10102: "No more results",
	10103: "Call has been canceled",
	10104: "Procedure call table is invalid",
	10105: "Service provider is invalid",
	10106: "Service provider failed to initialize",
	10107: "System call failure",
	10108: "Service not found",
	10109: "Class type not found",
	10110: "No more results",
	10111: "Call was canceled",
	10112: "Database query was refused",
	11001: "Host not found",
	11002: "Nonauthoritative host not found",
	11003: "This is a nonrecoverable error",
	11004: "Valid name, no data record of requested type",
	11005: "QoS receivers",
	11006: "QoS senders",
	11007: "No QoS senders",
	11008: "QoS no receivers",
	11009: "QoS request confirmed",
	11010: "QoS admission error",
	11011: "QoS policy failure",
	11012: "QoS bad style",
	11013: "QoS bad object",
	11014: "QoS traffic control error",
	11015: "QoS generic error",
	11016: "QoS service type error",
	11017: "QoS flowspec error",
	11018: "Invalid QoS provider buffer",
	11019: "Invalid QoS filter style",
	11020: "Invalid QoS filter type",
	11021: "Incorrect QoS filter count",
	11022: "Invalid QoS object length",
	11023: "Incorrect QoS flow count",
	11024: "Unrecognized QoS object",
	11025: "Invalid QoS policy object",
	11026: "Invalid QoS flow descriptor",
	11027: "Invalid QoS provider-specific flowspec",
	11028: "Invalid QoS provider-specific filterspec",
	11029: "Invalid QoS shape discard mode object",
	11030: "Invalid QoS shaping rate object",
	11031: "Reserved policy QoS element type",
}

// Widnows WSA error messages are way too long to print
// ex: An established connection was aborted by the software in your host machine.write tcp 127.0.0.1:8100->127.0.0.1:52466: wsasend: An established connection was aborted by the software in your host machine.
func tryShortenWSAError(err interface{}) (ret string) {
	defer func() {
		if recover() != nil {
			ret = fmt.Sprintf("%v", err)
		}
	}()

	if e, sysok := err.(*net.OpError).Err.(*os.SyscallError); sysok {
		errno := e.Err.(syscall.Errno)
		if msg, ok := WSAErrno[int(errno)]; ok {
			ret = msg
		} else {
			// messages on linux are short enough
			ret = fmt.Sprintf("C%d, %s", uintptr(errno), e.Error())
		}

		return
	}

	ret = err.(*net.OpError).Err.Error()
	return
}
