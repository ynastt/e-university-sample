CREATE TABLE IF NOT EXISTS Users (
	UserID UUID PRIMARY KEY,
	Login VARCHAR(50) NOT NULL,
	Passw VARCHAR(50) NOT NULL,
	UsersRights Boolean NOT NULL,
	CONSTRAINT user_unique UNIQUE(Login) --именование ограничение табилцы
);

CREATE TABLE IF NOT EXISTS StudentGroup (
	GroupID UUID PRIMARY KEY,
	GroupName VARCHAR NOT NULL UNIQUE,
	YearOfAdmission INT NOT NULL,
	Course INT NOT NULL,
	AmountOfStudents INT NOT NULL
);

CREATE TABLE IF NOT EXISTS Student (
	StudentID UUID PRIMARY KEY,
	StudentName VARCHAR(50) NOT NULL,
	Surname VARCHAR(50) NOT NULL,
	Patronymic VARCHAR(50),
	Email VARCHAR(255) UNIQUE NOT NULL,
	Phone VARCHAR(11) NOT NULL,
	YearOfAdmission INT NOT NULL,
	PassedCourses INT NOT NULL,
	NumInGroup INT NOT NULL,
	user_id UUID NOT NULL REFERENCES Users(UserID),
	group_id UUID NOT NULL REFERENCES StudentGroup(GroupID),
	CONSTRAINT student_unique UNIQUE(user_id, group_id)
);

CREATE TABLE IF NOT EXISTS Teacher (
	TeacherID UUID PRIMARY KEY,
	TeacherName VARCHAR(50) NOT NULL,
	Surname VARCHAR(50) NOT NULL,
	Patronymic VARCHAR(50),
	Email VARCHAR(255) UNIQUE NOT NULL,
	user_id UUID  NOT NULL REFERENCES Users(UserID)
);

CREATE TABLE IF NOT EXISTS Subject (
	SubjectID UUID PRIMARY KEY,
	Description TEXT NOT NULL, 
	SubjectProgram TEXT NOT NULL, 
	NumberOfHours INT NOT NULL,
	NumberOfCredits INT NOT NULL 
);

CREATE TABLE IF NOT EXISTS Modules (
	ModuleID UUID NOT NULL UNIQUE,
	subject_id UUID NOT NULL REFERENCES Subject(SubjectID) ON DELETE RESTRICT,
	ModuleName Varchar(50) NOT NULL, 
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	--CONSTRAINT module_unique UNIQUE(ModuleID, subject_id)
	PRIMARY KEY(ModuleID, subject_id)
);

