drop table Users cascade;
CREATE TABLE IF NOT EXISTS Users (
	UserID UUID PRIMARY KEY,
	Login VARCHAR(50) NOT NULL,
	Passw VARCHAR(50) NOT NULL,
	UsersRights INT NOT NULL,
	CONSTRAINT user_unique UNIQUE(Login) --именование ограничение табилцы
);

CREATE TABLE IF NOT EXISTS StudentGroup (
	GroupID UUID PRIMARY KEY,
	GroupName VARCHAR NOT NULL,
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
	Description TEXT NOT NULL UNIQUE,  --название предмета
	SubjectProgram TEXT NOT NULL, 
	NumberOfHours INT NOT NULL,
	NumberOfCredits INT NOT NULL, 
	CONSTRAINT subject_unique UNIQUE(SubjectID, Description)
);

CREATE TABLE IF NOT EXISTS Modules (
	ModuleID UUID NOT NULL UNIQUE,
	subject_id UUID NOT NULL REFERENCES Subject(SubjectID) ON DELETE CASCADE
	,
	ModuleName Varchar(50) NOT NULL UNIQUE, 
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
	CONSTRAINT exam_unique UNIQUE(ExamID, ExamDate),
	subject_id UUID NOT NULL REFERENCES Subject(SubjectID) ON DELETE CASCADE

);

CREATE TABLE IF NOT EXISTS Lecture (
	LectureID UUID PRIMARY KEY,
	Theme TEXT NOT NULL UNIQUE,
	LectureText TEXT NOT NULL, 
	module_id UUID NOT NULL REFERENCES Modules(ModuleID) ON DELETE CASCADE
	,
	LectNumber INT NOT NULL,
	LectDate date,
	CONSTRAINT lect_unique UNIQUE(Theme);
);

CREATE TABLE IF NOT EXISTS Seminar (
	SeminarID UUID PRIMARY KEY, 
	Theme TEXT NOT NULL UNIQUE,
	SeminarText TEXT NOT NULL, 
	module_id UUID NOT NULL REFERENCES Modules(ModuleID) ON DELETE CASCADE
	,
	SemNumber INT NOT NULL,
	SemDate date,
	CONSTRAINT sem_unique UNIQUE(Theme);
);

