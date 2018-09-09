# МГТУ имени Н. Э. Баумана

## API For BMSTU Info Database

### Исполнители:

- Кочетков Михаил Дмитриевич
- Сукиасян Владимир Мартунович

### Руководители:

- Терехов Валерий Игоревич
- Иван

### Описание:

 Api для получение данных из информационной базы данных МГТУ имени Н. Э. Баумана

### Доступные функции:

- Расписание групп(на 6 рабочих дней)
- Расписание преподавателей(на 6 рабочих дней)

### Дополнительные функции:

- Анекдоты

### Документация:

- "/api/v1/schedule/info/professors" -
  - GET - получение списка преподавателей без расписания (имя,
    фамилия, отчество, кафедра)
- "/api/v1/schedule/professors" -
  - GET - получение всего списка расписаний преподавателей
  - POST - добавление нового расписания преподавателя
  - days - список дней, разделенные запятыми, которые должны быть в
    расписании.
     В структуре JSON дни, не входящие в этот список, имеют в поле
    "lessons" значение null
     Пример:
    /api/v1/schedule/professors?days=monday,wednesday
  - limit - количество записей
    Пример:
  /api/v1/schedule/professors?limit=1
  - offset - сколько записей от начала списка нужно пропустить.
    Только, если установлен параметр limit
    Пример:
  /api/v1/schedule/professors?limit=1;offset=1
- "/api/v1/schedule/professors/{surname}" -
  - GET - получение списка расписаний преподавателей с фамилий
    surname
  - days - список дней, разделенные запятыми, которые должны быть в
    расписании.
     В структуре JSON дни, не входящие в этот список, имеют в поле
    "lessons" значение null
     Пример:
    /api/v1/schedule/professors/Козлов?days=monday,wednesday
  - limit - количество записей
    Пример:
  /api/v1/schedule/professors/Козлов?limit=1
  - offset - сколько записей от начала списка нужно пропустить.
    Только, если установлен параметр limit
    Пример:
  /api/v1/schedule/professors/Козлов?limit=1;offset=1
- "/api/v1/schedule/professors/{id}" -
  - PUT - замена уже существующего расписания преподавателя с
    идентификатором id
  - DELETE - удаление расписания преподавателя с идентификатором id
- "/api/v1/schedule/info/student_groups" -
  - GET - получение списка групп без расписания (номер группы)
- "/api/v1/schedule/student_groups" -
  - GET - получение всего списка расписаний групп
  - POST - добавление нового расписания группы
  - days - список дней, разделенные запятыми, которые должны быть в
    расписании.
     В структуре JSON дни, не входящие в этот список, имеют в поле
    "lessons" значение null
     Пример:
    /api/v1/schedule/student_groups?days=monday,wednesday
  - limit - количество записей
    Пример:
  /api/v1/schedule/student_groups?limit=1
  - offset - сколько записей от начала списка нужно пропустить.
    Только, если установлен параметр limit
    Пример:
  /api/v1/schedule/student_groups?limit=1;offset=1
- "/api/v1/schedule/student_groups/{group_name}" -
  - GET - получение расписание группы group_name
  - PUT - замена уже существующего расписания группы group_name
  - DELETE - удаление расписания группы group_name
  - days - список дней, разделенные запятыми, которые должны быть в
    расписании.
     В структуре JSON дни, не входящие в этот список, имеют в поле
    "lessons" значение null
     Пример:
    /api/v1/schedule/student_groups/ИУ5-31Б?days=monday,wednesday
- "/api/v1/other_themes/info/jokes" -
  - GET - получение списка анекдотов без текста (название категории)
- "/api/v1/other_themes/jokes" -
  - GET - получение всего списка анекдотов
  - POST - добавление нового анекдота
  - limit - количество записей
    Пример:
  /api/v1/other_themes/jokes?limit=1
  - offset - сколько записей от начала списка нужно пропустить.
    Только, если установлен параметр limit
    Пример:
  /api/v1/other_themes/jokes?limit=1;offset=1
