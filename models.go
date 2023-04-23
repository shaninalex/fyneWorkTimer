package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase(db_path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	// define schema
	scheme := `
	CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY,
		name VARCHAR(64) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY,
		name VARCHAR(64),
		project_id INTEGER NOT NULL,
		time_estimate INTEGER,
	    FOREIGN KEY (project_id)
			REFERENCES projects (id)
	);
	CREATE TABLE IF NOT EXISTS timepointes (
		id INTEGER PRIMARY KEY,
		task_id INTEGER NOT NULL,
		start_time INTEGER,
		end_time INTEGER,
	    FOREIGN KEY (task_id)
			REFERENCES task (id)	
	)
	`
	_, err = db.Exec(scheme)
	if err != nil {
		log.Fatal(err)
		// return nil, err
	}

	return db, nil

}

type Project struct {
	Id   int64
	Name string
}

func GetAllProjects(db *sql.DB) ([]Project, error) {

	rows, err := db.Query(`SELECT * FROM projects`)
	if err != nil {
		log.Printf("db query failed %s", err.Error())
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var all []Project
	for rows.Next() {
		var project Project
		if err := rows.Scan(&project.Id, &project.Name); err != nil {
			return nil, err
		}
		all = append(all, project)
	}

	return all, nil

}

func (p *Project) Create(db *sql.DB) error {

	res, err := db.Exec(`INSERT INTO projects (name) VALUES (?)`, p.Name)
	if err != nil {
		log.Fatal(err)
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return err
	}
	p.Id = id
	return nil
}

func (p *Project) Delete(db *sql.DB) error {
	_, err := db.Exec(`DELETE FROM projects WHERE id = ?`, p.Id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (p *Project) Update(db *sql.DB) error {
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

func (p *Task) Create() {

}

func (p *Task) Delete() {

}

func (p *Task) Update() {

}

// this struct relates to the Task
type TimePoint struct {
	Id        int64
	Name      string
	StartTime *time.Time
	EndTime   *time.Time
	TakId     int64
}

func (a *TimePoint) StartTimer() {

}

func (a *TimePoint) EndTimer() {

}