CREATE TABLE IF NOT EXISTS Lab (
	LabID UUID PRIMARY KEY,
	LabNumber INT NOT NULL,
	LabName TEXT NOT NULL,
	LabText TEXT NOT NULL, 
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	LabDate date NOT NULL,
	Deadline date NOT NULL,
	CONSTRAINT lab_unique UNIQUE(LabName),
	module_id UUID NOT NULL REFERENCES Modules(ModuleID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS BC (
	BCID uuid PRIMARY KEY, 
	Theme TEXT NOT NULL,
	Questions TEXT NOT NULL, 
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	CONSTRAINT rk_unique UNIQUE(Theme),
	module_id UUID NOT NULL REFERENCES Modules(ModuleID) ON DELETE CASCADE
	,
	BCNum int
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
	StartDate timestamp with time zone NOT NULL,
	subject_id UUID NOT NULL REFERENCES Subject(SubjectID) ON DELETE CASCADE
);

-- dop
CREATE TABLE IF NOT EXISTS StudentInQueue (
	student_id UUID NOT NULL REFERENCES Student(StudentID) ON DELETE CASCADE,
	queue_id UUID NOT NULL REFERENCES Queue(QueueID) ON DELETE CASCADE,
	NumInQueue Int NOT NULL,
	PRIMARY KEY(student_id, queue_id)
);

CREATE TABLE IF NOT EXISTS TeacherSubjectGroup (
	teacher_id UUID NOT NULL REFERENCES Teacher(TeacherID) ON DELETE CASCADE,
	subject_id UUID NOT NULL REFERENCES Subject(SubjectID) ON DELETE CASCADE,
	group_id UUID NOT NULL REFERENCES StudentGroup(GroupID) ON DELETE CASCADE,
	TeacherRole Int NOT NULL,
	PRIMARY KEY(teacher_id, subject_id, group_id)
);

CREATE TABLE IF NOT EXISTS Supervisor (
	teacher_id UUID NOT NULL REFERENCES Teacher(TeacherID) ON DELETE CASCADE
	,
	project_id UUID NOT NULL REFERENCES CourseProject (ProjectID) ON DELETE CASCADE,
	SupervisorRole Int NOT NULL,
	PRIMARY KEY(teacher_id, project_id)
);

CREATE TABLE IF NOT EXISTS StudentCourseProject (
	student_id UUID NOT NULL REFERENCES Student(StudentID) ON DELETE CASCADE
	,
	project_id UUID NOT NULL REFERENCES CourseProject(ProjectID) ON DELETE CASCADE,
	ProjAssignment TEXT NOT NULL,
	TitleOfProject VARCHAR(100) NOT NULL,
	RecievedScore INT NOT NULL,
	DateOdPassing DATE,
	PRIMARY KEY(student_id, project_id)
);

CREATE TABLE IF NOT EXISTS LectureAttendance (
	student_id UUID NOT NULL REFERENCES Student(StudentID) ON DELETE CASCADE
	,
	lecture_id UUID NOT NULL REFERENCES Lecture(LectureID) ON DELETE CASCADE
	,
	WasAttended BOOL NOT NULL,
	BonusScore INT,
	PRIMARY KEY(student_id, lecture_id)
);

CREATE TABLE IF NOT EXISTS SeminarAttendance (
	student_id UUID NOT NULL REFERENCES Student(StudentID) ON DELETE CASCADE
	,
	seminar_id UUID NOT NULL REFERENCES Seminar(SeminarID) ON DELETE CASCADE
	,
	WasAttended BOOL NOT NULL,
	BonusScore INT,
	PRIMARY KEY(student_id, seminar_id)
);

CREATE TABLE IF NOT EXISTS LabInstance (
	student_id UUID UNIQUE NOT NULL REFERENCES Student(StudentID) ON DELETE CASCADE
	,
	lab_id UUID UNIQUE NOT NULL REFERENCES Lab(LabID) ON DELETE CASCADE
	,
	DateOfPassing DATE NOT NULL,
	NumOfInstance INT UNIQUE NOT NULL,
	RecievedScore INT NOT NULL,
	Variant INT,
	Remarks TEXT,
	BonusScore INT,
	PRIMARY KEY(student_id, lab_id, NumOfInstance) -- вместо даты в ключ переменную попытки сдачи, а дату оставить атрибутом
);

CREATE TABLE IF NOT EXISTS BCInstance (
	student_id UUID UNIQUE NOT NULL REFERENCES Student(StudentID) ON DELETE CASCADE
	,
	bc_id UUID UNIQUE NOT NULL REFERENCES BC(BCID) ON DELETE CASCADE
	,
	DateOfPassing DATE NOT NULL,
	NumOfInstance INT UNIQUE NOT NULL,
	RecievedScore INT NOT NULL,
	Variant INT,
	Remarks TEXT,
	PRIMARY KEY(student_id, bc_id, NumOfInstance) -- вместо даты в ключ переменную попытки сдачи, а дату оставить атрибутом
);

CREATE TABLE IF NOT EXISTS ExamInstance (
	student_id UUID UNIQUE NOT NULL REFERENCES Student(StudentID) ON DELETE CASCADE
	,
	exam_id UUID UNIQUE NOT NULL REFERENCES Exam(ExamID) ON DELETE CASCADE
	,
	DateOfPassing DATE NOT NULL,
	NumOfInstance INT UNIQUE NOT NULL,
	RecievedScore INT NOT NULL,
	TicketNumber INT,
	PRIMARY KEY(student_id, exam_id, NumOfInstance) -- вместо даты в ключ переменную попытки сдачи, а дату оставить атрибутом
);

CREATE TABLE IF NOT EXISTS Task (
	student_id UUID NOT NULL REFERENCES Student(StudentID) ON DELETE CASCADE
	,
	subject_id UUID NOT NULL REFERENCES Subject(SubjectID) ON DELETE CASCADE
	,
	TaskID UUID NOT NULL, -- можно сделать это PRIMARY KEY и оставить внешние ключи 
	Description TEXT NOT NULL, -- какой тип данных?
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	Deadline DATE NOT NULL,
	RecievedScore INT NOT NULL,
	PRIMARY KEY(student_id, subject_id, TaskID) --мне кажется это бредом
);

CREATE TABLE IF NOT EXISTS TaskExam (
	exam_id UUID NOT NULL REFERENCES Exam(ExamID) ON DELETE CASCADE
	,
	TaskID UUID NOT NULL, -- можно сделать это PRIMARY KEY и оставить внешние ключи
	Description TEXT NOT NULL, 
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	Deadline DATE NOT NULL,
	RecievedScore INT NOT NULL,
	PRIMARY KEY(exam_id, TaskID) --мб тоже см Task
);

CREATE TABLE IF NOT EXISTS TaskBC (
	bc_id UUID NOT NULL REFERENCES BCInstance(bc_id) ON DELETE CASCADE
	,
	TaskID UUID NOT NULL, -- можно сделать это PRIMARY KEY и оставить внешние ключи
	Description TEXT NOT NULL, 
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	Deadline DATE NOT NULL,
	RecievedScore INT NOT NULL,
	PRIMARY KEY(bc_id, TaskID) --мб тоже см Task
);

-- views

drop view if exists labview;

CREATE VIEW labview AS
    SELECT LabNumber, LabName, LabText, MaxScore, MinScore, LabDate, Deadline, module_id,
		NumOfInstance, RecievedScore, Remarks, BonusScore
	FROM
    	Lab
	LEFT JOIN LabInstance 
   	ON LabID = lab_id;
SELECT * FROM labview;

drop view if exists rkview;

CREATE VIEW rkview AS
    SELECT BCNum, DateOdPassing, MaxScore, MinScore, module_id, NumOfInstance, Variant, RecievedScore, Remarks
	FROM
    	BC
	LEFT JOIN BCInstance 
   	ON BCID = bc_id;
SELECT * FROM rkview;

drop view if exists semview;
CREATE VIEW semview AS
    SELECT SemNumber, Theme, SemDate, module_id, WasAttended, BonusScore
	FROM
    	Seminar
	LEFT JOIN SeminarAttendance 
   	ON SeminarID = seminar_id;
SELECT * FROM semview;

drop view if exists lectview;
CREATE VIEW lectview AS
    SELECT LectNumber, Theme, LectDate, module_id, WasAttended, BonusScore
	FROM
    	Lecture
	LEFT JOIN LectureAttendance 
   	ON LectureID = lecture_id;
SELECT * FROM lectview;

drop view if exists cpview;
CREATE VIEW cpview AS
    SELECT Subject, Description, StartDate, Deadline, student_id, ProjAssignment, TitleOfProject, RecievedScore, DateOdPassing
	FROM
    	CourseProject
	LEFT JOIN StudentCourseProject 
   	ON ProjectID = project_id;
SELECT * FROM cpview;

drop view if exists StudentQueue
create view StudentQueue AS 
	SELECT student_id, queue_id, NumInQueue 
	FROM StudentInQueue


-- Sonya
INSERT INTO StudentGroup(GroupID, GroupName, YearOfAdmission, Course, AmountOfStudents) VALUES
	(gen_random_uuid(),'ИУ9-61Б', 2020, 3, 28),
	(gen_random_uuid(), 'ИУ9-62Б', 2020, 3, 24);

SELECT * FROM StudentGroup;

INSERT INTO Users(UserID, Login, Passw, UsersRights) VALUES
	(gen_random_uuid(),'Iliin', 'dskdfjkfdn',False),
    (gen_random_uuid(),'Kozoch', 'dskdjkfdn',False),
    (gen_random_uuid(),'SPKirich', 'dskdkfdn',False),
    (gen_random_uuid(),'Belyaev', 'skdfjkfdn',False),
    (gen_random_uuid(),'Yarov', 'dskdfjkfdn',False),
    (gen_random_uuid(),'IgorE', 'bdiscool',true),
    (gen_random_uuid(),'SoFa325', 'yyy',true),
    (gen_random_uuid(),'Anastasia', 'djkbfhkd',true),
    (gen_random_uuid(),'Tsarukan', 'fff',false),
	(gen_random_uuid(), 'Pos', 'kdjckdk', True);