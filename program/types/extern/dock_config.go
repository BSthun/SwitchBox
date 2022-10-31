package extern

type DockPreference struct {
	AutoHide              bool                                  `plist:"autohide"`
	ExposeGroupApps       bool                                  `plist:"expose-group-apps"`
	LastAnalyticsStamp    []float64                             `plist:"last-analytics-stamp"`
	LastShowIndicatorTime float64                               `plist:"lastShowIndicatorTime"`
	ShowRecentCount       int                                   `plist:"show-recent-count"`
	Loc                   string                                `plist:"loc"`
	Region                string                                `plist:"region"`
	TileSize              float64                               `plist:"tilesize"`
	LargeSize             float64                               `plist:"largesize"`
	Magnification         bool                                  `plist:"magnification"`
	ShowRecent            bool                                  `plist:"show-recents"`
	ModCount              int                                   `plist:"mod-count"`
	MruSpace              bool                                  `plist:"mru-spaces"`
	TrashFull             bool                                  `plist:"trash-full"`
	Version               int                                   `plist:"version"`
	WvousBrCorner         int                                   `plist:"wvous-br-corner"`
	WvousBrModifier       int                                   `plist:"wvous-br-modifier"`
	PersistentApps        []*DockPersistent[*DockAppTileData]   `plist:"persistent-apps"`
	RecentApps            []*DockPersistent[*DockAppTileData]   `plist:"recent-apps"`
	PersistentOthers      []*DockPersistent[*DockOtherTileData] `plist:"persistent-others"`
}

type DockPersistent[T any] struct {
	Guid     int    `plist:"guid"`
	TileData T      `plist:"tile-data"`
	TileType string `plist:"tile-type"`
}

type DockAppTileData struct {
	Book             []byte        `plist:"book"`
	BundleIdentifier string        `plist:"bundle-identifier"`
	DockExtra        bool          `plist:"dock-extra"`
	FileData         *DockFileData `plist:"file-data"`
	FileLabel        string        `plist:"file-label"`
	FileModDate      int           `plist:"file-mod-date"`
	FileType         int           `plist:"file-type"`
	IsBeta           bool          `plist:"is-beta"`
	ParentModDate    int           `plist:"parent-mod-date"`
}

type DockOtherTileData struct {
	Arrangement       int           `plist:"arrangement"`
	Book              []byte        `plist:"book"`
	DisplayAs         int           `plist:"displayas"`
	FileData          *DockFileData `plist:"file-data"`
	FileLabel         string        `plist:"file-label"`
	FileModDate       int           `plist:"file-mod-date"`
	FileType          int           `plist:"file-type"`
	IsBeta            bool          `plist:"is-beta"`
	ParentModDate     int           `plist:"parent-mod-date"`
	PreferredItemSize int           `plist:"preferreditemsize"`
	ShowAs            int           `plist:"showas"`
}

type DockFileData struct {
	CfUrlString     string `plist:"_CFURLString"`
	CfUrlStringType int    `plist:"_CFURLStringType"`
}
