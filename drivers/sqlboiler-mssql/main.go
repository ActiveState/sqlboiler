package main

import (
	"github.com/ActiveState/sqlboiler/drivers"
	"github.com/ActiveState/sqlboiler/drivers/sqlboiler-mssql/driver"
)

func main() {
	drivers.DriverMain(&driver.MSSQLDriver{})
}
