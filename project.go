package todoist

type Project struct {
	ID             string
	Name           string
	Color          string
	ParentID       string
	Order          uint8
	CommentCount   int
	IsShared       bool
	IsFavorite     bool
	IsInboxProject bool
	IsTeamInbox    bool
	ViewStyle      string
	URL            string
}
