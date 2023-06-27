from enum import Enum

usersset = set()
subj_data = {}
auth = False
subj_choice = False
login = ''

class Subject:
    def __init__(self, id, name):
        self.id = id
        self.name = name

class Queue:
    def __init__(self, id, date, subject_id):
        self.id = id
        self.date = date
        self.subject_id =  subject_id

class Task(Enum):
    LAB = 1
    RK = 2

class Queue:
    def __init__(self, student_id, queue_id, num, task):
        self.student_id = student_id
        self.queue_id = queue_id
        self.num = num
        self.task =  task