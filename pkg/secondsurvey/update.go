package secondsurvey

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	repo = `https://github.com/inexorabletash/travellermap`
)

func FetchOTU() error {
	tmp := os.TempDir()
	fmt.Println(tmp)
	cmd := exec.Command("git", "clone", repo)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to fetch data: %v", err)
	}

	return nil
}
