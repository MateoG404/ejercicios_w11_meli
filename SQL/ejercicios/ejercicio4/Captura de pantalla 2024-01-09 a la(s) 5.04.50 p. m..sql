-- Ejercicios de clase consultas avanzadas SQL pt2

-- Use the database

USE library;

-- Create table Libro

CREATE TABLE IF NOT EXISTS libro (
    idlibro INT AUTO_INCREMENT,
    titulo VARCHAR(255),
    editorial VARCHAR(100),
    area VARCHAR(50),
    PRIMARY KEY(idlibro)
);

-- Create table Autor

CREATE TABLE IF NOT EXISTS autor(
	idautor INT AUTO_INCREMENT,
	nombre VARCHAR(100),
	nacionalidad VARCHAR(100),
	PRIMARY KEY(idautor)

);

-- Create table LibroAutor

CREATE TABLE IF NOT EXISTS libroAutor (
    idlibro INT,
    idautor INT,
    PRIMARY KEY(idlibro, idautor),
    FOREIGN KEY(idlibro) REFERENCES libro(idlibro),
    FOREIGN KEY(idautor) REFERENCES autor(idautor)
);

-- Create table estudiante

CREATE TABLE IF NOT EXISTS estudiante(
    idlector INT AUTO_INCREMENT,
    nombre VARCHAR(100),
    apellido VARCHAR(100),
    direccion VARCHAR(255),
    carrera VARCHAR(255),
    edad INT,
    PRIMARY KEY(idlector)
);

-- Create table prestamo

CREATE TABLE IF NOT EXISTS prestamo(
    idlector INT,
    idlibro INT,
    fechaprestamo DATE,
    fechadevolucion DATE,
    devuelto BOOL,
    PRIMARY KEY (idlibro,idlector),
    FOREIGN KEY(idlibro) REFERENCES libro(idlibro),
    FOREIGN KEY(idlector) REFERENCES estudiante(idlector)
);

---- Data Insertion ---

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
