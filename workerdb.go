package main

import "time"

type Project struct {
	Id   int64
	Name string
}

func (p *Project) create() {

}

func (p *Project) delete() {

}

func (p *Project) update() {

}

// this struct relatest to the Project
type Task struct {
	Id   int64
	Name string
}

func (p *Task) create() {

}

func (p *Task) delete() {

}

func (p *Task) update() {

}

// this struct relates to the Task
type TimePoint struct {
	Id        int64
	Name      string
	StartTime *time.Time
	EndTime   *time.Time
	TakId     int64
}

func (a *TimePoint) startTime() {

}

func (a *TimePoint) endTime() {

}
