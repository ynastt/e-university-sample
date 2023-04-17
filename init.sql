DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS StudentGroup;
DROP TABLE IF EXISTS Student;
DROP TABLE IF EXISTS Teacher;
DROP TABLE IF EXISTS Subject;
DROP TABLE IF EXISTS Modules;
DROP TABLE IF EXISTS Exam;
DROP TABLE IF EXISTS Lecture;
DROP TABLE IF EXISTS Seminar;
DROP TABLE IF EXISTS Lab;
DROP TABLE IF EXISTS BC;
DROP TABLE IF EXISTS Task;
DROP TABLE IF EXISTS CourseProject;
DROP TABLE IF EXISTS Queue;

CREATE TABLE IF NOT EXISTS Users (
	UserID serial PRIMARY KEY,
	Login VARCHAR(50) NOT NULL,
	Passw VARCHAR(50) NOT NULL,
	UsersRights Boolean NOT NULL
);

CREATE TABLE IF NOT EXISTS StudentGroup (
	GroupID serial PRIMARY KEY,
	YearOfAdmission VARCHAR(4) NOT NULL,
	Course INT NOT NULL,
	AmountOfStudents INT NOT NULL
);

CREATE TABLE IF NOT EXISTS Student (
	StudentID serial PRIMARY KEY,
	StudentName VARCHAR(50) NOT NULL,
	Surname VARCHAR(50) NOT NULL,
	Patronymic VARCHAR(50),
	Email VARCHAR(255) UNIQUE NOT NULL,
	Phone VARCHAR(11),
	YearOfAdmission VARCHAR(4) NOT NULL,
	PassedCourses INT NOT NULL,
	NumInGroup serial,
	FOREIGN KEY (user_id) REFERENCES Users (UserID),
	FOREIGN KEY (group_id) REFERENCES StudentGroup (GroupID)
);

CREATE TABLE IF NOT EXISTS Teacher (
	TeacherID serial PRIMARY KEY,
	TeacherName VARCHAR(50) NOT NULL,
	Surname VARCHAR(50) NOT NULL,
	Patronymic VARCHAR(50),
	Email VARCHAR(255) UNIQUE NOT NULL,
	FOREIGN KEY (user_id) REFERENCES Users (UserID)
);

CREATE TABLE IF NOT EXISTS Subject (
	SubjectID serial PRIMARY KEY,
	Description TEXT NOT NULL, -- какой тип данных?
	SubjectProgram JSON NOT NULL, -- какой тип данных?
	NumberOfHours INT NOT NULL,
	NumberOfCredits INT NOT NULL -- что это?
);

CREATE TABLE IF NOT EXISTS Modules (
	ModuleID serial,
	SubjectID serial,
	PRIMARY KEY(ModuleID, SubjectID)
	ModuleName Varchar(50) NOT NULL, 
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL
);

CREATE TABLE IF NOT EXISTS Exam (
	ExamID serial PRIMARY KEY, 
	Questions JSON NOT NULL, -- какой тип данных?
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	ExamDate date,
	FOREIGN KEY (subject_id) REFERENCES Subject (SubjectID)
);

CREATE TABLE IF NOT EXISTS Lecture (
	LectureID serial PRIMARY KEY, -- как идентифицировать?
	Theme TEXT NOT NULL,
	LectureText JSON NOT NULL, -- какой тип данных?
	FOREIGN KEY (module_id) REFERENCES Modules (ModuleID)
);

CREATE TABLE IF NOT EXISTS Seminar (
	SeminarID serial PRIMARY KEY, -- как идентифицировать?
	Theme TEXT NOT NULL,
	SeminarText JSON NOT NULL, -- какой тип данных?
	FOREIGN KEY (module_id) REFERENCES Modules (ModuleID)
);

CREATE TABLE IF NOT EXISTS Lab (
	LabID serial PRIMARY KEY, -- как идентифицировать?
	LabName TEXT NOT NULL,
	LabText JSON NOT NULL, -- какой тип данных?
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	LabDate date NOT NULL,
	Deadline date NOT NULL,
	FOREIGN KEY (module_id) REFERENCES Modules (ModuleID)
);

CREATE TABLE IF NOT EXISTS BC (
	BCID serial PRIMARY KEY, -- как идентифицировать?
	Theme TEXT NOT NULL,
	Questions JSON NOT NULL, -- какой тип данных?
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	FOREIGN KEY (module_id) REFERENCES Modules (ModuleID)
);

