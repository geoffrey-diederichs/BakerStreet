package OSINT

import (
	"fmt"
	"os/exec"
)

func ScriptExec(string) {
	pseudo := "clem"
	cmd := exec.Command("/bin/sh", "/home/clem/Osint/BakerStreet/Back/scripts/test.sh", pseudo)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Script execution failed with error: %v\nOutput: %s\n", err, output)
	} else {
		fmt.Printf("Script output: %s\n", output)
	}

}
