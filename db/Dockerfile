# Dockerfile para MySQL
FROM mysql

# Variables de entorno para configurar la base de datos
ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=fifa

# Copiar el archivo SQL a la carpeta de inicialización del MySQL
COPY fifa.sql /docker-entrypoint-initdb.d/fifa.sql
