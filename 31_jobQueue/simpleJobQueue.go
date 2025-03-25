package main

import (
	"fmt"
	"time"
)

// task def
type Task struct {
	ID int
}

// process of task
func (t *Task) Process() {
	fmt.Println("Processing Task", t.ID)
	time.Sleep(time.Second * 2)
}

// Job Queue def
type JobQueue struct {
	tasks []Task
}

// Add a job to the queue
func (jq *JobQueue) AddTask(task Task) {
	jq.tasks = append(jq.tasks, task)
}

func (jq *JobQueue) Run() {
	for _, task := range jq.tasks {
		task.Process()
	}
	fmt.Println("All tasks processed!")
}

func simpleJobQueue() {
	//create tasks
	jq := JobQueue{}

	// Add tasks to the queue
	for i := 1; i <= 5; i++ {
		jq.AddTask(Task{ID: i})
	}

	// Run the queue (sequential processing)
	jq.Run()
}
