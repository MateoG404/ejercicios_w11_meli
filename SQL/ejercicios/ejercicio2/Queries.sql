USE ejercicio_1_SQL;

-- 1.Mostrar todos los clientes y sus detalles.
SELECT * FROM clients;

-- 2. Obtener la lista de clientes que residen en una provincia específica.
SELECT * FROM clients WHERE stateregion LIKE '%Atlantico%';

-- 3. Encontrar los clientes cuyos apellidos empiezan con una letra determinada.
SELECT * FROM clients WHERE last_name LIKE 'A%';

-- 4. Obtener la lista de planes de internet y sus detalles.
SELECT * FROM internet_service;

-- 5. Encontrar los planes de internet con una velocidad superior a cierta cantidad de megas.
SELECT * FROM internet_service WHERE speedInternet> 80;

-- 6. Mostrar los clientes que tienen un descuento mayor del 10 en su plan de internet.
SELECT * FROM internet_service WHERE discount> 10;

-- 7. Obtener la cantidad total de clientes registrados en cada provincia.
SELECT stateregion, COUNT(*) as total_clients
FROM clients
GROUP BY stateregion;

-- 8. Encontrar los clientes que tienen un plan de internet con un precio superior a cierto valor.
SELECT * FROM clients INNER JOIN internet_service on clients.dni = internet_service.clients_dni WHERE internet_service.speedInternet > 70 ;

-- 9. Obtener la edad promedio de los clientes.

SELECT AVG(YEAR(CURDATE()) - YEAR(clients.birthdate)) FROM clients;

-- 10. Encontrar los clientes que cumplen años en los proximos 11 meses.
SELECT * FROM clients WHERE MONTH(clients.birthdate) = MONTH(CURDATE()) + 11;
