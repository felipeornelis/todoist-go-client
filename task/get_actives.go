package task

type GetManyParams struct {
	ProjectID string
	SectionID string
	Label     string
	Filter    string
	Lang      string
	Ids       []int
}

func (t task) GetMany(params GetManyParams) (task, error) {
	return task{}, nil
}
