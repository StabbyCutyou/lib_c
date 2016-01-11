package main

import "github.com/StabbyCutyou/pinner"

func main() {
	pinner.Register("github.com/StabbyCutyou/lib_e", "~> 0.5")
	pinner.Register("github.com/StabbyCutyou/lib_d", "~> 1.0")

	pinner.RunMain()
}
