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

## Información del Desarrollador 🧑‍💻

- **Alumno:** Santiago A. Riveros Salomon
- **Mail:** 2000049@ucc.edu.ar
- **Materia:** Ingeniería de Software III

## Docker 🐋
### Construir imágenes
Para esto es necesario ejecuta en el CMD la linea:
- docker-compose build

### Levantar contenedores
Para esto es necesario ejecuta en el CMD la linea:
- docker-compose up -d

### Subir a Docker Hub sanrivsal

Para subir a las imagenes creadas por docker a esta plataforma, se utilizan los siguientes comando:

- **Para iniciar sesion:** docker login
- **Etiquetar la imagen:** hay que tener en cuanta que nuestro codigo general 3 imagenes por lo cual los comandos son:
    - docker tag partidos_de_futbol-backend sanrivsal/partidos_de_futbol-backend:latest
    - docker tag partidos_de_futbol-database sanrivsal/partidos_de_futbol-database:latest
    - docker tag partidos_de_futbol-frontend sanrivsal/partidos_de_futbol-frontend:latest
- **Subir las Imágenes:** Para subir todas las imagenes que se etiquetaron se utilizan los siguiente comando:
    - docker push sanrivsal/partidos_de_futbol-backend:latest
    - docker push sanrivsal/partidos_de_futbol-database:latest
    - docker push sanrivsal/partidos_de_futbol-frontend:latest


