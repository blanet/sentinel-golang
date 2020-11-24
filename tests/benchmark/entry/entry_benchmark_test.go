package entry

import (
	"fmt"
	"testing"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/stat"
	"github.com/alibaba/sentinel-golang/tests/benchmark"
)

func init() {
	benchmark.InitSentinel()
}

func newSlotChain() *base.SlotChain {
	slotChain := base.NewSlotChain()
	slotChain.AddStatPrepareSlot(stat.DefaultResourceNodePrepareSlot)
	slotChain.AddStatSlot(stat.DefaultSlot)
	return slotChain
}

func sentinelEntry(sc *base.SlotChain) {
	e, b := sentinel.Entry("entry_benchmark_test", sentinel.WithSlotChain(sc))
	if b != nil {
		fmt.Println("blocked")
	} else {
		e.Exit()
	}
}

func Benchmark_Entry_Concurrency_1(b *testing.B) {
	sc := newSlotChain()
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sentinelEntry(sc)
		}
	})
}

func Benchmark_Entry_Concurrency_4(b *testing.B) {
	sc := newSlotChain()
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(4)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sentinelEntry(sc)
		}
	})
}

func Benchmark_Entry_Concurrency_8(b *testing.B) {
	sc := newSlotChain()
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(8)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sentinelEntry(sc)
		}
	})
}

func Benchmark_Entry_Concurrency_16(b *testing.B) {
	sc := newSlotChain()
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(16)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sentinelEntry(sc)
		}
	})
}

func Benchmark_Entry_Concurrency_32(b *testing.B) {
	sc := newSlotChain()
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(32)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sentinelEntry(sc)
		}
	})
}

func Benchmark_Entry_Concurrency_48(b *testing.B) {
	sc := newSlotChain()
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(48)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sentinelEntry(sc)
		}
	})
}

func Benchmark_Entry_Concurrency_64(b *testing.B) {
	sc := newSlotChain()
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(64)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sentinelEntry(sc)
		}
	})
}