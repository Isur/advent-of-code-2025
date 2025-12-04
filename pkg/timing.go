package pkg

import (
	"fmt"
	"time"
)

func PrintTiming(elapsed time.Duration, part int) {
	seconds := elapsed.Seconds()

	if seconds >= 1 {
		fmt.Printf("Part %d took: %.3fs\n", part, seconds)
	} else if elapsed.Milliseconds() >= 1 {
		fmt.Printf("Part %d took: %.6fs (%.3fms)\n", part, seconds, float64(elapsed.Microseconds())/1000)
	} else if elapsed.Microseconds() >= 1 {
		fmt.Printf("Part %d took: %.9fs (%.3fµs)\n", part, seconds, float64(elapsed.Nanoseconds())/1000)
	} else {
		fmt.Printf("Part %d took: %.9fs (%dns)\n", part, seconds, elapsed.Nanoseconds())
	}
}
