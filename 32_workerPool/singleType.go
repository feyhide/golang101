package main

import (
	"fmt"
	"sync"
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

// worker pool def
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}

// func to execute worker pool

func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	//initialize task channel
	wp.tasksChan = make(chan Task, len(wp.Tasks))

	// Tell WaitGroup how many tasks we have (before starting workers)
	wp.wg.Add(len(wp.Tasks))

	//start workers
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	//send task in channels
	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}
	close(wp.tasksChan)

	//wait for all task to finish
	wp.wg.Wait()
}

func singleType() {
	//create tasks
	tasks := make([]Task, 20)
	for i := 0; i < 20; i++ {
		tasks[i] = Task{ID: i + 1}
	}

	//create workerpool
	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5, // no of workers running at the same time
	}

	// run the pool
	wp.Run()
	fmt.Println("All tasks processed")
}
