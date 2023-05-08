package dataBase 

import (
    "database/sql"
)

type Module struct {
    Id []uint8
	SubjectID []uint8
	Name string
	MaxScore int
	MinScore int
	Db *sql.DB
}

func (m Module) Get_id() ([]uint8, []uint8) { return m.Id, m.SubjectID}
func (m Module) Get_name() string { return m.Name }
func (m Module) Get_max_score() int { return m.MaxScore}
func (m Module) Get_min_score() int { return m.MinScore}

func (m *Module) Set_name(name string) {
    m.Name = name
    _, err := m.Db.Exec("update Module SET Name = $1 where ModuleID = $2", m.Name, m.Id)
    if err != nil {
        panic(err)
    }
}

func (m *Module) Set_max_score(score int) {
    m.MaxScore = score
    _, err := m.Db.Exec("update Module SET MaxScore = $1 where ModuleID = $2", m.MaxScore, m.Id)
    if err != nil {
        panic(err)
    }
}

func (m *Module) Set_min_score(score int) {
    m.MinScore = score
    _, err := m.Db.Exec("update Module SET MinScore = $1 where ModuleID = $2", m.MinScore, m.Id)
    if err != nil {
        panic(err)
    }
}
