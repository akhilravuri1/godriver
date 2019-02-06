// +build !windows

package api

//#include <sqlcli1.h>
import "C"

func SQLAllocHandle(handleType SQLSMALLINT, inputHandle SQLHANDLE, outputHandle *SQLHANDLE) (ret SQLRETURN) {
	r := C.SQLAllocHandle(C.SQLSMALLINT(handleType), C.SQLHANDLE(inputHandle), (*C.SQLHANDLE)(outputHandle))
	return SQLRETURN(r)
}
