# boilerplate-go

Este es un microservicio encargado de responder las pruebas

tenemos la carpeta kernel, que este se encarga de contener el core de la aplicacion, todos los utilitarios y ayudas requeridas
esto con el fin de no depender de externos

esta la carpeta microservicios donde se encuentran todos los microservicios separados, en este caso el microservicio Amaris

en collection esta un archivo de postman que puedes usar para probar los endpoint

esta en por defecto en 8080, este es un boilerplate en desarrollo, pero lo use para mejorarlo con esta oportunidad


para correr el microservicio debes entrar a su carpeta

cd microservices && cd amaris
y correr el clasico go run *.go
o usar el nodemon como es en mi caso preferido, ya que este me lo actualiza de una


-- ejecutar *go mod tidy* en caso necesario
-- ejecutar *go run *.go* en caso de desear probar

en caso de hacer todo con make usar 

-- ejecutar *make tidy* para descargar dependencias
-- ejecutar *make amaris*  para correr el microservicio