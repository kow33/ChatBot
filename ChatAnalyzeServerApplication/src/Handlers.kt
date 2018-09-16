package com.example

import com.google.gson.Gson
import io.ktor.application.ApplicationCall
import io.ktor.client.HttpClient
import io.ktor.client.engine.apache.Apache
import io.ktor.client.request.get
import io.ktor.client.response.HttpResponse
import io.ktor.client.response.readBytes
import io.ktor.client.response.readText
import io.ktor.content.TextContent
import io.ktor.http.ContentType
import io.ktor.http.HttpStatusCode
import io.ktor.http.cio.parseMultipart
import io.ktor.response.contentType
import io.ktor.response.respond
import io.ktor.response.respondText
import kotlinx.coroutines.experimental.runBlocking


fun professorHandler(call: ApplicationCall, params: Map<String, String>) {
    val client = HttpClient(Apache) {
        //configuration
    }
    runBlocking {
        when {
            (params["surname"] == Error.UndefinedProfessorSurname.message) -> {
                call.respond(HttpStatusCode.BadRequest, mapOf(
                        "message" to Error.UndefinedProfessorSurname.message
                ))
            }
            (params["surname"]!!.split(",").size > 1) -> {
                val response = client.get<HttpResponse>("$BD_SERVER_URL/api/v1/schedule/professors/${params["surname"]!!.split(",").first()}" +
                        "?" +
                        if(params["days"]?.isNotEmpty() == true) "days=${params["days"]}" else ""
                )
                val responseText=response.readText()
                val professors = deserializationArrayFromGsonProfessor(responseText)

                call.respond(TextContent(
                        "{" +
                                "\"message\": \" " +
                                    "${Error.ManyProfessorsFoundedFirst.message}\n" +
                                    "${professors.joinToString("\n") { "${professors.indexOf(it) + 1} - ${it.surname} ${it.firstname} ${it.patronymic}" }}\n" +
                                    "${Error.ManyProfessorsFoundedSecond.message}"+
                                "\","+
                                "\"body\" : $responseText,"+
                                "\"continue_dialog_on_client\" : true"+
                            "}"
                        ,ContentType.Application.Json))

            }
            else -> {
                val response = client.get<HttpResponse>("$BD_SERVER_URL/api/v1/schedule/professors/${params["surname"]}" +
                        "?" +
                        if(params["days"]?.isNotEmpty() == true) "days=${params["days"]}" else ""
                )
                call.respond(response.status, response.readText())
            }
        }
    }

}

fun jokeHandler(call: ApplicationCall, params: Map<String, String>) {
    val client = HttpClient(Apache) {
        //config
    }
    runBlocking {
        when (params["theme"]) {
            Error.UndefinedJokeTheme.message -> {
                call.respond(HttpStatusCode.BadRequest, mapOf(
                        "message" to Error.UndefinedJokeTheme.message
                ))
            }
            "" -> {
                val response = client.get<HttpResponse>("$BD_SERVER_URL/api/v1/other_themes/jokes")
                call.respond(response.status, response.readText())
            }
            else -> {
                val response = client.get<HttpResponse>("$BD_SERVER_URL/api/v1/other_themes/jokes/${params["theme"]}")
                call.respond(response.status, response.readText())
            }
        }
    }
}