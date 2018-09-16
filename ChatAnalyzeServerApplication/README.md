## Документация к AnalyzerService 

 Api для получения клиентом данных из информационной базы МГТУ им.Баумана посредством анализа текстового сообщения

### Доступные функции:

- Расписание преподавателей(на 6 рабочих дней)

### Дополнительные функции:

- Анекдоты

### Оптимизация текста запроса:

В случае, когда пользователь ошибается при вводе текста, срабатывает алгоритм расстояниия Левенштейна . 
Допустимое колличество ошибочных символов:
- Для фамилий преподавателя: 3 символа
- Для дней недели: 3 символа

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
    
В случае, если преподавателей с запрашиваемой фамилией несколько, сервис отправляет список всех преподавтелей с такой фамилией, возлагая дальнейшую обработку текущего запроса на клиента.

Пример ответа на запрос: api/v1/send_message

Тело запроса:
    
        {
        	 "message": "Покажи расписание Козлова"
        }
    
Пример ответа на запрос: api/v1/send_message

Ответ на запрос:

    {
        "message": " Какого преподавтеля вы имели ввиду?\n1 - Козлов Алекснадр Дмитриевич\n2 - Козлов Сергей Константинович\nВведите порядковый номер:",
        "body": [
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
            },
            {
                "id": 3,
                "firstname": "Сергей",
                "surname": "Козлов",
                "patronymic": "Константинович",
                "chair": "РЛ3",
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
        ],
        "continue_dialog_on_client": true
    }
    
Пара <b>"continue_dialog_on_content" : true </b> оповещает клиента о том, что дальнейшая работа с запросом ведется на его стороне.    

    
Тело запроса:
    
        {
        	 "message": "Расскажи анекдоты про роботов)"
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
        	 "message": "Хочу анекдоты)"
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
        
###Ошибки

Формат выдачи ошибок:

      {
          "message": string
      }
      
Возможные ошибки и примеры:

- Не удалось определить тему запроса

        {
            "message": "Извините, ваш запрос не получилось обработать на сервере. Пожалуйста, переформулируйте и отправьте еще раз."
        }

- Преподавтеля с запршиваемой фамилией не существует:

        {
            "message": "Такого преподавтеля не существует, по крайнер мере, мне так программисты сказали."
        }
        
- Анектдоты с запрашиваемой темой отсутствуют в базе данных:
    
        {
            "message": "Очень жаль, но анектдотов по это теме нет. Наши анектодтоведы уже начали заниматься этой проблемой."
        }
