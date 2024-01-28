-- Script created to solve the excercise to do normalization, creation temporary tables and indexes


USE movies_db;


-- 1.  Agregar una película a la tabla movies.

INSERT INTO movies  (created_at,updated_at,title,rating,awards,release_date,length,genre_id)
VALUES ('2002-04-08 10:30:01', '2024-11-11 11:10:10', 'PELICULA MATEO', 5, 2, '2022-01-01', 120, 1);

-- 2.  Agregar un género a la tabla genres.
INSERT INTO genres  (created_at,updated_at,name,ranking,active)
VALUES ('2002-04-08 10:30:01', '2024-11-11 11:10:10',"Sci-fi",13,1);

-- 3.  Asociar a la película del punto 1. genre el género creado en el punto 2.
UPDATE movies as m SET m.genre_id = 16 WHERE m.title LIKE "%PELICULA%MATEO";

-- 5.  Crear una tabla temporal copia de la tabla movies.

CREATE TEMPORARY TABLE TempMovies AS SELECT * FROM movies m ;

SELECT * FROM TempMovies;

-- 6.  Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.

DELETE FROM TempMovies WHERE awards < 5 ;

-- 7.  Obtener la lista de todos los géneros que tengan al menos una película.

SELECT * FROM genres g WHERE g.active = 1;

-- 8.  Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.

SELECT a.* FROM actors a
INNER JOIN actor_movie am ON a.id  = am.actor_id
INNER JOIN movies m ON am.movie_id = m.id
WHERE m.awards > 3;

-- 9.  Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX indxNameMovies ON movies(title);

-- 10. Chequee que el índice fue creado correctamente.

SHOW INDEX FROM movies;
