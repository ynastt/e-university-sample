import os
import telebot
from telebot import types
import psycopg2
from configparser import ConfigParser
import datetime
import data_storage as db
# using now() to get current time
current_time = datetime.datetime.now()
 

first = 0
API_TOKEN = '6042118462:AAFSsuIPOcaRocGuT9ImTnuaR8yXgy-_7e0'
bot = telebot.TeleBot (API_TOKEN)


def config(filename='../configs/database.ini', section='postgresql'):
    parser = ConfigParser()
    parser.read(filename)

    db = {}
    if parser.has_section(section):
        params = parser.items(section)
        for param in params:
            db[param[0]] = param[1]
    else:
        raise Exception('Section {0} not found in the {1} file'.format(section, filename))

    return db

def connectDB():
    """ Connect to the PostgreSQL database server """
    conn = None
    try:
        params = config()
        print('Connecting to the PostgreSQL database...')
        conn = psycopg2.connect(**params)
        # create a cursor
        cur = conn.cursor()
    except (Exception, psycopg2.DatabaseError) as error:
        print(error)
    return conn, cur
    

@bot.message_handler(commands=['start'])
def start(message):
    markup = types.ReplyKeyboardMarkup(resize_keyboard=True)
    btn1 = types.KeyboardButton("👋 Привет, запиши меня в очередь")
    markup.add(btn1)
    bot.send_message(message.from_user.id, "👋 - Привет, Я - бот-помощник для организации очереди на сдачу лабораторных работ! Для записи в очередь нажми на кнопку ниже", reply_markup=markup)

@bot.message_handler(commands=['help'])
def start(message):
    bot.send_message(message.from_user.id, "Этот бот создан того, чтобы студенты могли организовать очередь на сдачу лабораторных работ!\nДля запуска напиши сообщение /start")

@bot.message_handler(content_types=['text'])
def get_text_messages(message):
    if message.text == '👋 Привет, запиши меня в очередь':
        markup = types.ReplyKeyboardMarkup(resize_keyboard=True)
        btn1 = types.KeyboardButton('Войти в электронный дневник')
        markup.add(btn1)
        bot.send_message(message.from_user.id, '❗Для записи в очередь необходимо авторизоваться в электронном дневнике студента', reply_markup=markup)
    elif message.text == 'Войти в электронный дневник':
        bot.send_message(message.from_user.id, 'Введите логин:')
        conn, cur = connectDB()
        # print('PostgreSQL user table:')
        cur.execute('SELECT * from Users')
        usrs = cur.fetchall()
        for user in usrs:
            db.usersset.add(user[1])
        # print('PostgreSQL subject table:')
        cur.execute('SELECT SubjectID, Description from Subject')
        subjects = cur.fetchall()
        for s in subjects:
            # print(s)
            db.subj_data[s[1]] = s[0] 
        cur.close()
        conn.close()
    elif message.text in db.usersset and not db.auth:
        db.login = message.text
        print(message.text)
        print(message.text in db.usersset)
        db.auth = True
        markup = types.ReplyKeyboardMarkup(resize_keyboard=True)
        for s in db.subj_data.keys():
            print(s)
            btn1 = types.KeyboardButton(s)
            markup.add(btn1)
        bot.send_message(message.from_user.id, 'Авторизация прошла успешно! Выберете предмет, по которому Вы сдаете работу', reply_markup=markup)    
    elif message.text not in db.usersset and not db.auth:
        print(message.text in db.usersset)
        print(db.usersset)
        bot.send_message(message.from_user.id, 'Вы не зарегистрированы в электронном дневнике.\nК сожалению, вы не можете быть записаны в очередь.')
         
    elif message.text in db.subj_data.keys() and not db.subj_choice:
        db.subj_choice = True
        db.current_subject = db.Subject(db.subj_data[message.text], message.text)
        print(db.current_subject.name, db.current_subject.id)
        conn, cur = connectDB()
        print('PostgreSQL queue table for current subject:')
        
        cur.execute('SELECT * from Queue where subject_id = %s',
                     (db.current_subject.id,))
        active_queues = cur.fetchall()
        print(len(active_queues))
        if len(active_queues) == 0:
            markup = types.ReplyKeyboardMarkup(resize_keyboard=True)
            btn1 = types.KeyboardButton('Создать очередь')
            markup.add(btn1)
            bot.send_message(message.from_user.id, 'Сейчас нет активных очередей по данному предмету.\nХотите создать новую очередь?', reply_markup=markup)
        else:
            markup = types.ReplyKeyboardMarkup(resize_keyboard=True)
            for q in active_queues:
                print(q)
                name = 'Очередь {date}'.format(date=q[1][0:10])
                btn1 = types.KeyboardButton(name)
                markup.add(btn1)
            bot.send_message(message.from_user.id, 'Выберите одну из активных очередей', reply_markup=markup)
        
        # cur.execute('SELECT SubjectID, Description from Subject')
        # subjects = cur.fetchall()
        # for s in subjects:
        #     # print(s)\ 
        #     db.subj_data[s[1]] = s[0] 
        # cur.close()
        # conn.close()
        bot.send_message(message.from_user.id, 'aaaa')
    
    elif message.text == 'Посмотреть очередь':
        bot.send_message(message.from_user.id, 'Подробно про советы по оформлению публикаций прочитать по ' + '[ссылке](https://habr.com/ru/docs/companies/design/)', parse_mode='Markdown')
    elif message.text == 'Удалить запись':
        bot.send_message(message.from_user.id, 'Подробно про советы по оформлению публикаций прочитать по ' + '[ссылке](https://habr.com/ru/docs/companies/design/)', parse_mode='Markdown')


if __name__ == '__main__':
    
    # Printing value of now.
    print("Time now at greenwich meridian is:", current_time)
    bot.polling(none_stop=True, interval=0) #обязательная для работы бота часть
