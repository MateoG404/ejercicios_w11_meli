-- Ejercitacion para consultas Avanzadas en SQL

USE movies_db;

-- 1. Mostrar el título y el nombre del género de todas las series.

SELECT s.title, g.name FROM series as `s` INNER JOIN genres as `g` ON s.genre_id = g.id;

-- 2. Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.

SELECT e.title, a.first_name, a.last_name FROM episodes as e
INNER JOIN actor_episode ae ON e.id = ae.episode_id
INNER JOIN actors a ON a.id = ae.actor_id ;

-- 3. Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.

SELECT s.title, COUNT(ss.serie_id) as `TotalSeasons` FROM series as s
INNER JOIN seasons ss ON s.id = ss.serie_id GROUP BY s.title ;

-- 4. Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.

SELECT g.name, COUNT(m.id) as `CantidadPelicula` FROM genres as g
INNER JoIN movies as m ON g.id = m.genre_id
GROUP BY g.name HAVING COUNT(m.id) > 2;

-- 5. Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan


SELECT DISTINCT a.first_name, a.last_name FROM actors as a
INNER JOIN actor_movie as am ON a.favorite_movie_id = am.actor_id
INNER JOIN movies as m ON am.movie_id  = m.id
WHERE m.title LIKE '%Guerra%de%las%galaxias%';
