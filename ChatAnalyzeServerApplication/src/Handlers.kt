package com.example

import io.ktor.application.ApplicationCall
import io.ktor.client.engine.apache.Apache
import io.ktor.client.HttpClient
import io.ktor.client.request.get
import io.ktor.client.response.HttpResponse
import io.ktor.client.response.readText
import io.ktor.http.HttpStatusCode
import io.ktor.response.respond
import kotlinx.coroutines.experimental.runBlocking


fun professorHandler(call: ApplicationCall, params: Map<String,String>) {
    val client=HttpClient(Apache){
        //configuration
    }
    runBlocking {
        when(params["surname"]){
            Error.UndefinedProfessorSurname.message->{
                call.respond(HttpStatusCode.BadRequest, mapOf(
                        "message" to Error.UndefinedProfessorSurname.message
                ))
            }
            ""->{
                val response=client.get<HttpResponse>("$BD_SERVER_URL/api/v1/schedule/professors/")
                call.respond(response.status,response.readText())
            }
            else->{
                val response=client.get<HttpResponse>("$BD_SERVER_URL/api/v1/schedule/professors/${params["surname"]}"+
                        "?"+
                        "days=${params["days"]}"
                )
                call.respond(response.status,response.readText())
            }
        }
    }

}

fun jokeHandler(call: ApplicationCall,params: Map<String,String>) {

}