- "/api/v1/other_themes/jokes/{theme}" -
  - GET - получение списка анекдотов по заданной категории
  - limit - количество записей
    Пример:
  /api/v1/other_themes/jokes/robots?limit=1
  - offset - сколько записей от начала списка нужно пропустить.
    Только, если установлен параметр limit
    Пример:
  /api/v1/other_themes/jokes/robots?limit=1;offset=1
- "/api/v1/other_themes/jokes/{id}" -
  - PUT - замена уже существующего анекдота по его id
  - DELETE - удаление уже существующего анекдота по его id

Структура JSON файла расписания преподавателя

Поле id присутствует только в ответе на запрос пользователя

        {
            "id": __id__(int),
            "firstname": __name__(string),
            "surname": __surname__(string),
            "patronymic": __patronymic__(string),
            "chair": __chair__(string),
            "week": __week_object__(object)
        }

Структура JSON файла расписания групп

Поле id присутствует только в ответе на запрос пользователя

        {
            "id": __id__(int),
            "group_name": __group_name__(string),
            "week": __week_object__(object)
        }

Структура JSON объекта week

        {
            "monday": __day_object__(object),
            "tuesday": __day_object__(object),
            "wednesday": __day_object__(object),
            "thursday": __day_object__(object),
            "friday": __day_object__(object),
            "saturday": __day_object__(object)
        }

Структура JSON объекта day

Поле is_empty означает, что расписание пустое или нет

        {
        "lessons": __lesson_object__(object),
        "is_empty": __is_empty__(bool)
        }

Структура JSON объекта lesson

**Поле time содержит временные рамки урока в виде начало - конец.

Пример: "8:30 - 10:05"**

        {
            "time": __time__(string),
            "subject": __subject_object__(object)
        }

Структура JSON объекта lesson

**Поле numerator содержит название предмета и его кабинета по

числителям. Пример: "АСОИУ,501ю"**

 **Поле denominator содержит название предмета и его кабинета по

знаметелям. Пример: "Электротехника,502ю"**

 **Поле is_differ содержит информацию - отличаются ли предметы или

кабинеты в числитель и знаменатель**

        {
            "numerator": __ numerator__(string),
            "denominator": __denominator__(string),
            "is_differ": __is_differ__(bool)
        }

Структура JSON объекта joke

        {
            "id": __id__(int),
            "theme": __theme__(string),
            "body": __text__(string)
        }

Пример joke

        {
            "id": 1,
            "theme": "robots",
            "body": "\"Робот никогда не заменит человека!\" /Людоед/ "
        }

