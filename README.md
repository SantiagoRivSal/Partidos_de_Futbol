# Torneos de Conmebol - Libertadores y Sudamericana

## Descripción ⚽

Este proyecto consiste en una aplicación para crear torneos de futbol de Conmebol, mas precisamente de Copa Libertadores y Sudamericana. La aplicacion ya tiene precargados algunos equipos de todas los paises Afiliados a la Conmebol (10 paises Sudamericanos). 

## Métodos API 🚀
### Confederaciones
- **GET (/confederaciones):** Obtiene la lista de todas las confederaciones existentes. En nuestro caso CONMEBOL.
### Equipos por Edicion
- **GET (/equiposxedicion/:id_edicion_torneo):** Obtiene la lista de todos los equipos que participan en una de edicion de un torneo puntual.
- **POST (/equipoxedicion):** Agrega un nuevo equipo a una edicion de un torneo.
### Edicion de Torneos
- **GET (/ediciones_de_torneo/:id_torneo):** Obtiene la lista de todas las ediciones de un torneo en partcular.
- **POST (/edicion_de_torneo):** Agrega una nueva edicion de un torneo.
### Resultado
- **GET (/resultados):** Obtiene la lista con los resultados de todos los torneos.
- **POST (/resultado):** Agrega un resultado de un torneo.
- **GET (/resultadoxedicion/:id_edicion_torneos):** Obtiene el resultado de una edicion de un torneo puntual.
### Equipos
- **GET (/equipos):** Obtiene la lista de todos los equipos.
- **GET (/equiposxpais/:id_pais):** Obtiene la lista de todos los equipos que pertenecen a un pais determinado.
- **GET (/equipo/:id):** Obtiene un equipo puntual.
### Paises
- **GET (/paises):** Obtiene la lista de todos los paises.
### Torneos
- **GET (/torneos):** Obtiene la lista de todos los torneos.

## Tecnologias Utilizadas 💻
### Docker 🐋
Docker es una plataforma de software que facilita la creación, implementación y administración de aplicaciones utilizando contenedores. Los contenedores permiten empaquetar una aplicación junto con todas sus dependencias en una unidad estándar, asegurando que la aplicación funcionará de la misma manera en cualquier entorno.

### GitHub Actions 💾
GitHub Actions es un servicio de automatización proporcionado por GitHub que te permite crear flujos de trabajo personalizados para tu repositorio. Puedes usarlo para automatizar tareas como pruebas de código, implementaciones y otras acciones basadas en eventos de tu repositorio de GitHub.

### Google Cloud ☁️
Google Cloud es una plataforma de servicios en la nube proporcionada por Google. Ofrece una amplia gama de servicios, incluyendo computación en la nube, almacenamiento de datos, aprendizaje automático, análisis de datos y más. Es utilizada por empresas y desarrolladores para construir, desplegar y escalar aplicaciones en la nube.

### Unit Test
Las pruebas unitarias (Unit Tests) son pruebas automatizadas que verifican el comportamiento de unidades individuales de código, como funciones o métodos, de manera aislada. Se utilizan para garantizar que cada unidad de código funcione correctamente de acuerdo con sus especificaciones.

### Test de Integracion
Los tests de integración son pruebas que verifican que las distintas unidades de código, módulos o sistemas interactúen entre sí de manera correcta cuando se combinan. Estas pruebas evalúan el comportamiento integrado de componentes más grandes y su interacción con sistemas externos o dependencias.

## Información del Desarrollador 🧑‍💻

- **Alumno:** Santiago A. Riveros Salomon
- **Mail:** 2000049@ucc.edu.ar
- **Materia:** Ingeniería de Software III

