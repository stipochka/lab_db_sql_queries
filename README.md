![procedure1](https://github.com/user-attachments/assets/eb3a404c-0d99-4754-a747-90ac562dee9b)База данных, 4 Вариант:
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

Лабораторная №4:
Запросы для создания процедур/view/trigger находятся в файле: procedures_queries.txt
Задание №1: ![procedure1](https://github.com/user-attachments/assets/8c83c7da-aa96-4e48-822d-a3c69ba500b8)
Задание №2: ![procedure2](https://github.com/user-attachments/assets/1e1e2dfb-67ab-4576-ac82-e363a4e32166)
Задание №3: ![procedure3](https://github.com/user-attachments/assets/159f064d-3c7f-4fca-bfda-99cd0c15a467)
Задание №4: ![procedure4](https://github.com/user-attachments/assets/907be783-2cc4-42b1-898b-59a6dea047b0)


