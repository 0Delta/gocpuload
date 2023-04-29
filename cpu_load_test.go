package gocpuload_test

import (
	"fmt"
	"math"
	"runtime"
	"testing"
	"time"

	"github.com/0Delta/gocpuload"
	cpu "github.com/shirou/gopsutil/v3/cpu"
)

var (
	target_cpuload    = 35
	cpuload_threshold = 5.0
	target_time       = 3
	target_core       = imin(runtime.NumCPU()/2, 2)
	time_threshold    = 3
	failed_threshold  = 3
)

func imin(in ...int) int {
	ret := math.MaxInt
	for _, i := range in {
		if ret > i {
			ret = i
		}
	}
	return ret
}

func TestSimple(t *testing.T) {
	cp_prev, err := cpu.Percent(time.Second, true)
	if err != nil {
		t.Fatalf("get cpupercent error %s:", err)
	}
	cp_prev_avg := f64avg(cp_prev)
	t.Logf("initial load %s:", f64tstr(cp_prev))
	t.Logf("    avg %1.1f:", cp_prev_avg)

	go gocpuload.RunCPULoad(target_core, target_time, target_cpuload)

	failed_count := 0
	timer := 0
	for {
		timer += 1
		time.Sleep(time.Millisecond * 500)
		cp_load, err := cpu.Percent(time.Second, true)
		if err != nil {
			t.Fatalf("get cpupercent error %s:", err)
		}

		loadcount := 0
		for _, l := range cp_load {
			if failed_count >= failed_threshold {
				t.Fatal("Test failed.")
				return
			}
			if (-1*cpuload_threshold) < (l-cp_prev_avg) && (l-cp_prev_avg) < cpuload_threshold {
				continue
			}
			loadcount += 1
		}
		if loadcount != 2 {
			if f64avg(cp_load) < cpuload_threshold && timer >= 3 {
				t.Logf("load power %s:", f64tstr(cp_load))
				if loadcount == 0 {
					t.Logf("load successful shutdown, Test Clear")
					return
				}
			}
			t.Errorf("load power error %s:", f64tstr(cp_load))
			failed_count += 1
		}
		if timer > target_time+time_threshold {
			t.Errorf("load time error")
			failed_count += 1
		}
		t.Logf("load power %s:", f64tstr(cp_load))
	}
}

func f64tstr(f64s []float64) string {
	ret := "{"
	for _, f := range f64s {
		ret += fmt.Sprintf("%1.1f, ", f)
	}
	ret = ret[:len(ret)-2] + "}"
	return ret
}

func f64avg(in []float64) float64 {
	ret := 0.0
	for _, f := range in {
		ret += f
	}
	return ret / float64(len(in))
}
