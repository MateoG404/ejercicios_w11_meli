-- Ejercicios de clase consultas avanzadas SQL

-- Use the database

USE empresa_db;


-- 1. Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores

SELECT e.nombre, e.apellido,e.puesto ,d.localidad FROM empleados as e
INNER JOIN departamentos d ON e.depto_nro = d.depto_nro
HAVING e.puesto = "Vendedor";

-- 2. Visualizar los departamentos con más de dos empleados.

SELECT d.*
FROM departamentos as d
INNER JOIN empleados as e ON e.depto_nro = d.depto_nro
GROUP BY d.depto_nro HAVING COUNT(e.depto_nro)>2;

-- 3. Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.


SELECT e.nombre,e.salario, d.nombre_depto  FROM empleados as e
INNER JOIN departamentos as d ON e.depto_nro = d.depto_nro
WHERE e.puesto = (SELECT e.puesto FROM empleados e WHERE e.nombre LIKE 'Mito' AND e.apellido LIKE 'Barchuk');


-- 4. Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.

SELECT e.* FROM empleados as e
INNER JOIN departamentos d ON d.depto_nro = e.depto_nro
WHERE d.nombre_depto LIKE 'Contabilidad'
ORDER BY e.nombre ;


-- 5. Mostrar el nombre del empleado que tiene el salario más bajo.

SELECT * FROM 	empleados e
ORDER BY salario  LIMIT 1;

-- 6. Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
SELECT e.* FROM 	empleados e
INNER JOIN departamentos d ON e.depto_nro = d.depto_nro
WHERE d.nombre_depto LIKE 'Ventas'
ORDER BY salario  DESC LIMIT 1;
