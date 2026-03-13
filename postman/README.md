# Colección Postman para la API de Bandas

Esta carpeta contiene la colección de Postman (`ej4-web.postman_collection.json`) con todas las peticiones necesarias para interactuar y probar la API de bandas.

## Contenido de la Colección

La colección incluye ejemplos de peticiones para los siguientes endpoints y métodos:
-   **GET /api/ping**: Para verificar que la API está activa.
-   **GET /api/bands**: Para obtener todas las bandas o filtrar por `id`, `name`, `genre`, `year`, `albums` o `members`.
-   **POST /api/bands**: Para crear una nueva banda.
-   **PUT /api/bands?id={id}**: Para reemplazar completamente una banda existente por su ID.
-   **PATCH /api/bands?id={id}**: Para actualizar parcialmente los campos de una banda existente por su ID.
-   **DELETE /api/bands?id={id}**: Para eliminar una banda por su ID.

## Ejemplos de Uso (en Postman)

Asegúrate de que tu API esté corriendo (ej. con `docker compose up`) y accesible en `http://localhost:41263`.

### 1. Verificar estado de la API
*   **Método:** `GET`
*   **URL:** `http://localhost:41263/api/ping`

### 2. Obtener todas las bandas
*   **Método:** `GET`
*   **URL:** `http://localhost:41263/api/bands`

### 3. Obtener una banda por ID
*   **Método:** `GET`
*   **URL:** `http://localhost:41263/api/bands?id=1`

### 4. Filtrar bandas por nombre
*   **Método:** `GET`
*   **URL:** `http://localhost:41263/api/bands?name=metallica`

### 5. Filtrar bandas por género y miembros
*   **Método:** `GET`
*   **URL:** `http://localhost:41263/api/bands?genre=shoegaze&members=<5`

### 6. Crear una nueva banda
*   **Método:** `POST`
*   **URL:** `http://localhost:41263/api/bands`
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

### 7. Reemplazar completamente una banda
*   **Método:** `PUT`
*   **URL:** `http://localhost:41263/api/bands?id=1`
*   **Body** (raw, JSON):
    ```json
    {
      "name": "Updated Geese",
      "genre": "Indie Rock",
      "year": 2016,
      "albums": 3,
      "members": 4
    }
    ```

### 8. Actualizar parcialmente una banda
*   **Método:** `PATCH`
*   **URL:** `http://localhost:41263/api/bands?id=1`
*   **Body** (raw, JSON):
    ```json
    {
      "name": "Geese (Patched)"
    }
    ```

### 9. Eliminar una banda
*   **Método:** `DELETE`
*   **URL:** `http://localhost:41263/api/bands?id=1`

## Documentación en Postman Cloud

Puedes explorar la documentación interactiva de esta colección directamente en Postman Cloud a través del siguiente enlace:

Documentación de la Colección de Bandas
https://diego-glitch-cloud-4830510.postman.co/workspace/0c38c282-732f-41d3-8097-5f282a18b838/documentation/52958973-96c9e8b9-eb61-4139-82ef-2376e8f9e6d9

## Cómo Usar

1.  Abre Postman (o la extensión de Postman en VS Code).
2.  Utiliza la opción "Import" y selecciona el archivo `ej4-web.postman_collection.json` de esta carpeta.
3.  Una vez importada, podrás ejecutar todas las peticiones preconfiguradas contra tu API local.
