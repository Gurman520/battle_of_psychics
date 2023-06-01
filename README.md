# battle_of_psychics

## Описание
Это тестовый проект суть которого заключается в том, что пользователь загадывает число, а несколько экстрасенсов пытаются отгадать загаданное число. 
По результатам предсказаний и ответа пользователя, меняется рейтинг каждого экстрасенса.

## API
Для реализации API используется Swagger V 2.0

## Точка входа
Точкой входа в приложение является _http://127.0.0.1:8080/Game_

## Запуск
Для запуска проекта необходимо выполнить команду _go run ./main.go_

## Быстрый запуск
Для быстро запуска проекта, можно воспользоваться docker-compose

Для этого необходимо:
1. Склонировать репозиторий
2. Перейти в него
3. Выполнить команду _docker compose build_
4. После этого выполнить вторую команду _docker compose up_
5. Дождаться запуска проекта и после этого перейти на страницу _http://127.0.0.1:8080/Game_
Именно с этой страницы начинается работа сервиса

