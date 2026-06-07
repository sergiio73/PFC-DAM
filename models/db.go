package models

// Esto hace de "base de datos falsa" usando slices en memoria.
// No se guarda en disco, así que al cerrar el programa se pierde todo.

var Characters []Character
var Games []Game
var Registrations []Registration

var NextCharacterID = 1
var NextGameID = 1
var NextRegistrationID = 1
