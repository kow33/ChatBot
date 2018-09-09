package com.example

enum class Error(val message: String){
    UndefinedClass("""Извините, ваш запрос не получилось обработать на сервере.
        |Пожалуйста, переформулируйте и отправьте еще раз.""".trimMargin()),
    UndefinedProfessorSurname("Такого преподавтеля не существует, по крайнер мере, мне так программисты сказали."),
    UndefinedJokeTheme("Очень жаль, что анектдотов по это теме нет.")
}