# Microservicio de Archivos

Dependencias: 

	$ go get github.com/joho/godotenv
    $ go get -u github.com/gin-gonic/gin
    $ go get -u github.com/golang-jwt/jwt/v5

Archivo .env

    JWT_SECRET=mi_secreto_jwt_fuerte
    AUTH_HEADER=dXNlci1zdGlja3lfc2VjcmV0XzEyMzQ1Njc
    PORT=4000
    SECURE=false||true