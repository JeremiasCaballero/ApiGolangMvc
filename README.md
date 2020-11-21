# ApiGolangMvc
<h2>Pequeño proyecto en formato mvc para el seminario de golang<h2>

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