CREATE TABLE IF NOT EXISTS Exam (
	ExamID UUID PRIMARY KEY, 
	Questions TEXT NOT NULL,
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	ExamDate date,
	subject_id UUID NOT NULL REFERENCES Subject(SubjectID) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS Lecture (
	LectureID UUID PRIMARY KEY,
	Theme TEXT NOT NULL,
	LectureText TEXT NOT NULL, 
	module_id UUID NOT NULL REFERENCES Modules(ModuleID) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS Seminar (
	SeminarID UUID PRIMARY KEY, 
	Theme TEXT NOT NULL,
	SeminarText TEXT NOT NULL, 
	module_id UUID NOT NULL REFERENCES Modules(ModuleID) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS Lab (
	LabID UUID PRIMARY KEY,
	LabName TEXT NOT NULL,
	LabText TEXT NOT NULL, 
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	LabDate date NOT NULL,
	Deadline date NOT NULL,
	module_id UUID NOT NULL REFERENCES Modules(ModuleID) ON DELETE RESTRICT --запрет на удаление модуля через таблицу лабы
);

CREATE TABLE IF NOT EXISTS BC (
	BCID uuid PRIMARY KEY, 
	Theme TEXT NOT NULL,
	Questions TEXT NOT NULL, 
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	module_id UUID NOT NULL REFERENCES Modules(ModuleID) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS CourseProject (
	ProjectID UUID PRIMARY KEY, 
	Subject VARCHAR(50) NOT NULL,
	Description TEXT NOT NULL, 
	NumberOfHours INT NOT NULL,
	StartDate DATE NOT NULL,
	Deadline DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS Queue (
	QueueID UUID PRIMARY KEY,
	StartDate DATE NOT NULL
);

-- dop
CREATE TABLE IF NOT EXISTS StudentInQueue (
	student_id UUID NOT NULL REFERENCES Student(StudentID) ON DELETE RESTRICT,
	queue_id UUID NOT NULL REFERENCES Queue(QueueID) ON DELETE RESTRICT,
	NumInQueue Int NOT NULL,
	Task INT NOT NULL, -- делать enum
	PRIMARY KEY(student_id, queue_id)
);

CREATE TABLE IF NOT EXISTS TeacherSubject (
	teacher_id UUID NOT NULL REFERENCES Teacher(TeacherID) ON DELETE RESTRICT,
	subject_id UUID NOT NULL REFERENCES Subject(SubjectID) ON DELETE RESTRICT,
	TeacherRole Int NOT NULL,
	PRIMARY KEY(teacher_id, subject_id)
);

CREATE TABLE IF NOT EXISTS Supervisor (
	teacher_id UUID NOT NULL REFERENCES Teacher(TeacherID) ON DELETE RESTRICT,
	project_id UUID NOT NULL REFERENCES CourseProject (ProjectID) ON DELETE CASCADE,
	SupervisorRole Int NOT NULL,
	PRIMARY KEY(teacher_id, project_id)
);

CREATE TABLE IF NOT EXISTS StudentCourseProject (
	student_id UUID NOT NULL REFERENCES Student(StudentID) ON DELETE RESTRICT,
	project_id UUID NOT NULL REFERENCES CourseProject(ProjectID) ON DELETE CASCADE,
	ProjAssignment TEXT NOT NULL,
	TitleOfProject VARCHAR(100) NOT NULL,
	RecievedScore INT NOT NULL,
	DateOdPassing DATE,
	PRIMARY KEY(student_id, project_id)
);

CREATE TABLE IF NOT EXISTS LectureAttendance (
	student_id UUID NOT NULL REFERENCES Student(StudentID) ON DELETE RESTRICT,
	lecture_id UUID NOT NULL REFERENCES Lecture(LectureID) ON DELETE RESTRICT,
	WasAttended BOOL NOT NULL,
	BonusScore INT NOT NULL,
	PRIMARY KEY(student_id, lecture_id)
);

CREATE TABLE IF NOT EXISTS SeminarAttendance (
	student_id UUID NOT NULL REFERENCES Student(StudentID) ON DELETE RESTRICT,
	seminar_id UUID NOT NULL REFERENCES Seminar(SeminarID) ON DELETE RESTRICT,
	WasAttended BOOL NOT NULL,
	BonusScore INT NOT NULL,
	PRIMARY KEY(student_id, seminar_id)
);

CREATE TABLE IF NOT EXISTS LabInstance (
	student_id UUID NOT NULL REFERENCES Student(StudentID) ON DELETE RESTRICT,
	lab_id UUID NOT NULL REFERENCES Lab(LabID) ON DELETE RESTRICT,
	DateOdPassing DATE NOT NULL,
	NumOfInstance INT NOT NULL,
	RecievedScore INT NOT NULL,
	Variant INT,
	Remarks TEXT,
	BonusScore INT NOT NULL,
	PRIMARY KEY(student_id, lab_id, NumOfInstance) -- вместо даты в ключ переменную попытки сдачи, а дату оставить атрибутом
);

CREATE TABLE IF NOT EXISTS BCInstance (
	student_id UUID UNIQUE NOT NULL REFERENCES Student(StudentID) ON DELETE RESTRICT,
	bc_id UUID UNIQUE NOT NULL REFERENCES BC(BCID) ON DELETE RESTRICT,
	DateOdPassing DATE NOT NULL,
	NumOfInstance INT UNIQUE NOT NULL,
	RecievedScore INT NOT NULL,
	Variant INT,
	PRIMARY KEY(student_id, bc_id, NumOfInstance) -- вместо даты в ключ переменную попытки сдачи, а дату оставить атрибутом
);

CREATE TABLE IF NOT EXISTS ExamInstance (
	student_id UUID NOT NULL REFERENCES Student(StudentID) ON DELETE RESTRICT,
	exam_id UUID NOT NULL REFERENCES Exam(ExamID) ON DELETE RESTRICT,
	DateOdPassing DATE NOT NULL,
	NumOfInstance INT UNIQUE NOT NULL,
	RecievedScore INT NOT NULL,
	TicketNumber INT,
	PRIMARY KEY(student_id, exam_id, NumOfInstance) -- вместо даты в ключ переменную попытки сдачи, а дату оставить атрибутом
);

CREATE TABLE IF NOT EXISTS Task (
	student_id UUID NOT NULL REFERENCES Student(StudentID) ON DELETE RESTRICT,
	subject_id UUID NOT NULL REFERENCES Subject(SubjectID) ON DELETE RESTRICT,
	TaskID UUID NOT NULL, -- можно сделать это PRIMARY KEY и оставить внешние ключи 
	Description TEXT NOT NULL, -- какой тип данных?
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	Deadline DATE NOT NULL,
	RecievedScore INT NOT NULL,
	PRIMARY KEY(student_id, subject_id, TaskID) --мне кажется это бредом
);

CREATE TABLE IF NOT EXISTS TaskExam (
	exam_id UUID NOT NULL REFERENCES Exam(ExamID) ON DELETE RESTRICT,
	TaskID UUID NOT NULL, -- можно сделать это PRIMARY KEY и оставить внешние ключи
	Description TEXT NOT NULL, 
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	Deadline DATE NOT NULL,
	RecievedScore INT NOT NULL,
	PRIMARY KEY(exam_id, TaskID) --мб тоже см Task
);

CREATE TABLE IF NOT EXISTS TaskBC (
	bc_id UUID NOT NULL REFERENCES BCInstance(bc_id) ON DELETE RESTRICT,
	TaskID UUID NOT NULL, -- можно сделать это PRIMARY KEY и оставить внешние ключи
	Description TEXT NOT NULL, 
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	Deadline DATE NOT NULL,
	RecievedScore INT NOT NULL,
	PRIMARY KEY(bc_id, TaskID) --мб тоже см Task
);


INSERT INTO StudentGroup(GroupID, GroupName, YearOfAdmission, Course, AmountOfStudents) VALUES
	(gen_random_uuid(),'ИУ9-61Б', 2020, 3, 28),
	(gen_random_uuid(), 'ИУ9-62Б', 2020, 3, 24);

SELECT * FROM StudentGroup;