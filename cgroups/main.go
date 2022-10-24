package main
import (
	"fmt"
	"github.com/containerd/cgroups"
	cgroupsv2 "github.com/containerd/cgroups/v2"
)

func main() {
	fmt.Printf("hello\n")
	cgroupV2 := false
	if cgroups.Mode() == cgroups.Unified {
		cgroupV2 = true
	}
	if cgroupV2 {
		fmt.Printf("v2\n")
		m, err := cgroupsv2.LoadSystemd("/", "")
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}
		ctrs, err := m.Controllers()
		fmt.Printf("%s\n", ctrs)
		metrics, err := m.Stat()
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}
		fmt.Printf("max: %d\n", metrics.Memory.UsageLimit)
		return
	}
	control, err := cgroups.Load(cgroups.V1, cgroups.StaticPath("/"))
	if err != nil {
		fmt.Printf("error: %s\n", err)
			
	}
	fmt.Printf("%x\n", control)
}
