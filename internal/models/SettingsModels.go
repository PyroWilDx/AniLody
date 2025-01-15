package models

type UserSettings struct {
	UserName string
	UserSite string

	OutPath      string
	ThreadsCount int

	MusicNameFormat string
	CapWords        bool
	LowWords        bool
	FmtNums         bool

	ApplyImage        bool
	UpScaleImageWidth string

	IncOp bool
	IncEd bool

	MinScore   float32
	MaxScore   float32
	StatusList []string
}
