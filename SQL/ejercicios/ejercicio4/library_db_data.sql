-- Script to get and put new data in the db
USE library;

---- INSERT DATA ----

-- Insertar datos en la tabla libro
INSERT INTO libro (titulo, editorial, area) VALUES
('El Universo: Guía de viaje', 'Editorial XYZ', 'Ciencia'),
('Harry Potter', 'Salamandra', 'Literatura'),
('Base de Datos', 'Editorial ABC', 'Informática'),
('El código Da Vinci', 'Salamandra', 'Literatura'),
('El señor de los anillos', 'Editorial XYZ', 'Literatura');

-- Insertar datos en la tabla autor
INSERT INTO autor (nombre, nacionalidad) VALUES
('J.K. Rowling', 'Reino Unido'),
('Dan Brown', 'Estados Unidos'),
('J.R.R. Tolkien', 'Reino Unido');

-- Insertar datos en la tabla libroAutor (asumiendo que los id de los libros y autores insertados anteriormente son 3, 4, 5 y 3, 4, 5 respectivamente)
INSERT INTO libroAutor (idlibro, idautor) VALUES
(3, 3),
(4, 4),
(5, 5);

-- Insertar datos en la tabla estudiante
INSERT INTO estudiante (nombre, apellido, direccion, carrera, edad) VALUES
('Filippo', 'Galli', 'Calle 789', 'Informática', 23),
('Maria', 'Gonzalez', 'Avenida 012', 'Medicina', 25),
('Pedro', 'Gomez', 'Calle 345', 'Informática', 19),
('Lucia', 'Perez', 'Avenida 678', 'Ingeniería', 21);

-- Insertar datos en la tabla prestamo (asumiendo que los id de los estudiantes y libros insertados anteriormente son 3, 4, 5, 6 y 3, 4, 5 respectivamente)
INSERT INTO prestamo (idlector, idlibro, fechaprestamo, fechadevolucion, devuelto) VALUES
(3, 3, '2022-01-01', '2022-01-15', true),
(4, 4, '2022-01-05', '2022-01-20', false),
(5, 5, '2022-01-10', '2022-01-25', true),
(6, 3, '2022-01-15', '2022-01-30', false);


--- GET DATA---


-- 1. Listar los datos de los autores.

SELECT * FROM autor;

-- Listar nombre y edad de los estudiantes

SELECT e.nombre, e.edad FROM estudiante as e;

-- ¿Qué estudiantes pertenecen a la carrera informática?

SELECT e.* FROM estudiante as e WHERE e.carrera LIKE 'informatica';

-- ¿Qué autores son de nacionalidad francesa o italiana?

SELECT a.* FROM autor as a WHERE a.nacionalidad LIKE 'Francia' OR a.nacionalidad LIKE 'Italia';

-- ¿Qué libros no son del área de la informatica?

SELECT l.area FROM libro as l  WHERE l.area NOT LIKE 'informatica';

-- Listar los libros de la editorial Salamandra.

SELECT l.editorial  FROM libro as l  WHERE l.editorial  LIKE 'Salamandra';


-- Listar los datos de los estudiantes cuya edad es mayor al promedio.

SELECT e.* FROM estudiante as e WHERE  e.edad < (SELECT AVG(e2.edad) FROM estudiante as e2);

-- Listar los nombres de los estudiantes cuyo apellido comience con la letra G.

SELECT e.* FROM estudiante as e WHERE  e.apellido LIKE 'G%';

-- Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).

SELECT a.* FROM autor as a
INNER JOIN libroAutor as lb ON a.idautor = lb.idautor
INNER JOIN libro as l ON lb.idlibro = l.idlibro
WHERE l.titulo LIKE '%El universo%';


-- ¿Qué libros se prestaron al lector “Filippo Galli”?

SELECT l.titulo FROM libro as l
INNER JOIN prestamo as p ON l.idlibro = p.idlibro
INNER JOIN estudiante as e ON p.idlector = e.idlector
WHERE e.nombre LIKE 'Filippo%';

-- Listar el nombre del estudiante de menor edad.
SELECT e.* FROM estudiante as e WHERE e.edad < 18;

-- Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.
SELECT e.* FROM estudiante as e;

-- Listar los libros que pertenecen a la autora J.K. Rowling.

SELECT l.* FROM libro as l
INNER JOIN  libroAutor as lb ON lb.idlibro = l.idlibro
INNER JOIN autor  as a ON a.idautor = lb.idautor
WHERE a.nombre LIKE '%j.k%';
