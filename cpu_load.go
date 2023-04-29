package gocpuload

import (
	"context"
	"runtime"
	"time"
)

// RunCPULoad run CPU load in specify cores count and percentage
func RunCPULoad(ctx context.Context, coresCount int, timeSeconds int, percentage int) (context.Context, context.CancelFunc) {
	ctx, cancelfunc := context.WithTimeout(ctx, time.Duration(timeSeconds*1000*1000*1000))
	runtime.GOMAXPROCS(coresCount)

	// 1 unit = 100 ms may be the best
	unitHundresOfMicrosecond := 1000
	runMicrosecond := unitHundresOfMicrosecond * percentage
	sleepMicrosecond := unitHundresOfMicrosecond*100 - runMicrosecond

	for i := 0; i < coresCount; i++ {
		go func(ctx context.Context) {
			runtime.LockOSThread()
			// endless loop
			for {
				select {
				case <-ctx.Done():
					return
				default:
					begin := time.Now()
					for {
						// run 100%
						if time.Now().Sub(begin) > time.Duration(runMicrosecond)*time.Microsecond {
							break
						}
					}
					// sleep
					time.Sleep(time.Duration(sleepMicrosecond) * time.Microsecond)
				}
			}
		}(ctx)
	}

	return ctx, cancelfunc
}
