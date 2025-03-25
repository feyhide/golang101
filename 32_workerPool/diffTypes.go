package main

import (
	"fmt"
	"sync"
	"time"
)

type Task2 interface {
	Process2()
}

// EmailTask struct
type EmailTask struct {
	Email   string
	Subject string
	Message string
}

// Implement Process2 for EmailTask
func (t *EmailTask) Process2() {
	fmt.Println("📩 Sending Email to", t.Email)
	time.Sleep(time.Second * 2)
}

// ImageProcessingTask struct
type ImageProcessingTask struct {
	ImageURL string
}

// Implement Process2 for ImageProcessingTask
func (t *ImageProcessingTask) Process2() {
	fmt.Println("🖼️ Processing image:", t.ImageURL)
	time.Sleep(time.Second * 5)
}

type WorkerPool2 struct {
	Tasks       []Task2
	concurrency int
	tasksChan   chan Task2
	wg          sync.WaitGroup
}

func (wp *WorkerPool2) worker2() {
	for task := range wp.tasksChan {
		task.Process2() // Calls Process2() based on Task2 interface
		wp.wg.Done()
	}
}

func (wp *WorkerPool2) Run2() {
	wp.tasksChan = make(chan Task2, len(wp.Tasks))

	for i := 0; i < wp.concurrency; i++ {
		go wp.worker2()
	}

	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}
	close(wp.tasksChan)

	wp.wg.Wait()
}

func diffTypes() {
	tasks := []Task2{
		&EmailTask{"user1@example.com", "Hello!", "Welcome to our service."},
		&ImageProcessingTask{"https://example.com/image1.jpg"},
		&EmailTask{"user2@example.com", "Reminder", "Don't forget your meeting."},
		&EmailTask{"user3@example.com", "Hello!", "Welcome to our service."},
		&ImageProcessingTask{"https://example.com/image2.jpg"},
		&EmailTask{"user4@example.com", "Reminder", "Don't forget your meeting."},
		&EmailTask{"user5@example.com", "Hello!", "Welcome to our service."},
		&ImageProcessingTask{"https://example.com/image3.jpg"},
		&EmailTask{"user6@example.com", "Reminder", "Don't forget your meeting."},
		&EmailTask{"user7@example.com", "Hello!", "Welcome to our service."},
		&ImageProcessingTask{"https://example.com/image4.jpg"},
		&EmailTask{"user8@example.com", "Reminder", "Don't forget your meeting."},
		&ImageProcessingTask{"https://example.com/image5.jpg"},
		&EmailTask{"user9@example.com", "Hello!", "Welcome to our service."},
		&ImageProcessingTask{"https://example.com/image6.jpg"},
		&EmailTask{"user10@example.com", "Reminder", "Don't forget your meeting."},
		&ImageProcessingTask{"https://example.com/image7.jpg"},
	}

	wp := WorkerPool2{
		Tasks:       tasks,
		concurrency: 5,
	}

	wp.Run2()
	fmt.Println("✅ All tasks processed!")
}
