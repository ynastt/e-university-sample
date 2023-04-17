INSERT into Users Values(1, 'SoFa325', 'yyy',True),
(2, 'Anastasia', 'djkbfhkd', True),
(3, 'IgorE', 'bdiscool', True),
(4, 'Tsarukan', 'fff', False);

INSERT into StudentGroup(Group_name,YearOfAdmission, Course, AmountOfStudents ) Values('ISC9-62B',2020, 3, 24),
                                ('ISC9-22B', 2022, 1, 31),
                                ('ISC9-22M', 2018, 5, 10),
                                ('ISC7-61B',2020, 3,  26),
                                ('ISC9-61B',2020, 3, 26 );
Insert into Student(user_id, group_id,StudentName,Surname, Email, YearOfAdmission, PassedCourses, NumInGroup)
            Values(1, 1,'Sofa', 'Bel', 'sofyabel@inbox.ru', 2020, 0000000, 2 );
	
Insert into Teacher(user_id, TeacherName, Surname, Email) Values (3, 'Igor', 'Vishnyakov', 'vcsi@mail.com');

Insert into Subject(Description, NumberOfHours, NumberOfCredits) Values
('Database learning course', 72, 42),
('Discreth Math', 128, 42),
('Object-oriented languages', 80, 60);

                    

--SELECT * FROM Users;
--SELECT * FROM StudentGroup;
--SELECT * FROM Student;
--SELECT * From Teacher;
--Select * from Subject;