package com.example

import com.google.gson.Gson
import com.google.gson.GsonBuilder
import io.ktor.client.HttpClient
import io.ktor.client.engine.apache.Apache
import io.ktor.client.features.json.GsonSerializer
import io.ktor.client.features.json.JsonFeature
import io.ktor.client.request.get
import io.ktor.client.response.HttpResponse
import io.ktor.client.response.readText
import io.ktor.gson.GsonConverter
import kotlinx.coroutines.experimental.runBlocking
import java.util.ArrayList
import com.google.gson.reflect.TypeToken




fun messageParamAnalyze(message: String, clazz: Classifier): Map<String, String> {
    val params = mutableMapOf<String, String>()
    val client = HttpClient(Apache) {
        install(JsonFeature) {
            serializer = GsonSerializer()
        }
    }
    runBlocking {
        when (clazz) {
            Classifier.Professor -> {
                val response = client.get<HttpResponse>("$BD_SERVER_URL/api/v1/schedule/info/professors")
                val professors= deserializationArrayFromGson<Professor>(response.readText())

                var professorSurname = message.split(",", " ").find { word ->
                    professors.any { word.contains(it.surname, ignoreCase = true) }
                } ?: Error.UndefinedProfessorSurname.message

                val daysOfWeek = message.split(","," ").filter {word->
                    daysOfWeekKeyWord.any { it.contains(word,ignoreCase = true) }
                }

                //Избавляемся от склонения
                professorSurname=professors.find { professorSurname.contains(it.surname,ignoreCase = true)}!!.surname

                params["surname"] = professorSurname
                params["days"]=daysOfWeek.joinToString(separator = ",")
            }
            Classifier.Joke -> {
                val jokes = client.get<List<Joke>>()
                val jokesTheme = message.split(",", " ").find { word ->
                    jokes.any { word.contains(it.theme, ignoreCase = true) }
                } ?: Error.UndefinedJokeTheme.message

                params["theme"] = jokesTheme

            }
        }
    }
    return params
}

fun messageClassAnalyze(message: String): Classifier {
    val wordList = message.split(",", " ")
    val checkClassProfessor: Boolean = wordList.any { word ->
        professorsKeyWord.any {
            word.contains(it, ignoreCase = true)
        }
    }
    val checkClassJoke: Boolean = wordList.any { word ->
        jokesKeyWord.any {
            word.contains(it, ignoreCase = true)
        }
    }

    return when (setOf(checkClassProfessor, checkClassJoke)) {
        setOf(true, false) -> Classifier.Professor
        setOf(false, true) -> Classifier.Joke
        else -> Classifier.Undefined
    }

}