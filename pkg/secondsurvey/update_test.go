package secondsurvey

import (
	"fmt"
	"testing"
)

func TestFetchOUT(t *testing.T) {
	err := FetchOTU()
	fmt.Println(err)
}
