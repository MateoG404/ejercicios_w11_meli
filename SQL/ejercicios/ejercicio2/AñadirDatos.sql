USE ejercicio_1_SQL;

-- Incorporacion de 10 nuevos datos en la tabla clientes.

INSERT INTO clients (dni,first_name,last_name,birthdate,stateregion,city)
	VALUES
		("12233","Mariam","Zambrano","2003-12-05 12:34:56","cundinamarca","Madrid"),
        ("42433","Camila","Torres","1994-01-23 12:34:56","cundinamarca","bogota"),
        ("34535","Juan","Peralta","2014-04-25 12:34:56","cundinamarca","Choconta"),
        ("4256","mateo","Izquierdo","2002-01-25 12:34:56","cundinamarca","Bogota"),
        ("90423","Santiago","Peña","1994-01-25 12:34:56","cundinamarca","Suesca"),
        ("9453","mateo","Melo","2010-05-25 12:34:56","cundinamarca","bogota"),
        ("4243","Theofilo","gutierrez","2024-01-25 12:34:56","cundinamarca","Paipa"),
        ("120513","Ana","gutierrez","2021-11-25 12:34:56","cundinamarca","bogota"),
        ("6453","Catalina","Alfonso","2021-12-09 12:34:56","cundinamarca","Funza"),
        ("23403","Estefani","Fuentes","2022-01-15 12:34:56","cundinamarca","Madrid");

-- Incorporacion de 5 nuevos registros para la tabla internet_service

INSERT INTO internet_service (idinternet_service,speedInternet,price,discount,clients_dni)
	VALUES
    ("1","100","50","1","12233"),
	("2","24","23","31","34535"),
	("3","134","12","4","23403"),
	("4","997","312","5","23403"),
	("5","34","89","12","4243");
