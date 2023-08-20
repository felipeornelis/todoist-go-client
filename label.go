package todoist

type Label struct {
	ID         string
	Name       string
	Color      string
	Order      uint8
	IsFavorite bool
}
