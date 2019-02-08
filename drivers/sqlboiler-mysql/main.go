package main

import (
	"github.com/ActiveState/sqlboiler/drivers"
	"github.com/ActiveState/sqlboiler/drivers/sqlboiler-mysql/driver"
)

func main() {
	drivers.DriverMain(&driver.MySQLDriver{})
}
