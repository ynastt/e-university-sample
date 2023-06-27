

select * from StudentGroup;
select * from Users;
select * from Student where Email = 'yarovnast@mail.ru';
select * from Subject;
select * from Exam where subject_id = (SELECT SubjectID from Subject where Description = 'БАЗЫ ДАННЫХ');

select * from Modules where subject_id = (SELECT SubjectID from Subject where Description = 'БАЗЫ ДАННЫХ');

SELECT ModuleID from Modules where subject_id = (SELECT SubjectID from Subject where Description = 'БАЗЫ ДАННЫХ');
-- select * from Lab where module_id = (SELECT ModuleID from Modules where subject_id = (SELECT SubjectID from Subject where Description = 'БАЗЫ ДАННЫХ'))

-- select * from Lab where module_id = '84d9bb22-939a-469c-b27e-01ee3a43e63b';
-- select * from Lab where module_id = '04b03c25-59be-42f7-bb1f-4c452a36aa6e';
-- drop table lab cascade;
-- CREATE TABLE IF NOT EXISTS Lab (
-- 	LabID UUID PRIMARY KEY,
-- 	LabNumber INT NOT NULL,
-- 	LabName TEXT NOT NULL,
-- 	LabText TEXT NOT NULL, 
-- 	MaxScore INT NOT NULL,
-- 	MinScore INT NOT NULL,
-- 	LabDate date NOT NULL,
-- 	Deadline date NOT NULL,
-- 	CONSTRAINT lab_unique UNIQUE(LabName),
-- 	module_id UUID NOT NULL REFERENCES Modules(ModuleID) ON DELETE RESTRICT --запрет на удаление модуля через таблицу лабы
-- );

-- labs
delete from Lab;

INSERT INTO Lab(LabID, LabNumber, LabName, LabText, MaxScore, MinScore, LabDate, Deadline, module_id) Values 
	 (gen_random_uuid(), 1, 'ER', 'Моделирование данных с использованием модели сущность связь.', 
      8, 5, '2022-09-09', '2022-09-16', '84d9bb22-939a-469c-b27e-01ee3a43e63b'),
      (gen_random_uuid(), 2, 'SOM', 'Моделирование данных с использованием модели семантических объектов.', 
      5, 4, '2022-09-16', '2022-09-23', '84d9bb22-939a-469c-b27e-01ee3a43e63b'),
      (gen_random_uuid(), 3, 'ER->R', 'Преобразование модели "сущность-связь" в реляционную модель.', 
      7, 5, '2022-09-23', '2022-09-30', '84d9bb22-939a-469c-b27e-01ee3a43e63b'),
      (gen_random_uuid(), 4, 'SOM->R', 'Преобразование модели семантических объектов в реляционную модель.',  
      5, 4, '2022-09-30', '2022-10-07', '84d9bb22-939a-469c-b27e-01ee3a43e63b');
