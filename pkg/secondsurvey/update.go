package secondsurvey

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	repo         = `https://github.com/inexorabletash/travellermap`
	localSource  = `c:\Users\pemaltynov\go\src\github.com\Galdoba\travellermap\res\t5ss\data\`
	localProcess = `\travellermap\res\t5ss\data\`
)

func FetchOTU() error {
	tmp := os.TempDir()
	fmt.Println("fetch to", tmp)
	cmd := exec.Command("git", "clone", repo)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to fetch data: %v", err)
	}
	

	return os.Rename(tmp, fetchRoot)
}
