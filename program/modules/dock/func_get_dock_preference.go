package dock

import (
	"os"
	"path"

	"howett.net/plist"

	"github.com/bsthun/switchbox/program/entries/val"
	"github.com/bsthun/switchbox/program/types/extern"
	"github.com/bsthun/switchbox/program/wrappers/interactive"
)

func GetDockPreference() extern.DockPreference {
	// * Get home directory
	dirname, err := os.UserHomeDir()
	if err != nil {
		interactive.Throw(val.DefaultTag, "UNABLE TO GET USER HOME DIRECTORY", err, nil)
	}

	// * Read dock configuration file
	dat, err := os.ReadFile(path.Join(dirname, "/Library/Preferences/com.apple.dock.plist"))
	if err != nil {
		interactive.Throw(val.DefaultTag, "UNABLE TO READ DOCK CONFIGURATION FILE", err, nil)
	}

	var dockItems extern.DockPreference
	if _, err := plist.Unmarshal(dat, &dockItems); err != nil {
		interactive.Throw(val.DefaultTag, "UNABLE TO PARSE DOCK CONFIGURATION FILE", err, nil)
	}

	return dockItems
}
