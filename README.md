База данных, 4 Вариант:
Запросы на создание таблиц находятся в папке: internal/sql/schema/, 

Таблица Medpersonal: 
id - SERIAL INT PRIMARY KEY, сделан уникальным и автоматически увеличавающимся
lastname VARCHAR(50) - фамилия сотрудника, ограничено изз-за размера фамилии
address VARCHAR(100) - VARCHAR(100) выбрано для более детального адреса
taxpercentage NUMERIC(4, 2) - вещественное число, чтобы отображать количество процентов до 99.99

таблица Workplace:
id - аналогичен id из таблицы Medpersonal
institution VARCHAR(100) - VARCHAR(100) выбрано для более детального названия учереждения
address VARCHAR(100) - VARCHAR(100) выбрано для более детального адреса
taxpercentage NUMERIC(4, 2) - вещественное число, чтобы отображать количество процентов до 99.99

таблица OperationType:
id - аналогичен id из таблицы Medpersonal
name VARCHAR(100) - название операции VARCHAR(100) выбрано для более детального название,
basepoint VARCHAR(100) NOT NULL - Опорный пункт (место выполнения операции),
stock INTEGER NOT NULL CHECK (Stock >= 0) - провера запасов, не могут быть отрицательными,
сost NUMERIC(10, 2) NOT NULL CHECK (Cost >= 0) - стоимость операции, не может быть отрицательным

таблица WorkActivity:
contract SERIAL PRIMARY KEY - уникальный номер контракта
date VARCHAR(50) - для хранения данных о дне недели
medpersonalid INTEGER REFERENCES Medpersonal(id) - cсылка на медперсонал
workplaceid INTEGER REFERENCES Workplace(id) - cсылка на место работы
operationid INTEGER REFERENCES OperationTypes(id) - cсылка на тип операции
quantity INTEGER CHECK (quantity > 0) - количество выполненных операций, проверка исключает неверные значения
payment NUMERIC(10, 2) - оплата за операцию

Запросы на вставку находятся в папке internal/sql/schema/queries/
Запросы находятся в папке sql_queries

