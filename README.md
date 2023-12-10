# ZenLog

A minimalist logging library for Go. Inspired by [Charm Log](https://github.com/charmbracelet/log).

![Made with VHS](https://vhs.charm.sh/vhs-7e0m3qWldNrjb7ZyVUODat.gif)

Features:

-   Colored and human readable output.
-   Multiple log levels (Debug, Info, Success, Warn, Error, Fatal).
-   Customizable time format.
-   Toggleable report caller.
-   Toggleable debug mode.

## Installation

Use `go get` to install the package.

```bash
go get github.com/zenitria/zenlog@latest
```

## Usage

Import package to your project.

```go
import "github.com/zenitria/zenlog"
```

### Send logs

```go
zenlog.Debug("Debug message") // Works only in debug mode
zenlog.Info("Info message")
zenlog.Success("Success message")
zenlog.Warning("Warning message")
zenlog.Error("Error message")
zenlog.Fatal("Fatal message") // Exits with code 1
```

### Format messages

```go
zenlog.Debugf("Debug message with %s", "formatting") // Works only in debug mode
zenlog.Infof("Info message with %s", "formatting")
zenlog.Successf("Success message with %s", "formatting")
zenlog.Warningf("Warning message with %s", "formatting")
zenlog.Errorf("Error message with %s", "formatting")
zenlog.Fatalf("Fatal message with %s", "formatting") // Exits with code 1
```

### Set custom time format

```go
zenlog.SetTimeFormat("YOUR TIME FORMAT") // Default: 2006-01-02 15:04:05
```

[Available time formats](https://gosamples.dev/date-time-format-cheatsheet/)

![Made with VHS](https://vhs.charm.sh/vhs-6XsBZzgAiLe0bHDLgvwABd.gif)

### Set debug mode

```go
zenlog.SetDebug(true) // Default: false
```

### Toggle report caller

```go
zenlog.SetReportCaller(true) // Default: false
```

![Made with VHS](https://vhs.charm.sh/vhs-4gE111CpFyQiUTi2bdfkDd.gif)
