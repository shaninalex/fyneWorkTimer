package main

import (
	"database/sql"
	"log"
	"time"
)

type Project struct {
	Id   int64
	Name string
}

func (p *Project) create(db *sql.DB) error {

	res, err := db.Exec(`INSERT INTO projects (name) VALUES (?)`, p.Name)
	if err != nil {
		log.Fatal(err)
		return err
	}
	id, err := res.LastInsertId()
	p.Id = id
	return nil
}

func (p *Project) delete(db *sql.DB) error {
	_, err := db.Exec(`DELETE FROM projects WHERE id = ?`, p.Id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (p *Project) update(db *sql.DB) error {
	_, err := db.Exec(`UPDATE projects SET name = '?' WHERE id = ?`, p.Name, p.Id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// this struct relatest to the Project
type Task struct {
	Id           int64
	Name         string
	ProjectId    int64
	TimeEstimate int64
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
