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
    MAX_FILE_SIZE_MB=5
    ALLOWED_FILE_EXTENSIONS=pdf,png,docx
    ALLOWED_ORIGINS=https://tudominio.com,http://localhost:3000
    ALLOWED_METHODS=GET,POST,PUT,DELETE,OPTIONS
    ALLOWED_HEADERS=Content-Type,Authorization,X-Requested-With
    CORS_ENABLED=true