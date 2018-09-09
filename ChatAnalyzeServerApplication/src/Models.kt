package com.example

data class PostMessage(val message: String="")
data class Item(val key: String, val value: String)
data class Professor(val name:String, val surname:String, val patronymic: String, val chair: String)
data class Joke(val theme:String, val body:String)