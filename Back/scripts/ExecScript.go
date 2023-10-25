package OSINT

import (
	"fmt"
	"net/http"
	"os/exec"
)

func ScriptExec(w http.ResponseWriter, r *http.Request) {
	pseudo := "clem"
	cmd := exec.Command("/bin/sh", "/home/clem/Osint/Bakerstreet/Back/server/test.sh", pseudo)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Script execution failed with error: %v\nOutput: %s\n", err, output)
	} else {
		fmt.Printf("Script output: %s\n", output)
	}

}
