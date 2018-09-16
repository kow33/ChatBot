package com.example

import io.ktor.client.HttpClient
import io.ktor.client.engine.apache.Apache
import io.ktor.client.features.json.GsonSerializer
import io.ktor.client.features.json.JsonFeature
import io.ktor.client.request.get
import io.ktor.client.response.HttpResponse
import io.ktor.client.response.readText
import kotlinx.coroutines.experimental.runBlocking


fun messageParamAnalyze(message: String, clazz: Classifier): Map<String, String> {
    val params = mutableMapOf<String, String>()
    val client = HttpClient(Apache) {
        install(JsonFeature) {
            serializer = GsonSerializer()
        }
    }
    val messageWord=message.split(" ",",")
    runBlocking {
        when (clazz) {
            Classifier.Professor -> {
                val response = client.get<HttpResponse>("$BD_SERVER_URL/api/v1/schedule/info/professors")
                val professors = deserializationArrayFromGsonProfessor(response.readText())


                val professorSurnames = professors.filter { professor->
                    messageWord.any {word->
                        levensteinDistance(word,professor.surname)<=ProfessorLevensteinDistance
                    }
                }.map { it.surname }

                val daysOfWeek=daysOfWeekKeyWordRus.filter {day->
                    messageWord.any {word->
                        levensteinDistance(word,day)<= DaysLevensteinDistance
                    }
                }

                params["surname"] = if(professorSurnames.isNotEmpty()) professorSurnames.joinToString(",") else Error.UndefinedProfessorSurname.message
                params["days"]= daysOfWeek.joinToString(separator = ",") { day->
                    val index=daysOfWeekKeyWordRus.indexOf(day)
                    daysOfWeekKeyWordEng[index]
                }
            }
            Classifier.Joke -> {
                val response = client.get<HttpResponse>("$BD_SERVER_URL/api/v1/other_themes/jokes")
                //TODO do only one func for deserializing from GSON to Array
                val jokes = deserializationArrayFromGsonJoke(response.readText())
                val jokeThemes=if(messageWord.contains("про")){
                    jokes.find {joke->
                        messageWord.any {word->
                            word.contains(joke.theme, ignoreCase = true)
                        }
                    }?.theme?:Error.UndefinedJokeTheme.message
                }
                else {
                    ""
                }

                params["theme"] = jokeThemes
            }
            Classifier.Undefined->{

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

    return when {
        checkClassProfessor && !checkClassJoke -> Classifier.Professor
        !checkClassProfessor && checkClassJoke -> Classifier.Joke
        else -> Classifier.Undefined
    }

}