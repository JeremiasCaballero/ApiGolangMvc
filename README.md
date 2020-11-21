# ApiGolangMvc
<img src="https://miro.medium.com/max/700/1*CdjOgfolLt_GNJYBzI-1QQ.jpeg">
<h3>Pequeño proyecto en formato mvc para el seminario de golang<h3>
<p>La rama main cuenta con la estructura mvc, y el modelo implementando con <a href="https://gorm.io/index.html">GORM</a><p>
<p>Se debe importar la base de datos o crear una (vuelo), las tablas que esta contenga se crean solas al iniciar el main.go</p>
<h3>Como iniciar la ApiRest</h3>
<ul>
<li>Importar la base de datos que se encuentra en /database, a su dominio</li>
<li>En caso de no importarla se creara automaticamente</li>
<li>Ejecutar main.go (go run main.go)</li>
<li>Abrir Postman o en el navegador localhost:8080/endpoint</li>
 </ul>


<h3>Tabla de ruteo</h3>
<table>
  <tr>
    <td>GetVuelos</td>
    <td>GetVuelo</td>
    <td>Addvuelo</td>
    <td>Deletevuelo</td>
    <td>Updatevuelo</td>
  </tr>
  <tr>
    <td>localhost:8080/getvuelos (trae todos los vuelos)</td>
    <td>localhost:8080/getvuelo/:id (trae un vuelo dado un id)</td>
    <td>localhost:8080/addvuelo/:id (agrega un vuelo dado un id)</td>
    <td>localhost:8080/deletevuelo/:id (elimina un vuelo dado un id)</td>
    <td>localhost:8080/updatevuelo/:id (actualiza un vuelo dado un id)</td>
  </tr>
</table>
<h3>Enviando un JSON</h3>
<p>Para enviar datos al servidor se necesita crear un <a href="https://www.json.org/json-es.html">JSON</a> con la siguiente estructura<p>
  

```javascript
{
   "Destino" "agregeUnDestino" // string
}

```
