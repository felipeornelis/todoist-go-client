package task

type Task struct {
	ID           string
	ProjectID    string
	SectionID    string
	Content      string
	Description  string
	IsCompleted  bool
	Labels       []string
	ParentID     string
	Order        uint8
	Priority     uint8
	Due          taskDue
	URL          string
	CommentCount int
	CreatedAt    string
	AssigneeID   string
	AssignerID   string
	Duration     taskDuration
}

type taskDue struct {
	String      string
	Date        string
	IsRecurring bool
	Datetime    string
	Timezone    string
}

type taskDuration struct {
	Amount uint
	Unit   string
}
