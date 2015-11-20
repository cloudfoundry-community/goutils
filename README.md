# goutils
S&amp;W Go Utilities

## How To Use

Get:

```
go get github.com/starkandwayne/goutils/{{pkg}}
```

Import:

```
import "github.com/starkandwayne/goutils/{{pkg}}"
```

## Log

Log has the following levels defined:

* Debug
* Info
* Notice
* Warn
* Error
* Crit
* Alert
* Emerg

Usage is the same as `Sprintf`/`Printf` statements - simply append an `f` to the desired level. e.g.:

```
dbug_mesg := "This isn't a bug."
log.Debugf("I really need to know this in debug mode: %s", dbug_mesg)
```
