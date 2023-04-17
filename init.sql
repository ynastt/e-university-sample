DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Student;
DROP TABLE IF EXISTS StudentGroup;
DROP TABLE IF EXISTS Teacher;
DROP TABLE IF EXISTS Subject;
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

CREATE TABLE IF NOT EXISTS Student (
	StudentID serial PRIMARY KEY,
	StudentName VARCHAR(50) NOT NULL,
	Surname VARCHAR(50) NOT NULL,
	Patronymic VARCHAR(50),
	Email VARCHAR(255) UNIQUE NOT NULL,
	Phone VARCHAR(11),
	YearOfAdmission VARCHAR(4) NOT NULL,
	PassedCourses INT NOT NULL,
	NumInGroup serial
);

CREATE TABLE IF NOT EXISTS StudentGroup (
	GroupID serial PRIMARY KEY,
	YearOfAdmission VARCHAR(4) NOT NULL,
	Course INT NOT NULL,
	AmountOfStudents INT NOT NULL
);

CREATE TABLE IF NOT EXISTS Teacher (
	TeacherID serial PRIMARY KEY,
	TeacherName VARCHAR(50) NOT NULL,
	Surname VARCHAR(50) NOT NULL,
	Patronymic VARCHAR(50),
	Email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS Subject (
	SubjectID serial PRIMARY KEY,
	Description TEXT NOT NULL, -- какой тип данных?
	SubjectProgram JSON NOT NULL, -- какой тип данных?
	NumberOfHours INT NOT NULL,
	NumberOfCredits INT NOT NULL -- что это?
);

CREATE TABLE IF NOT EXISTS Exam (
	ExamID serial PRIMARY KEY, -- как идентифицировать экзамены?
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL 
);

CREATE TABLE IF NOT EXISTS Lecture (
	LectureID serial PRIMARY KEY, -- как идентифицировать?
	Theme TEXT NOT NULL,
	LectureText JSON NOT NULL -- какой тип данных?
);

CREATE TABLE IF NOT EXISTS Seminar (
	SeminarID serial PRIMARY KEY, -- как идентифицировать?
	Theme TEXT NOT NULL,
	SeminarText JSON NOT NULL -- какой тип данных?
);

CREATE TABLE IF NOT EXISTS Lab (
	LabID serial PRIMARY KEY, -- как идентифицировать?
	LabName TEXT NOT NULL,
	LabText JSON NOT NULL, -- какой тип данных?
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	LabDate date NOT NULL,
	Deadline date NOT NULL
);

CREATE TABLE IF NOT EXISTS BC (
	BCID serial PRIMARY KEY, -- как идентифицировать?
	Theme TEXT NOT NULL,
	Questions JSON NOT NULL, -- какой тип данных?
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL
);

CREATE TABLE IF NOT EXISTS Task (
	TaskID serial PRIMARY KEY, -- как идентифицировать?
	Description TEXT NOT NULL, -- какой тип данных?
	MaxScore INT NOT NULL,
	MinScore INT NOT NULL,
	Deadline date NOT NULL, --зачем?
	RecievedScore INT NOT NULL
);

CREATE TABLE IF NOT EXISTS CourseProject (
	ProjectID serial PRIMARY KEY, -- как идентифицировать?
	Subject VARCHAR(50) NOT NULL,
	Description TEXT NOT NULL, -- какой тип данных?
	NumberOfHours INT NOT NULL,
	StartDate date NOT NULL,
	Deadline date NOT NULL
);

CREATE TABLE IF NOT EXISTS Queue (
	QueueID serial PRIMARY KEY,
	StartDate date NOT NULL
);

INSERT INTO  Student(StudentName, Surname, Patronymic, Email, Phone, YearOfAdmission, PassedCourses, NumInGroup)
VALUES
	('Анастасия', 'Яровикова', 'Сергеевна', 'yarovnast@bmstu.ru', '89598675865', '2020', 101, DEFAULT);

SELECT * FROM Student;