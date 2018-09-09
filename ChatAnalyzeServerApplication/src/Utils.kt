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
val daysOfWeekKeyWord= listOf(
        "понедельник",
        "вторник",
        "сред",
        "четверг",
        "пятница",
        "суббота",
        "воскресенье"
)

fun <T> deserializationArrayFromGson(json:String): List<T>{
    val professorListType = object : TypeToken<ArrayList<Professor>>(){}.type
    return Gson().fromJson<List<T>>(json,professorListType)
}