# API de Bandas en Go

<!-- (Readme elaborado con Gemini) -->

Este proyecto consiste en una API RESTful construida completamente con la librería estándar de Go. Permite realizar operaciones CRUD (Crear, Leer, Actualizar, Borrar) sobre una colección de bandas de música.

El objetivo es demostrar un entendimiento práctico del protocolo HTTP, el diseño de APIs y la persistencia de datos sin depender de frameworks externos.

---

## 🚀 Características Principales

*   **API REST Completa:** Soporte para los métodos `GET`, `POST`, `PUT`, `PATCH` y `DELETE`.
*   **Persistencia de Datos:** Los cambios (crear, actualizar, borrar) se guardan de forma permanente en el archivo `data/bands.json`.
*   **Filtrado Avanzado:** El endpoint `GET /api/bands` permite filtrar por múltiples campos (`id`, `name`, `genre`, `year`, etc.) y combinar filtros.
*   **Validación Robusta:** Se valida la presencia de campos requeridos en las peticiones `POST` y `PUT`, y se comprueba la validez de los tipos de datos.
*   **Contenerización con Docker:** El proyecto incluye un `Dockerfile` y está listo para ser ejecutado de forma aislada y consistente con Docker.

---

## 📁 Estructura del Proyecto

```
.
├── main.go                 # Código fuente principal de la API
├── data/
│   └── bands.json          # Base de datos en formato JSON
├── postman/
│   ├── ej4-web.postman_collection.json # Colección de Postman
│   └── README.md           # Documentación de la colección
├── Dockerfile              # Define la imagen del contenedor
├── .dockerignore           # Archivos a ignorar por Docker
└── README.md               # Documentación principal del proyecto
```

---

## 🐳 Ejecución con Docker

La forma más sencilla de ejecutar el proyecto es utilizando Docker.

### Prerrequisitos
*   Tener Docker instalado.

### Instrucciones

1.  Clona el repositorio en tu máquina.
2.  Abre una terminal en la raíz del proyecto.
3.  Ejecuta el siguiente comando para construir la imagen y levantar el contenedor:
    ```bash
    docker compose up --build
    ```
4.  La API estará disponible en la siguiente URL: `http://localhost:41263`

---

## 📖 Endpoints de la API

A continuación se detallan los endpoints disponibles.

### 1. Verificar estado de la API
*   **Método:** `GET`
*   **URL:** `http://localhost:41263/api/ping`
*   **Descripción:** Endpoint simple para verificar que la API está en funcionamiento. Devuelve un `{"message":"pong"}`.

### 2. Obtener Bandas
*   **Método:** `GET`
*   **URL:** `http://localhost:41263/api/bands`
*   **Descripción:** Devuelve una lista de todas las bandas. Acepta query parameters para filtrar los resultados.
*   **Ejemplos de filtrado:**
    *   Por ID: `?id=1`
    *   Por nombre: `?name=metallica`
    *   Combinado: `?genre=Grunge&year=1987`

### 3. Crear una nueva Banda
*   **Método:** `POST`
*   **URL:** `http://localhost:41263/api/bands`
*   **Descripción:** Crea una nueva banda. Todos los campos son requeridos en el body.
*   **Body** (raw, JSON):
    ```json
    {
      "name": "New Band",
      "genre": "Rock",
      "year": 2023,
      "albums": 1,
      "members": 4
    }
    ```

### 4. Reemplazar una Banda (Actualización Completa)
*   **Método:** `PUT`
*   **URL:** `http://localhost:41263/api/bands?id={id}`
*   **Descripción:** Reemplaza completamente los datos de una banda existente. Requiere que todos los campos sean enviados en el body.
*   **Body** (raw, JSON):
    ```json
    {
      "name": "Updated Name",
      "genre": "Updated Genre",
      "year": 2024,
      "albums": 2,
      "members": 5
    }
    ```

### 5. Actualizar una Banda (Actualización Parcial)
*   **Método:** `PATCH`
*   **URL:** `http://localhost:41263/api/bands?id={id}`
*   **Descripción:** Actualiza uno o más campos de una banda existente sin afectar los demás.
*   **Body** (raw, JSON):
    ```json
    {
      "genre": "Progressive Rock"
    }
    ```

### 6. Eliminar una Banda
*   **Método:** `DELETE`
*   **URL:** `http://localhost:41263/api/bands?id={id}`
*   **Descripción:** Elimina una banda de la colección por su ID.

---

## 🧪 Pruebas y Documentación

Para facilitar las pruebas, el proyecto incluye una colección de Postman.

### Colección de Postman
*   **Archivo:** Puedes encontrar la colección para importar en `postman/ej4-web.postman_collection.json`.
*   **Cómo usar:** Abre Postman, haz clic en "Import" y selecciona el archivo.

### Documentación en la Nube
También puedes consultar la documentación interactiva y actualizada de la colección en Postman Cloud:

*   **Enlace:** Documentación de la API de Bandas

https://diego-glitch-cloud-4830510.postman.co/workspace/0c38c282-732f-41d3-8097-5f282a18b838/documentation/52958973-96c9e8b9-eb61-4139-82ef-2376e8f9e6d9
