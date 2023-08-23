package todoist

import "time"

const (
	BASE_URL    = "https://api.todoist.com/rest/v2"
	MAX_TIMEOUT = time.Second * 15
)

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
