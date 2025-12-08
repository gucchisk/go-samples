package main

import (
	"fmt"

	"github.com/hashicorp/go-version"
)

func main() {
	v120, _ := version.NewVersion("1.2.0")
	v123, _ := version.NewVersion("1.2.3")
	v131, _ := version.NewVersion("1.3.1")
	v201, _ := version.NewVersion("2.0.1")

	constraints, _ := version.NewConstraint("1.2")
	fmt.Printf("*** Constraints: %s ***\n", constraints)
	printVersionInfo(v123)
	if constraints.Check(v123) {
		fmt.Printf("Version %s satisfies the constraints\n", v123.String())
	}
	printVersionInfo(v120)
	if constraints.Check(v120) {
		fmt.Printf("Version %s satisfies the constraints\n", v120.String())
	}
	printVersionInfo(v131)
	if constraints.Check(v131) {
		fmt.Printf("Version %s satisfies the constraints\n", v131.String())
	}
	printVersionInfo(v201)
	if constraints.Check(v201) {
		fmt.Printf("Version %s satisfies the constraints\n", v201.String())
	}

	constraints, _ = version.NewConstraint(">= 1.2.1, < 2.0.1")
	fmt.Printf("*** Constraints: %s ***\n", constraints)
	printVersionInfo(v123)
	if constraints.Check(v123) {
		fmt.Printf("Version %s satisfies the constraints\n", v123.String())
	}
	printVersionInfo(v120)
	if constraints.Check(v120) {
		fmt.Printf("Version %s satisfies the constraints\n", v120.String())
	}
	printVersionInfo(v131)
	if constraints.Check(v131) {
		fmt.Printf("Version %s satisfies the constraints\n", v131.String())
	}
	printVersionInfo(v201)
	if constraints.Check(v201) {
		fmt.Printf("Version %s satisfies the constraints\n", v201.String())
	}
}

func printVersionInfo(v *version.Version) {
	fmt.Printf("Version: %s\n", v.String())
	l := len(v.Segments())
	for i := 0; i < l; i++ {
		var label string
		switch i {
		case 0:
			label = "Major"
		case 1:
			label = "Minor"
		case 2:
			label = "Patch"
		default:
			label = fmt.Sprintf("Segment %d", i)
		}
		fmt.Printf(" %s: %d", label, v.Segments()[i])
	}
	fmt.Println()
}
