# app-exercise API

[![Build Status](https://github.com/edurdo1901/app-exercise/workflows/Build/badge.svg?branch=main)](https://github.com/edurdo1901/app-exercise/actions?query=branch%3Amain) [![Coverage Status](https://coveralls.io/repos/github/edurdo1901/app-exercise/badge.svg?branch=main)](https://coveralls.io/github/edurdo1901/app-exercise?branch=main)

## Sitio de la aplicación

[Ejercicio amaris](https://app-amaris.prouddune-046dbdf6.eastus.azurecontainerapps.io) `https://app-amaris.prouddune-046dbdf6.eastus.azurecontainerapps.io`

## Documentación API

Los 3 `endpoints` se encuentran documentados en el siguiente archivo [ejercicios amaris swagger](docs/swagger.yaml), tambien se hizo, colección en postman para ejecutar los request [ejercicios amaris postman](docs/test.json)


## Requerimientos

- docker
- make
- go

## Ejecutar aplicación

ejecutamos `go run cmd/api/main.go` en el directorio raiz para correr el proyecto.

Se puede modificar el .env para ajustar el puerto por el que se está ejecutando la aplicación.

## Pruebas

Para ejecutar las pruebas unitarias, ejecutar el comando `make test`.

## Cobertura de pruebas

Para verificar la cobertura del código, ejecutar el comando `make test-cover`.

