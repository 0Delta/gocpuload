gocpuload
=========
[![Go Reference](https://pkg.go.dev/badge/golang.org/x/pkgsite.svg)][goRef]

Usage:
------

### as CLI

example 01: run 30% of all CPU cores for 10 seconds

```sh
gocpuload -p 30 -t 10
```

example 02: run 30% of all CPU cores forver

```sh
gocpuload -p 30
```

example 03: run 30% of 2 of CPU cores for 10 seconds

```sh
gocpuload -p 30 -c 2 -t 10
```

- `all CPU load` = (num of para `c` _ num of `p`) / (all cores count of CPU _ 100)
- may not specify cores run the load only, it just promise the `all CPU load`, and not promise each cores run the same load

#### CLI Parameters

```
--coresCount value, -c value   how many cores (optional, default: 8)

--timeSeconds value, -t value  how long (optional, default: 2147483647)

--percentage value, -p value   percentage of each specify cores (required)

--help, -h                     show help
```

### as Liblary

```go
import "github.com/0Delta/gocpuload"

func main(){
    // RunCPULoad(coresCount int, timeSeconds int, percentage int)
    //
    // This function currently does not return control immediately.
    // If you want something action under pressure, use goroutine.
    go gocpuload.RunCPULoad(runtime.NumCPU()/2, 3, 80)

}
```

See [GoDoc][goRef]

Requirements:
-------------
+ go

for test in Windows
  + cgo

Install:
--------
+ go1.6 or higher
```
go install github.com/0Delta/gocpuload@latest
```

+ go1.5 or lower
```
go get -u github.com/0Delta/gocpuload
```

How it runs:
--------

- Giving a range of time(e.g. 100ms)
- Want to run 30% of all CPU cores
  - 30ms: run (CPU 100%)
  - 70ms: sleep(CPU 0%)

license:
--------
MIT

Author:
-------
0Δ(0deltast@gmail.com)


[goRef]:https://pkg.go.dev/github.com/0Delta/gocpuload
