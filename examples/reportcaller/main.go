package main

import (
	"time"

	"github.com/zenitria/zenlog"
)

func main() {
	zenlog.SetReportCaller(true)

	zenlog.Debug("cp -r /backup /production")
	time.Sleep(300 * time.Millisecond)

	zenlog.Debug("unzip /production/backup.zip")
	time.Sleep(100 * time.Millisecond)

	zenlog.Debug("find /production -name base.sql")
	time.Sleep(100 * time.Millisecond)
}
