package secondsurvey

import (
	"os"
	"path/filepath"
)

var sep string
var fetchRoot string
var dbPath string
var dbRoot string

/*
XDG_DATA_HOME
For user application's own data files
Default to $HOME/.local/share
XDG_CONFIG_HOME
For user's app configuration files
Default to $HOME/.config
XDG_STATE_HOME
For user-specific app session data, which should be stored for future reuse
Default to $HOME/.local/state
May include logs, recently used files, application-specific information (e.g. window layout, views, opened files, undo history, etc.), akin to session data that should be stored by app by request of system session manager, like X session manager
XDG_CACHE_HOME
For user-specific apps cache files
Default to $HOME/.cache
*/

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	fetchRoot = filepath.Join(home, "travellermap")
	dbRoot = filepath.Join(home, ".local", "share", "t5")
	dbPath = filepath.Join(home, ".cache", "t5", "otu", "M1105", "spacemap")
	// fmt.Println(fetchRoot)
	// fmt.Println(dbPath)
	// fmt.Println(os.TempDir())
	// for _, dir := range []string{fetchRoot, dbRoot, dbPath} {
	// 	os.MkdirAll(dir, 0755)
	// }

}
