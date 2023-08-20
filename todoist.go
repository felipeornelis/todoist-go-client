package todoist

type Todoist struct {
	authToken string
}

func New(authToken string) Todoist {
	return Todoist{
		authToken: authToken,
	}
}

// type todoist struct {
// 	token string
// 	Task  task.Task
// }

// func New(token string) todoist {
// 	return todoist{
// 		token: token,
// 	}
// }
