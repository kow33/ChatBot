package com.example

data class PostMessage (val message: String, val clarify: Boolean)
data class Item (val key: String, val value: String)
data class Professor (val firstname:String, val surname:String, val patronymic: String, val chair: String)
data class Joke (val theme:String, val body:String)