package models

type UserSettings struct {
	UserName string
	UserSite string

	OutPath         string
	MusicNameFormat string
	CapWords        bool
	AddImage        bool

	IncOp bool
	IncEd bool
}
