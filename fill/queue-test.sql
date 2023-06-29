insert into Users(userid, login, passw, usersrights) values
	(gen_random_uuid(), 'vscie', 'bdteach', 1),
	(gen_random_uuid(), 'yogi', 'aaa', 2),
	(gen_random_uuid(), 'belezza', 'sduchwe', 2);

insert into Users values (gen_random_uuid(),'SoFa325', 'yyy', 2)
select * from users

select * from StudentInQueue;

select * from StudentInQueue where queue_id = 'bdaf68b7-435f-4a2f-93fa-1be0fc2ff9a3'
SELECT COUNT(*) + 1 from StudentInQueue
 

drop table Queue cascade;

CREATE TABLE IF NOT EXISTS Queue (
	QueueID UUID PRIMARY KEY,
	StartDate timestamp with time zone NOT NULL,
	subject_id UUID NOT NULL REFERENCES Subject(SubjectID) ON DELETE CASCADE
);

insert into Queue(QueueID, StartDate, subject_id) values
 	(gen_random_uuid(), '2023-06-27', (SELECT SubjectID from Subject where Description = 'БАЗЫ ДАННЫХ'));

select * from StudentGroup

Insert into StudentGroup(GroupID, GroupName, YearOfAdmission, course, amountofstudents) values
	(gen_random_uuid(), 'ИУ9-62Б', 2020, 3, 24)

Insert into Student(StudentID, StudentName, Surname, Patronymic, Email, Phone, YearOfAdmission, PassedCourses, NumInGroup, user_id, group_id) values
	(gen_random_uuid(), 'Софья', 'Белякова', 'Сергеевна', 'sofabel@gmail.com', '89345678911', 2020, 3, 1, 'aed02e6d-8899-4ff8-ad88-5751df5493d0', 'f4e35b6c-156b-4ab4-b4ef-2b00468b0975'),
	(gen_random_uuid(), 'Александр', 'Федоров', 'Николаевич', 'fdx@gmail.com', '89344328126', 2020, 3, 27, '65860bd9-cc4d-4c0a-a5cd-99f59f078d9a', '52fc2308-aa38-4f65-9080-e96774be65dd'),
	(gen_random_uuid(), 'Дарья', 'Емельяненко', 'Сергеевна', 'belez@gmail.com', '89572328135', 2020, 3, 27, '463f0885-6525-41d9-992c-89c92739a96a', '52fc2308-aa38-4f65-9080-e96774be65dd');

select * from Student