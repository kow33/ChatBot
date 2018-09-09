package com.example

import io.ktor.application.Application
import io.ktor.application.call
import io.ktor.application.install
import io.ktor.features.ContentNegotiation
import io.ktor.gson.gson
import io.ktor.request.receive
import io.ktor.response.respond
import io.ktor.routing.post
import io.ktor.routing.route
import io.ktor.routing.routing
import java.text.DateFormat

fun main(args: Array<String>): Unit = io.ktor.server.netty.DevelopmentEngine.main(args)

@Suppress("unused") // Referenced in application.conf
fun Application.module() {
    install(ContentNegotiation) {
        gson {
            setDateFormat(DateFormat.LONG)
            setPrettyPrinting()
        }
    }
    routing {

        route("/api/v1") {
            post("/send_message") {
                val post=call.receive<PostMessage>()

                val clazz=messageClassAnalyze(post.message)
                val params=messageParamAnalyze(post.message, clazz)

                when (clazz) {
                    Classifier.Professor -> professorHandler(call,params)
                    Classifier.Joke -> jokeHandler(call,params)
                    else -> call.respond(mapOf("message" to Error.UndefinedClass))
                }
            }
//            get("/item/{key}") {
////                val item = model.items.firstOrNull { it.key == call.parameters["key"] }
//                if (item == null)
//                    call.respond(HttpStatusCode.NotFound)
//                else
//                    call.respond(item)
//            }

        }
    }
}

