package configs

type subjectsLinks struct {
	Subjects map[string]string
}

var SubjectsLinks subjectsLinks

func init() {
	SubjectsLinks.Subjects = make(map[string]string)
	SubjectsLinks.Subjects["БАЗЫ ДАННЫХ"] = "subject/databases"
	SubjectsLinks.Subjects["РПиРП"] = "subject/rprp"
}
