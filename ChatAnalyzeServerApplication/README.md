## Документация к AnalyzerService 

 Api для получения клиентом данных из информационной базы МГТУ им.Баумана посредством анализа текстового сообщения

### Доступные функции:

- Расписание преподавателей(на 6 рабочих дней)

### Дополнительные функции:

- Анекдоты

### Документация:

- "/api/v1/send_message" -
  - POST - отправка текстового сообщения для анализа и получения ответа
  
Структура JSON тела запроса

        {
            "message": __message__(string)
        }

Структура JSON файла расписания преподавателя

        {
            "id": __id__(int),
            "firstname": __name__(string),
            "surname": __surname__(string),
            "patronymic": __patronymic__(string),
            "chair": __chair__(string),
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

Ключ "theme" всегда должен иметь значение на русском языке

        {
            "id": 1,
            "theme": "робот",
            "body": "\"Робот никогда не заменит человека!\" /Людоед/ "
        }

Пример ответа на запрос: api/v1/send_message

Тело запроса:

    {
	    "message": "Отправь расписание Козлова на понедельник и пятницу"
    }
    
Ответ на запрос:

    [
        {
            "id": 1,
            "firstname": "Алекснадр",
            "surname": "Козлов",
            "patronymic": "Дмитриевич",
            "chair": "ИУ5",
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
                "wednesday": null,
                "thursday": null,
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
                "saturday": null
            }
        }
    ]
    
Пример ответа на запрос: api/v1/send_message
    
    Тело запроса:
    
        {
        	 "message": "Расскажи анектодты про роботов)"
        }
        
Ответ на запрос:
    
        [
            {
                "id": 1,
                "theme": "робот",
                "body": "\"Робот никогда не заменит человека!\" /Людоед/ "
            }
        ]

Пример ответа на запрос: api/v1/send_message
    
    Тело запроса:
    
        {
        	 "message": "Хочу анектодты)"
        }
        
Ответ на запрос:
    
        [
            {
                "id": 1,
                "theme": "робот",
                "body": "\"Робот никогда не заменит человека!\" /Людоед/ "
            },
            {
                "id": 2,
                "theme": "студент",
                "body": "— Профессор, а что вы говорите своим выпускникам при встрече? — Большую колу и картошечку фри, пожалуйста."
            }
        ]