INSERT INTO Lab(LabID, LabNumber, LabName, LabText, MaxScore, MinScore, LabDate, Deadline, module_id) Values 	  
	(gen_random_uuid(), 5, 'DB files', 'Операции с базой данных, файлами, схемами.', 3, 2, '2022-10-07', '2022-10-14', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
    (gen_random_uuid(), 6, 'Constraints', 'Ключи, ограничения, значения по умолчанию.', 3, 2, '2022-10-14', '2022-10-21', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
    (gen_random_uuid(), 7, 'Views', 'Представления и индексы.',3, 2, '2022-10-21', '2022-10-28', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
    (gen_random_uuid(), 8, 'SP/Cursor/F', 'Хранимые процедуры, курсоры, пользовательские функции.',4, 3, '2022-10-28', '2022-11-11', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
    (gen_random_uuid(), 9, 'Trigger', 'Триггеры DML.', 4, 3, '2022-10-28', '2022-11-11', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
    (gen_random_uuid(), 10, 'Transactions', 'Режимы выполнения транзакций.', 4, 2, '2022-11-11', '2022-11-18', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
    (gen_random_uuid(), 11, 'DB App', 'Создание базы данных и запросов к ней в СУБД SQL Server 2012.',  7, 5, '2022-11-18', '2022-11-25', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
    (gen_random_uuid(), 12, 'ADO.NET', 'Программирование клиентского приложения доступа к данным.',4, 2, '2022-11-25', '2022-12-02', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
    (gen_random_uuid(), 13, 'horizontal', 'Создание распределенных баз данных на основе секционированных представлений.', 3, 2, '2022-12-02', '2022-12-09', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
    (gen_random_uuid(), 14, 'vertical', 'Создание вертикально фрагментированных таблиц средствами СУБД SQL Server 2012.',3, 2, '2022-12-09', '2022-12-16', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
    (gen_random_uuid(), 15, 'linked', 'Создание распределенных баз данных со связанными таблицами средствами СУБД SQL Server 2012.',3, 2, '2022-12-16', '2022-12-23', '04b03c25-59be-42f7-bb1f-4c452a36aa6e');
select * from Lab;
delete from labinstance;

insert into LabInstance values
	((select STudentID from Student where Email = 'yarovnast@mail.ru'), (SELECT LabID from Lab where Labname = 'ER'), '2022-09-16', 1, 8, NULL, 'БД отлично спроектирована!', 1),
	((select STudentID from Student where Email = 'yarovnast@mail.ru'), (SELECT LabID from Lab where Labname = 'Views'), '2022-09-23', 2, 2, NULL, 'ЛР сдана позже дедлайна', 0),
	((select STudentID from Student where Email = 'yarovnast@mail.ru'), (SELECT LabID from Lab where Labname = 'Trigger'), '2022-10-28', 1, 3, NULL, '-', 0);
select * from LabInstance;

drop view if exists labview;


select * from lab
select * from LabInstance;

SELECT LabNumber, LabName, LabText, MaxScore, MinScore, LabDate, Deadline, module_id,
		NumOfInstance, RecievedScore, Remarks, BonusScore
FROM
    Lab
LEFT JOIN LabInstance 
   ON LabID = lab_id
   where module_id = '04b03c25-59be-42f7-bb1f-4c452a36aa6e';
   
CREATE VIEW labview AS
    SELECT LabNumber, LabName, LabText, MaxScore, MinScore, LabDate, Deadline, module_id,
		NumOfInstance, RecievedScore, Remarks, BonusScore
	FROM
    	Lab
	LEFT JOIN LabInstance 
   	ON LabID = lab_id;
SELECT * FROM labview;

-- rks
delete from BC where module_id = '84d9bb22-939a-469c-b27e-01ee3a43e63b';
delete from BC where module_id = '04b03c25-59be-42f7-bb1f-4c452a36aa6e';

alter table BC add column BCNum int;
add constraint rk_unique UNIQUE(Theme);
INSERT INTO BC(BCID, Theme, Questions, MaxScore, MinScore, module_id) Values 
	 (gen_random_uuid(), 'Проектирование баз данных. Модели "сущность-связь", семантических объектов, реляционные', 'Список вопросов для подготовки к РК в письме на почте',  
      10, 5, '84d9bb22-939a-469c-b27e-01ee3a43e63b'),
	  (gen_random_uuid(), 'СУБД SQL Server 2012', 'Список вопросов для подготовки к РК в письме на почте',  
      10, 5, '04b03c25-59be-42f7-bb1f-4c452a36aa6e');
update BC  set BCNum = 1 where BCID ='05355ab8-77e4-4641-b86d-39d58a057af8';
update BC  set BCNum = 2 where BCID ='8f962858-94b8-4b0c-9c00-d3841a3fa6b8';	  
select * from BC;

alter table BCInstance add column Remarks TEXT;	
insert into BCInstance(student_id, bc_id, DateOdPassing, NumOfInstance, RecievedScore, Variant, Remarks) values
-- 	((select StudentID from Student where Email = 'yarovnast@mail.ru'), (SELECT BCID from BC where Theme = 'Проектирование баз данных. Модели "сущность-связь", семантических объектов, реляционные'), '2022-10-10', 1, 8, 24, NULL),
	((select StudentID from Student where Email = 'yarovnast@mail.ru'), (SELECT BCID from BC where Theme = 'СУБД SQL Server 2012'), '2022-12-23', 1, 5, 29, NULL);
select * from BCInstance;

drop view if exists rkview;
CREATE VIEW rkview AS
    SELECT BCNum, DateOdPassing, MaxScore, MinScore, module_id, NumOfInstance, Variant, RecievedScore, Remarks
	FROM
    	BC
	LEFT JOIN BCInstance 
   	ON BCID = bc_id;
SELECT * FROM rkview;

-- sems
-- alter table Seminar add column SemNumber int;
-- alter table Seminar add column SemDate date;

select * from Seminar;
select * from SeminarAttendance;

					
drop view if exists semview;
CREATE VIEW semview AS
    SELECT SemNumber, Theme, SemDate, module_id, WasAttended, BonusScore
	FROM
    	Seminar
	LEFT JOIN SeminarAttendance 
   	ON SeminarID = seminar_id;
SELECT * FROM semview;
--lect
-- alter table Lecture add column LectNumber int;
alter table Lecture add column LectDate date;
-- select * from Lab where module_id = '84d9bb22-939a-469c-b27e-01ee3a43e63b';
-- select * from Lab where module_id = '04b03c25-59be-42f7-bb1f-4c452a36aa6e';

-- alter table Seminar
-- add constraint sem_unique UNIQUE(Theme);

INSERT INTO Lecture(LectureID, Theme, LectureText, module_id, LectNumber, LectDate) Values 
	 (gen_random_uuid(), 'Введение', 'https://disk.yandex.ru/d/UOifGiyLv5K7w', '84d9bb22-939a-469c-b27e-01ee3a43e63b', 1, '2022-09-10'),
     (gen_random_uuid(), 'Модель "сущность-связь"', '-','84d9bb22-939a-469c-b27e-01ee3a43e63b', 3 ,'2022-10-10'),
	  (gen_random_uuid(), 'СУБД SQL Server 2012 – Представления.Индексы.', 'https://disk.yandex.ru/d/UOifGiyLv5K7w', '04b03c25-59be-42f7-bb1f-4c452a36aa6e', 7, '2022-11-20'),
      (gen_random_uuid(), 'СУБД SQL Server 2012 – Создание распределенных баз данных.', 'https://disk.yandex.ru/d/UOifGiyLv5K7w', '04b03c25-59be-42f7-bb1f-4c452a36aa6e', 14, '2022-12-15');

INSERT INTO LectureAttendance(student_id, lecture_id, WasAttended, BonusScore) Values 
	 ((select StudentID from Student where Email = 'yarovnast@mail.ru'), (select LectureID from Lecture where Theme = 'Введение'), TRUE, 0),
	 ((select StudentID from Student where Email = 'yarovnast@mail.ru'), (select LectureID from Lecture where Theme = 'Модель "сущность-связь"'), TRUE, 1),
	 ((select StudentID from Student where Email = 'yarovnast@mail.ru'), (select LectureID from Lecture where Theme = 'СУБД SQL Server 2012 – Представления.Индексы.'), TRUE, 0),
	 ((select StudentID from Student where Email = 'yarovnast@mail.ru'), (select LectureID from Lecture where Theme = 'СУБД SQL Server 2012 – Создание распределенных баз данных.'), TRUE, 0);


CREATE TABLE IF NOT EXISTS Lecture (
	LectureID UUID PRIMARY KEY,
	Theme TEXT NOT NULL UNIQUE,
	LectureText TEXT NOT NULL, 
	module_id UUID NOT NULL REFERENCES Modules(ModuleID) ON DELETE RESTRICT,
	LectNumber INT NOT NULL,
	LectDate date
);

drop view if exists lectview;
CREATE VIEW lectview AS
    SELECT LectNumber, Theme, LectDate, module_id, WasAttended, BonusScore
	FROM
    	Lecture
	LEFT JOIN LectureAttendance 
   	ON LectureID = lecture_id;
SELECT * FROM lectview;

select * from Lecture;
select * from LectureAttendance;
--courseproj

CREATE TABLE IF NOT EXISTS CourseProject (
	ProjectID UUID PRIMARY KEY, 
	Subject VARCHAR(50) NOT NULL,
	Description TEXT NOT NULL, 
	NumberOfHours INT NOT NULL,
	StartDate DATE NOT NULL,
	Deadline DATE NOT NULL
);
insert into CourseProject(ProjectID, Subject, Description, NumberOfHours, StartDate, Deadline) values 
	(gen_random_uuid(), 'Алгоритмы компьютерной графики', 'Курсовая работа', 180, '2022-09-10', '2023-01-24');

select * from CourseProject;

insert into StudentCourseProject(student_id, project_id, ProjAssignment, TitleOfProject, RecievedScore, DateOdPassing) values 
	((select StudentID from Student where Email = 'yarovnast@mail.ru'), '424b9848-ea66-4762-ae6a-d48ede9a7388', 'разработка приложения для визуализации блокчейна Биткоин', 'Приложение визуализации графа Биткоин', 5, '2023-01-24');

select * from StudentCourseProject;

drop view if exists cpview;
CREATE VIEW cpview AS
    SELECT Subject, Description, StartDate, Deadline, student_id, ProjAssignment, TitleOfProject, RecievedScore, DateOdPassing
	FROM
    	CourseProject
	LEFT JOIN StudentCourseProject 
   	ON ProjectID = project_id;
SELECT * FROM cpview;