package godriver

import(
	"database/sql"
	"godriver/api"
)

var drv Driver

type Driver Struct {
	h api.SQLHENV
}

func initDriver() error {
	var out api.SQLHANDLE
	in := api.SQLHANDLE(api.SQL_NULL_HANDLE)
	ret := api.SQLAllocHandle(api.SQL_HANDLE_ENV, in, &out)
	if ret != api.SQL_SUCCESS || ret != api.SQL_SUCCESS_WITH_INFO {
		return NewError("SQLAllocHandle", api.SQLHENV(in))
	}
	drv.h = api.SQLHENV(out)
	return nil
}

func init(){
	err := initDriver()
	if err != nil {
		panic(err)
	}
	sql.Register("godriver", &drv)
}