CREATE DATABASE IF NOT EXISTS `schedule`;
USE `schedule`;

create table if not exists `schedule.professors`
(
  id         int auto_increment
  comment 'Уникальный идентификатор преподавателя'
    primary key,
  firstname  varchar(50) not null
  comment 'Имя преподавателя',
  surname    varchar(50) not null
  comment 'Фамилия преподавателя',
  patronymic varchar(50) null
  comment 'Отчество преподавателя',
  chair      varchar(15) not null
  comment 'Кафедра, на которой работает преподаватель',
  monday     json        null,
  tuesday    json        null,
  wednesday  json        null,
  thursday   json        null,
  friday     json        null,
  saturday   json        null
)
  comment 'Расписание преподавателей
Пример:
{
    "lessons": [
		   {"time": "8:30 - 10:05", "subject": {"numerator": "", "denominator": "", "is_differ": false}},
		   {"time": "10:15 - 11:50", "subject": {"numerator": "", "denominator": "", "is_differ": false}},
		   {"time": "12:00 - 13:35", "subject": {"numerator": "", "denominator": "", "is_differ": false}},
		   {"time": "13:50 - 15:25", "subject": {"numerator": "", "denominator": "", "is_differ": false}},
		   {"time": "15:40 - 17:15", "subject": {"numerator": "", "denominator": "", "is_differ": false}},
		   {"time": "17:25 - 19:00", "subject": {"numerator": "", "denominator": "", "is_differ": false}},
		   {"time": "19:10 - 20:45", "subject": {"numerator": "", "denominator": "", "is_differ": false}}],
    "is_empty": true
}';

create table if not exists `schedule.student_groups`
(
  id         int auto_increment
  comment 'Уникальный идентификатор группы'
    primary key,
  group_name varchar(15) not null
  comment 'Номер группы',
  monday     json        null,
  tuesday    json        null,
  wednesday  json        null,
  thursday   json        null,
  friday     json        null,
  saturday   json        null,
  constraint student_groups_group_name_uindex
  unique (group_name)
)
  comment 'Расписание групп
Пример:
{
    "lessons": [
		   {"time": "8:30 - 10:05", "subject": {"numerator": "", "denominator": "", "is_differ": false}},
		   {"time": "10:15 - 11:50", "subject": {"numerator": "", "denominator": "", "is_differ": false}},
		   {"time": "12:00 - 13:35", "subject": {"numerator": "", "denominator": "", "is_differ": false}},
		   {"time": "13:50 - 15:25", "subject": {"numerator": "", "denominator": "", "is_differ": false}},
		   {"time": "15:40 - 17:15", "subject": {"numerator": "", "denominator": "", "is_differ": false}},
		   {"time": "17:25 - 19:00", "subject": {"numerator": "", "denominator": "", "is_differ": false}},
		   {"time": "19:10 - 20:45", "subject": {"numerator": "", "denominator": "", "is_differ": false}}],
    "is_empty": true
}';

CREATE DATABASE IF NOT EXISTS `other_themes`;
USE `other_themes`;

create table if not exists jokes
(
  id    int                            not null
  comment 'Уникальный идентификатор шутки'
    primary key auto_increment,
  theme varchar(50) default 'no_theme' null
  comment 'Категория или тема анекдота',
  body  text                           not null
  comment 'Текст анекдота'
)
  comment 'Анекдоты'
