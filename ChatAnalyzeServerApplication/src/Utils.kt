package com.example

import com.google.gson.Gson
import com.google.gson.reflect.TypeToken
import java.util.Arrays



const val BD_SERVER_PORT=7777

const val BD_SERVER_URL="http://127.0.0.1:$BD_SERVER_PORT"

const val ProfessorLevensteinDistance=3

const val DaysLevensteinDistance=3

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
        "среда",
        "четверг",
        "пятница",
        "суббота",
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

fun levensteinDistance(x: String, y: String): Int {
    val dp = Array(x.length + 1) { IntArray(y.length + 1) }

    for (i in 0..x.length) {
        for (j in 0..y.length) {
            when {
                i == 0 -> dp[i][j] = j
                j == 0 -> dp[i][j] = i
                else -> dp[i][j] = min(dp[i - 1][j - 1] + costOfSubstitution(x[i - 1], y[j - 1]),
                        dp[i - 1][j] + 1,
                        dp[i][j - 1] + 1)
            }
        }
    }

    return dp[x.length][y.length]
}

fun costOfSubstitution(a: Char, b: Char): Int {
    return if (a == b) 0 else 1
}

fun min(vararg numbers: Int): Int {
    return Arrays.stream(numbers)
            .min().orElse(Integer.MAX_VALUE)
}