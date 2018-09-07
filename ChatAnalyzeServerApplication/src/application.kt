package com.example

import io.ktor.application.Application
import io.ktor.application.call
import io.ktor.application.install
import io.ktor.features.CallLogging
import io.ktor.features.Compression
import io.ktor.features.ContentNegotiation
import io.ktor.features.DefaultHeaders
import io.ktor.gson.gson
import io.ktor.http.HttpStatusCode
import io.ktor.response.respond
import io.ktor.routing.get
import io.ktor.routing.route
import io.ktor.routing.routing
import java.text.DateFormat
import java.time.LocalDate

fun main(args: Array<String>): Unit = io.ktor.server.netty.DevelopmentEngine.main(args)

data class Model(val name: String, val items: List<Item>, val date: LocalDate = LocalDate.of(2018, 4, 13))
data class Item(val key: String, val value: String)

val model = Model("root", listOf(Item("A", "Apache"), Item("B", "Bing")))

@Suppress("unused") // Referenced in application.conf
fun Application.module() {
    install(DefaultHeaders)
    install(Compression)
    install(CallLogging)
    install(ContentNegotiation) {
        gson {
            setDateFormat(DateFormat.LONG)
            setPrettyPrinting()
        }
    }
    install(Authentication) {
        basic {
            realm = "myrealm"
            validate { if (it.name == "user" && it.password == "password") UserIdPrincipal("user") else null }
        }
    }
    routing {

        route("/v1") {
            get {
                call.respond(model)
            }
            get("/item/{key}") {
                val item = model.items.firstOrNull { it.key == call.parameters["key"] }
                if (item == null)
                    call.respond(HttpStatusCode.NotFound)
                else
                    call.respond(item)
            }

        }
    }
}