Пример расписания группы(расписание преподователя аналогично)

    {
        "group_name": "ИУ5-31Б",
        "week": {
            "monday": {
                "lessons": [
                    {
                        "time": "8:30 - 10:05",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "10:15 - 11:50",
                        "subject": {
                            "numerator": "Электротехника УЦ,362",
                            "denominator": "",
                            "is_differ": true
                        }
                    },
                    {
                        "time": "12:00 - 13:35",
                        "subject": {
                            "numerator": "Электротехника УЦ,362",
                            "denominator": "",
                            "is_differ": true
                        }
                    },
                    {
                        "time": "13:50 - 15:25",
                        "subject": {
                            "numerator": "Базовые компоненты Интернет-технологий,362",
                            "denominator": "Модели данных,306э",
                            "is_differ": true
                        }
                    },
                    {
                        "time": "15:40 - 17:15",
                        "subject": {
                            "numerator": "Электротехника,362",
                            "denominator": "Модели данных,306э",
                            "is_differ": true
                        }
                    },
                    {
                        "time": "17:25 - 19:00",
                        "subject": {
                            "numerator": "",
                            "denominator": "Модели данных,306э",
                            "is_differ": true
                        }
                    },
                    {
                        "time": "19:10 - 20:45",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    }
                ],
                "is_empty": false
            },
            "tuesday": {
                "lessons": [
                    {
                        "time": "8:30 - 10:05",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "10:15 - 11:50",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "12:00 - 13:35",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "13:50 - 15:25",
                        "subject": {
                            "numerator": "Элективный курс по физической культуре",
                            "denominator": "Элективный курс по физической культуре",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "15:40 - 17:15",
                        "subject": {
                            "numerator": "Базовые компоненты Интернет-технологий,224л",
                            "denominator": "Теория вероятностей и математическая статистика,224л",
                            "is_differ": true
                        }
                    },
                    {
                        "time": "17:25 - 19:00",
                        "subject": {
                            "numerator": "Архитектура АСОИУ,224л",
                            "denominator": "Архитектура АСОИУ,224л",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "19:10 - 20:45",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    }
                ],
                "is_empty": false
            },
            "wednesday": {
                "lessons": [
                    {
                        "time": "8:30 - 10:05",
                        "subject": {
                            "numerator": "Физика,кафедра",
                            "denominator": "",
                            "is_differ": true
                        }
                    },
                    {
                        "time": "10:15 - 11:50",
                        "subject": {
                            "numerator": "Физика,кафедра",
                            "denominator": "Физика,427ю",
                            "is_differ": true
                        }
                    },
                    {
                        "time": "12:00 - 13:35",
                        "subject": {
                            "numerator": "Модели данных,501ю",
                            "denominator": "Модели данных,501ю",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "13:50 - 15:25",
                        "subject": {
                            "numerator": "Электротехника,501ю",
                            "denominator": "Электротехника,501ю",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "15:40 - 17:15",
                        "subject": {
                            "numerator": "",
                            "denominator": "Электротехника,427ю",
                            "is_differ": true
                        }
                    },
                    {
                        "time": "17:25 - 19:00",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "19:10 - 20:45",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    }
                ],
                "is_empty": false
            },
            "thursday": {
                "lessons": [
                    {
                        "time": "8:30 - 10:05",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "10:15 - 11:50",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "12:00 - 13:35",
                        "subject": {
                            "numerator": "Модели данных УЦ,306э",
                            "denominator": "",
                            "is_differ": true
                        }
                    },
                    {
                        "time": "13:50 - 15:25",
                        "subject": {
                            "numerator": "Модели данных УЦ,306э",
                            "denominator": "",
                            "is_differ": true
                        }
                    },
                    {
                        "time": "15:40 - 17:15",
                        "subject": {
                            "numerator": "Физика,328",
                            "denominator": "Физика,328",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "17:25 - 19:00",
                        "subject": {
                            "numerator": "Правоведение,501ю",
                            "denominator": "Правоведение,501ю",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "19:10 - 20:45",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    }
                ],
                "is_empty": false
            },
            "friday": {
                "lessons": [
                    {
                        "time": "8:30 - 10:05",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "10:15 - 11:50",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "12:00 - 13:35",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "13:50 - 15:25",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "15:40 - 17:15",
                        "subject": {
                            "numerator": "Иностранный язык",
                            "denominator": "Иностранный язык",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "17:25 - 19:00",
                        "subject": {
                            "numerator": "Теория вероятности и математическая статистика,534л",
                            "denominator": "Теория вероятности и математическая статистика,534л",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "19:10 - 20:45",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    }
                ],
                "is_empty": false
            },
            "saturday": {
                "lessons": [
                    {
                        "time": "8:30 - 10:05",
                        "subject": {
                            "numerator": "Экология,114л",
                            "denominator": "Экология,218л",
                            "is_differ": true
                        }
                    },
                    {
                        "time": "10:15 - 11:50",
                        "subject": {
                            "numerator": "Теория вероятности и математическая статистика,218л",
                            "denominator": "Теория вероятности и математическая статистика,218л",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "12:00 - 13:35",
                        "subject": {
                            "numerator": "Правоведение,536л",
                            "denominator": "",
                            "is_differ": true
                        }
                    },
                    {
                        "time": "13:50 - 15:25",
                        "subject": {
                            "numerator": "Элективный курс по физической культуре",
                            "denominator": "Элективный курс по физической культуре",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "15:40 - 17:15",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "17:25 - 19:00",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    },
                    {
                        "time": "19:10 - 20:45",
                        "subject": {
                            "numerator": "",
                            "denominator": "",
                            "is_differ": false
                        }
                    }
                ],
                "is_empty": false
            }
        }
    }

Пример ответа на запрос: /api/v1/schedule/student_groups?days=monday,wednesday;limit=1

    [
        {
            "id": 10,
            "group_name": "ИУ5-31Б",
            "week": {
                "monday": {
                    "lessons": [
                        {
                            "time": "8:30 - 10:05",
                            "subject": {
                                "numerator": "",
                                "denominator": "",
                                "is_differ": false
                            }
                        },
                        {
                            "time": "10:15 - 11:50",
                            "subject": {
                                "numerator": "Электротехника УЦ,362",
                                "denominator": "",
                                "is_differ": true
                            }
                        },
                        {
                            "time": "12:00 - 13:35",
                            "subject": {
                                "numerator": "Электротехника УЦ,362",
                                "denominator": "",
                                "is_differ": true
                            }
                        },
                        {
                            "time": "13:50 - 15:25",
                            "subject": {
                                "numerator": "Базовые компоненты Интернет-технологий,362",
                                "denominator": "Модели данных,306э",
                                "is_differ": true
                            }
                        },
                        {
                            "time": "15:40 - 17:15",
                            "subject": {
                                "numerator": "Электротехника,362",
                                "denominator": "Модели данных,306э",
                                "is_differ": true
                            }
                        },
                        {
                            "time": "17:25 - 19:00",
                            "subject": {
                                "numerator": "",
                                "denominator": "Модели данных,306э",
                                "is_differ": true
                            }
                        },
                        {
                            "time": "19:10 - 20:45",
                            "subject": {
                                "numerator": "",
                                "denominator": "",
                                "is_differ": false
                            }
                        }
                    ],
                    "is_empty": false
                },
                "tuesday": null,
                "wednesday": {
                    "lessons": [
                        {
                            "time": "8:30 - 10:05",
                            "subject": {
                                "numerator": "Физика,кафедра",
                                "denominator": "",
                                "is_differ": true
                            }
                        },
                        {
                            "time": "10:15 - 11:50",
                            "subject": {
                                "numerator": "Физика,кафедра",
                                "denominator": "Физика,427ю",
                                "is_differ": true
                            }
                        },
                        {
                            "time": "12:00 - 13:35",
                            "subject": {
                                "numerator": "Модели данных,501ю",
                                "denominator": "Модели данных,501ю",
                                "is_differ": false
                            }
                        },
                        {
                            "time": "13:50 - 15:25",
                            "subject": {
                                "numerator": "Электротехника,501ю",
                                "denominator": "Электротехника,501ю",
                                "is_differ": false
                            }
                        },
                        {
                            "time": "15:40 - 17:15",
                            "subject": {
                                "numerator": "",
                                "denominator": "Электротехника,427ю",
                                "is_differ": true
                            }
                        },
                        {
                            "time": "17:25 - 19:00",
                            "subject": {
                                "numerator": "",
                                "denominator": "",
                                "is_differ": false
                            }
                        },
                        {
                            "time": "19:10 - 20:45",
                            "subject": {
                                "numerator": "",
                                "denominator": "",
                                "is_differ": false
                            }
                        }
                    ],
                    "is_empty": false
                },
                "thursday": null,
                "friday": null,
                "saturday": null
            }
        }
    ]


