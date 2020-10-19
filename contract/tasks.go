package contract

import (
	"github.com/brunvieira/lotus"
	"time"
)

type TaskId struct {
	id string
}

type Task struct {
	*TaskId
	Checked bool
	Title string
	Description string
	Date time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

var (
	CreateTaskRouteContract = lotus.RouteContract{
		Label: "CreateTaskRoute",
		Description: "Creates a task",
		Method: lotus.POST,
		Path: "/add",
		Data: Task{},
	}
	DeleteTaskRouteContract = lotus.RouteContract{
		Label: "DeleteTaskRoute",
		Description: "Deletes a task",
		Method: lotus.DELETE,
		Path: "/remove/:id",
		Data: TaskId{},
	}
	GetTaskRouteContract = lotus.RouteContract{
		Label: "GetTaskRoute",
		Description: "Retrieves a task identified by id",
		Method: lotus.GET,
		Path: "/get/:id",
		Data: TaskId{},
	}
	GetTasksFromUserRouteContract = lotus.RouteContract{
		Label: "GetTasksFromUserRoute",
		Description: "Retrieves tasks for the user identified by id",
		Method: lotus.GET,
		Path: "/from-user/:id",
		Data: UserId{},
	}
	UpdateTaskRouteContract = lotus.RouteContract{
		Label: "UpdateTaskRoute",
		Description: "Updates the task identified by id",
		Method: lotus.PUT,
		Path: "update/:id",
		Data: Task{},
	}
	TasksServiceContract = lotus.ServiceContract{
		Label: "TasksService",
		Description: "Creates, Deletes and Retrieves tasks",
		Host: "localhost",
		Namespace: "tasks",
		Port: 8082,
		RoutesContracts: []lotus.RouteContract{
			CreateTaskRouteContract,
			DeleteTaskRouteContract,
			GetTaskRouteContract,
			GetTasksFromUserRouteContract,
			UpdateTaskRouteContract,
		},
	}
)
