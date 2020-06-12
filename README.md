Descripción 

Se desarrolla microservicio que expone una API con 3 endpoints. 
Su dominio son Sales, y el alcance:

 1. sales x month 
 2. sales by product 
 3. latest sales

Componentes
Existirá un servicio SalesServicios con los métodos:

 1. getSales
 2. getSalesByProduct
 3. getLatestSales

El cuál obtendrá la data de un repositorio SalesRepository que puede satisfacerse con una db o un mock local.
Las peticiones entrantes se manejarán con un Router y un SalesHandler que hará la llamada al servicio y se encargará de responder con un código http y un mensaje.

Pendientes

#server
#middleware(validaciones)
#configuracion
#testing
#mocks
#dockerizar

steps

Crear directorio y archivos iniciales
Iniciar módulo go mod init ms-sale-transaction
Crear main
Crear Interfaz de servicio
Métodos del servicio
Pruebas unitarias servicio
Repositorio
Implementación origen de datos mysql
Pruebas unitarias repositorio
Router
Definir contrato
Handlers
Probar
