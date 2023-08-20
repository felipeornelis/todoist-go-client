package todoist

type Comment struct {
	ID         string
	TaskID     string
	ProjectID  string
	PostedAt   string
	Content    string
	Attachment CommentAttachment
}

type CommentAttachment struct {
	FileName     string
	FileType     string
	FileURL      string
	ResourceType string
}
