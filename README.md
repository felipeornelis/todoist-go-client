# Todoist Client Go

This is an **unofficial** version of Go client for Todoist app. It's still in progress so it's definitely not recommended to be used in production.

## How to install
```bash
go get github.com/felipeornelis/todoist-go-client
```

### How to use

Simple example of how to initialise a Todoist client and fetch some account's tasks:

```go
package main

import (
    "log"
    "fmt"

    "github.com/felipeornelis/todoist-go-client"
)

func main() {
    t := todoist.New("<your token goes here>")

    tasks, err := t.GetTasks()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(tasks)
}
```

## Documentation

### Tasks

A task has a `Task` type with the following fields:

| Field | Type   |
|-------|----------------------|
| ID    | string |
| ProjectID |    string |
| SectionID |    string |
| Content |      string |
| Description |  string |
| IsCompleted |  bool |
| Labels |       []string |
| ParentID |     string |
| Order |        uint8 |
| Priority |     uint8 |
| Due |          taskDue |
| URL |          string |
| CommentCount | int |
| CreatedAt |    string       |
| AssigneeID |   string       |
| AssignerID |   string       |
| Duration |     taskDuration |

#### Add new task

To add a new task, the right method for it is the `AddTask(args AddTaskArgs)` method, which expects a parameter of type `AddTaskArgs`.

`AddTaskArgs` is a struct and only the `Content` field is required, while the other ones are optionals and might be omitted.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Content | string | Yes | Task content, which may contain either markdown-formatted text and hyperlinks; |
| Description | string | No | A description for the task, which also may container either markdown-formatted text and hyperlinks. |
| ProjectID | string | No | The ID of the project where the task should be set. If not set, then it is put to inbox. |
| SectionID | string | No | The ID of the section to put task into. |
| ParentID | string | No | The ID of task's parent. |
| Order | uint8 | No | Non-zero integer value used to sort tasks under the same parent. It is used by Todoist's clients (mobile and web). |
| Labels | []string | No | A list of words that may represent either personal or shared labels. |
| Priority | uint8 | No | Task priority from 0 to 4, where 0 means normal priority and 4 urgent. |
| DueString | string | No | Human readable task due date. It is set using local time, not UTC. Read more on the [official documentation](https://todoist.com/help/articles/due-dates-and-times). |
| DueDate | string | No | Specific due date in `YYYY-MM-DD` format. As `DueString`, it is set on user's local time. |
| DueDatetime | string | No | Specific date and time. It must follow [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) specifications. |
| DueLang | string | No | 2-letter code to specify what language `DueString` is written in. If it's written in English, then the field might be omitted.  |
| AssigneeID | string | No | ID of the user who is responsible for the task. Useful only for shared tasks. |
| Duration | uint | No | A positive integer number that represents the amount of `DurationUnit` the task will take. If specified, then `DurationUnit` must be defined. |
| DurationUnit | string | No | The unit of time that represents `Duration` field. Must be either `minute` or `day`. If specified, then `duration` must be defined. |

<!-- **Note**: if `Duration` field is specified, you must also specify `DurationUnit` and vice-versa. `DurationUnit` must be either `minute` or `day`. Check the [official documentation](https://developer.todoist.com/rest/v2/#create-a-new-task) for more information. -->
**Note:** `DueString`, `DueDate` and `DueDatetime` are exclusive fields, that is, only one of them must be filled in.

Sample of how to add a new task:

```go
package main

import (
    "log"
    "fmt"

    "github.com/felipeornelis/todoist-go-client"
)

func main() {
    t := todoist.New("<your token goes here, as I said before>")

    args := todoist.AddTaskArgs{
        Content: "Run the world after lunch"
    }
    task, err := t.AddTask(args)
    if err != nil {
        log.Fatal(er)
    }

    fmt.Println(task)
}
```

## Feedback

This package is under development, so any feedback is welcome. It can be reported as *Issues* in this repository or you can reach me on *hello@felipeornelis.com*.