CREATE TABLE IF NOT EXISTS Task (
	TaskID serial PRIMARY KEY, -- как идентифицировать?
	Description TEXT NOT NULL, -- какой тип данных?
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	Deadline DATE NOT NULL, --зачем?
	RecievedScore INT NOT NULL
);

CREATE TABLE IF NOT EXISTS CourseProject (
	ProjectID serial PRIMARY KEY, -- как идентифицировать?
	Subject VARCHAR(50) NOT NULL,
	Description TEXT NOT NULL, -- какой тип данных?
	NumberOfHours INT NOT NULL,
	StartDate DATE NOT NULL,
	Deadline DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS Queue (
	QueueID serial PRIMARY KEY,
	StartDate DATE NOT NULL
);

-- dop
CREATE TABLE IF NOT EXISTS StudentInQueue (
	FOREIGN KEY (student_id) REFERENCES Student (StudentID),
	FOREIGN KEY (queue_id) REFERENCES Queue (QueueID),
	PRIMARY KEY(student_id, queue_id),
	NumInQueue Int NOT NULL,
	Task VARCHAR(50) NOT NULL, -- решить что с типом
);

CREATE TABLE IF NOT EXISTS TeacherSubject (
	FOREIGN KEY (teacher_id) REFERENCES Teacher (TeacherID),
	FOREIGN KEY (subject_id) REFERENCES Subject (SubjectID),
	PRIMARY KEY(teacher_id, subject_id),
	TeacherRole Int NOT NULL
);

CREATE TABLE IF NOT EXISTS Supervisor (
	FOREIGN KEY (teacher_id) REFERENCES Teacher (TeacherID),
	FOREIGN KEY (project_id) REFERENCES CourseProject (ProjectID),
	PRIMARY KEY(teacher_id, project_id),
	SupervisorRole Int NOT NULL
);

CREATE TABLE IF NOT EXISTS StudentCourseProject (
	FOREIGN KEY (student_id) REFERENCES Student (StudentID),
	FOREIGN KEY (project_id) REFERENCES CourseProject (ProjectID),
	PRIMARY KEY(student_id, project_id),
	ProjAssignment TEXT NOT NULL,
	TitleOfProject VARCHAR(100) NOT NULL,
	RecievedScore INT NOT NULL,
	DateOdPassing DATE
);

CREATE TABLE IF NOT EXISTS LectureAttendance (
	FOREIGN KEY (student_id) REFERENCES Student (StudentID),
	FOREIGN KEY (lecture_id) REFERENCES Lecture (LectureID),
	PRIMARY KEY(student_id, lecture_id),
	WasAttended BOOL NOT NULL,
	BonusScore INT NOT NULL
);

CREATE TABLE IF NOT EXISTS SeminarAttendance (
	FOREIGN KEY (student_id) REFERENCES Student (StudentID),
	FOREIGN KEY (seminar_id) REFERENCES Seminar (SeminarID),
	PRIMARY KEY(student_id, seminar_id),
	WasAttended BOOL NOT NULL,
	BonusScore INT NOT NULL
);

CREATE TABLE IF NOT EXISTS LabInstance (
	FOREIGN KEY (student_id) REFERENCES Student (StudentID),
	FOREIGN KEY (lab_id) REFERENCES Lab (LabID),
	PRIMARY KEY(student_id, lab_id),
	NumOfInstance INT NOT NULL,
	RecievedScore INT NOT NULL,
	Variant INT,
	DateOdPassing DATE NOT NULL,
	Remarks TEXT,
	BonusScore INT NOT NULL
);

CREATE TABLE IF NOT EXISTS BCInstance (
	FOREIGN KEY (student_id) REFERENCES Student (StudentID),
	FOREIGN KEY (bc_id) REFERENCES BC (BCID),
	PRIMARY KEY(student_id, bc_id),
	NumOfInstance INT NOT NULL,
	RecievedScore INT NOT NULL,
	Variant INT,
	DateOdPassing DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS ExamInstance (
	FOREIGN KEY (student_id) REFERENCES Student (StudentID),
	FOREIGN KEY (exam_id) REFERENCES Exam (ExamID),
	PRIMARY KEY(student_id, bc_id),
	NumOfInstance INT NOT NULL,
	RecievedScore INT NOT NULL,
	TicketNumber INT,
	DateOdPassing DATE NOT NULL
);

INSERT INTO  Student(StudentName, Surname, Patronymic, Email, Phone, YearOfAdmission, PassedCourses, NumInGroup)
VALUES
	('Анастасия', 'Яровикова', 'Сергеевна', 'yarovnast@bmstu.ru', '89598675865', '2020', 101, DEFAULT);

SELECT * FROM Student;