package configs

type links struct {
	Name string
	Link string
}
type subjectsForStudents struct {
	Subjects map[string][]links
}

var SubjectsForStudents subjectsForStudents

func init() {
	SubjectsForStudents.Subjects = make(map[string][]links)
	SubjectsForStudents.Subjects["ИУ9-61Б"] = []links{{"БАЗЫ ДАННЫХ", "subject/databases"}, {"РПиРП", "subject/rprp"}}
	SubjectsForStudents.Subjects["ИУ9-62Б"] = []links{{"БАЗЫ ДАННЫХ", "subject/databases"}, {"РПиРП", "subject/rprp"}}
}
