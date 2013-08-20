package main
import (
	"fmt"
	"testing"
	"flag"
	"math"
	"os"
	"log"
	"runtime/pprof"
	"github.com/bjartwolf/abs"
)

var numbers  = [10]int{-340,-4253656356263,454265436245435,0,-24,-2324542544354352,-4,42,-24,-math.MaxInt64}
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func BenchmarkSaneAbs(b *testing.B) {
	for i:= 0; i< b.N; i++ {
		for _, number := range numbers {
			abs.SaneAbs(number)
		}
	}
}

func BenchmarkInSaneAbs(b *testing.B) {
	for i:= 0; i< b.N; i++ {
		for _, number := range numbers {
			abs.InSaneAbs(number)
		}
	}
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	fmt.Println("Testing")
        fmt.Println(testing.Benchmark(BenchmarkSaneAbs))
        fmt.Println(testing.Benchmark(BenchmarkInSaneAbs))
}
