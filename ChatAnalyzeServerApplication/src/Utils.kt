package com.example

import com.google.gson.Gson
import com.google.gson.reflect.TypeToken
import io.ktor.client.response.readText
import java.util.ArrayList

const val BD_SERVER_PORT=7777

const val BD_SERVER_URL="http://127.0.0.1:$BD_SERVER_PORT"

val professorsKeyWord= listOf(
        "расписани",
        "препод",
        "пара",
        "пары",
        "лектор",
        "семинар",
        "консульт",
        "заста",
        "искать",
        "найти"
)
val jokesKeyWord= listOf(
        "анек",
        "шут",
        "рассмеши"
)
val daysOfWeekKeyWordRus= listOf(
        "понедельник",
        "вторник",
        "сред",
        "четверг",
        "пятниц",
        "суббот",
        "воскресенье"
)
val daysOfWeekKeyWordEng= listOf(
        "monday",
        "tuesday",
        "wednesday",
        "thursday",
        "friday",
        "saturday",
        "sunday"
)

//TODO Обобщенка не работает, читать доки!!!
fun deserializationArrayFromGsonProfessor(json:String): List<Professor>{
    val listType = object : TypeToken<List<Professor>>(){}.type
    return Gson().fromJson<List<Professor>>(json,listType)
}

fun deserializationArrayFromGsonJoke(json:String): List<Joke>{
    val listType = object : TypeToken<List<Joke>>(){}.type
    return Gson().fromJson<List<Joke>>(json,listType)
}