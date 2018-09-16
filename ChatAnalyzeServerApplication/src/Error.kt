package com.example

enum class Error(val message: String){
    UndefinedClass("""Извините, ваш запрос не получилось обработать на сервере. Пожалуйста, переформулируйте и отправьте еще раз."""),
    UndefinedProfessorSurname("Такого преподавтеля не существует, по крайнер мере, мне так программисты сказали."),
    UndefinedJokeTheme("Очень жаль, но анектдотов по это теме нет. Наши специалисты уже начали заниматься этой проблемой."),
    ManyProfessorsFoundedFirst("Какого преподавтеля вы имели ввиду?"),
    ManyProfessorsFoundedSecond("Введите порядковый номер:")
}