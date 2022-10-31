package dock

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"howett.net/plist"

	"github.com/bsthun/switchbox/program/entries/val"
	"github.com/bsthun/switchbox/program/types/extern"
	"github.com/bsthun/switchbox/program/wrappers/interactive"
)

func ApplyDock(preference *extern.DockPreference) {
	var dat []byte
	if bytes, err := plist.Marshal(preference, plist.XMLFormat); err != nil {
		interactive.Throw(val.DefaultTag, "UNABLE TO PARSE DOCK CONFIGURATION FILE", err, nil)
	} else {
		dat = bytes
	}

	// * Get home directory
	dirname, err := os.UserHomeDir()
	if err != nil {
		interactive.Throw(val.DefaultTag, "UNABLE TO GET USER HOME DIRECTORY", err, nil)
	}

	// * Write data to file
	if err := ioutil.WriteFile(path.Join(dirname, "/Library/Preferences/com.apple.dock.plist"), dat, 0644); err != nil {
		interactive.Throw(val.DefaultTag, "UNABLE TO WRITE DOCK CONFIGURATION FILE", err, nil)
	}

	// * Reload dock
	if _, err := exec.Command("killall", "Dock").Output(); err != nil {
		interactive.Throw(val.DefaultTag, "UNABLE TO RESTART DOCK", err, nil)
	}
}
