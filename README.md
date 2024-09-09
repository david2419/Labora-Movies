# **Proyecto Final: Sistema de Gestión de Películas al Estilo Netflix con Integración de API Externa**

## **Descripción**

El objetivo de este proyecto es que desarrollen un sistema completo de Administración de Películas, inspirado en la funcionalidad de Netflix. El sistema deberá incluir autenticación de usuarios, consumir datos de una API externa (The Movie Database), realizar operaciones de ABM (Alta, Baja, Modificación) sobre películas y usuarios, y ser desplegado en AWS utilizando servicios como EC2 y RDS.

### **Características Principales:**

1. **Consumo de API Externa:**
    - Utilizar la API de The Movie Database (TMDb) para obtener información detallada sobre películas. Deberán integrar los endpoints proporcionados por TMDb para acceder a la información necesaria.
2. **API:**
    - **Endpoints:**
        - **Visualizar Detalles de Película y Contador de Visualizaciones:** Solo funciona si eres un usuario registrado y estas loggeado
            - **Método:** GET
            - **Descripción:** Permite visualizar detalles específicos de una película identificada por su movie_id en TMDb (incluyendo comentarios de los usuarios). Además, incrementa en 1 el contador de visualizaciones en una tabla interna del sistema.
            - **Parámetros de Entrada:**
            - `movie_id` (identificación única de la película en TMDb)
        - **Obtener las n películas más visualizadas:**
            - **Método:** GET
            - **Descripción:** Retorna información sobre las n películas más visualizada en el sistema. “n” es un valor constante y arbitrario del sistema.
        - **Agregar Comentario:** Solo funciona si estas loggeado
            - **Método:** POST
            - **Descripción:** Permite a los usuarios agregar comentarios a una película en particular.
            - **Parámetros de Entrada:**
                - `usuario_id` (identificación única del usuario)
                - `movie_id` (identificación única de la película en TMDb)
                - `comentario` (texto del comentario)
        - **Excluir Comentario:** solo puedes eliminar comentarios tuyos
            - **Método:** DELETE
            - **Descripción:** Permite a los usuarios eliminar uno de sus comentarios previamente agregados.
            - **Parámetros de Entrada:**
                - `usuario_id` (identificación única del usuario)
                - `comentario_id` (identificación única del comentario)
        - **Editar Comentario:** solo puedes editar tus comentarios
            - **Método:** PUT
            - **Descripción:** Permite a los usuarios editar uno de sus comentarios previamente agregados.
            - **Parámetros de Entrada:**
                - `usuario_id` (identificación única del usuario)
                - `comentario_id` (identificación única del comentario)
                - `nuevo_texto` (nuevo texto del comentario)
        - **Endpoints para Gestión de Usuarios:** solo puedes editar informacion de tu usuario, debes estar loggeado
            - Permitir a los usuarios gestionar su cuenta, incluyendo la modificación de información personal (email, password, nickname).
                - **Endpoint para Modificar Información del Usuario:**
                    - **Método:** PUT
                    - **Descripción:** Permite al usuario modificar su información personal.
                    - **Parámetros de Entrada:**
                        - `usuario_id` (identificación única del usuario)
                        - `nuevos_datos` (datos actualizados del usuario, como email, password o nickname)


## **Recursos:**

- La documentación de la API de The Movie Database: [TMDb API](https://developer.themoviedb.org/docs/getting-started)
