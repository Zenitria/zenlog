package main

import (
	"time"

	"github.com/zenitria/zenlog"
)

func main() {
	zenlog.SetDebug(true)
	zenlog.SetTimeFormat(time.RFC822)

	zenlog.Info("Starting app...")
	time.Sleep(100 * time.Millisecond)

	zenlog.Info("Coping files...")
	zenlog.Debug("cp -r /backup /production")
	time.Sleep(300 * time.Millisecond)

	zenlog.Success("Files copied successfully!")
	time.Sleep(1 * time.Second)

	zenlog.Info("Unziping files...")
	zenlog.Debug("unzip /production/backup.zip")
	time.Sleep(100 * time.Millisecond)

	zenlog.Error("Files were not unziped!")
	zenlog.Warn("This files are not important!")
	time.Sleep(500 * time.Millisecond)

	zenlog.Info("Finding base.sql")
	zenlog.Debug("find /production -name base.sql")
	time.Sleep(100 * time.Millisecond)

	zenlog.Fatal("base.sql not found!")
}
