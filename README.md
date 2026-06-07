# Rol Manager

Aplicación web para la gestión de partidas de rol de mesa.
Desarrollada en Go 1.22 utilizando únicamente la librería estándar.

## Requisitos

- Go 1.22 o superior

## Instalación y ejecución

1. Descargar o clonar el repositorio
2. Abrir una terminal en la carpeta del proyecto
3. Ejecutar el siguiente comando:

go run main.go

4. Abrir el navegador en: http://localhost:8080

## Estructura del proyecto

rolapp/
├── main.go              → punto de entrada y registro de rutas
├── go.mod               → módulo sin dependencias externas
├── models/
│   ├── db.go            → slices globales de datos
│   ├── character.go     → modelo de personaje
│   ├── game.go          → modelo de partida
│   └── registration.go  → modelo de inscripción
├── controllers/
│   ├── render.go        → función de renderizado de plantillas
│   ├── auth_controller.go
│   ├── home_controller.go
│   ├── character_controller.go
│   ├── game_controller.go
│   └── guide_controller.go
└── views/
    ├── layout.html      → plantilla base con navegación
    ├── login.html
    ├── home.html
    ├── games.html
    ├── game_detail.html
    ├── characters.html
    └── guide.html

## Uso básico

1. Acceder a la aplicación e introducir un nick de usuario
2. Como Master: crear una partida con límite de plazas
3. Como Jugador: crear un personaje e inscribirse en una partida
4. El Master acepta o rechaza las inscripciones desde el detalle de la partida

## Tecnología

- Lenguaje: Go 1.22
- Servidor HTTP: net/http (librería estándar)
- Plantillas: html/template (librería estándar)
- Persistencia: slices en memoria
- CSS: sin frameworks externos

## Autor

Sergio Mateos López