package main

import (
	"github.com/ActiveState/sqlboiler/drivers"
	"github.com/ActiveState/sqlboiler/drivers/sqlboiler-psql/driver"
)

func main() {
	drivers.DriverMain(&driver.PostgresDriver{})
}
