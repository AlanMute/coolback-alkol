<p align="center">
      <img src="https://i.ibb.co/ckY7GkK/Logo-Eptanit.png" alt="Logo-Eptanit" border="0">
</p>

<p align="center">
   <img src="https://img.shields.io/badge/golang-v1.20.5-blue?style=plastic&logo=Go&label=Golang" alt="Golang Version">
      
   <img src="https://img.shields.io/badge/gin-v1.9.1-blue?style=plastic&label=Gin&labelColor=orange&link=https%3A%2F%2Fgithub.com%2Fgin-gonic%2Fgin" alt="Gin Version">
</p>

## О проекте

Этот проект представляет собой REST сервер, разработанный на языке Go с использованием фреймворка Gin. Он служит API для нашего веб-сайта с курсами eptanit.com.

## Функциональности

- Аутентификация с использованием JWT-токенов с помощью фреймфорка [jwt](https://github.com/golang-jwt/jwt)
- Взаимодействие с базой данных PostgreSQL с помощью ORM [Gorm](https://github.com/go-gorm/gorm).
- Настройка Swagger для документации API с помощью библиотеки [swaggo/swagg](https://github.com/swaggo/swag).
- Реализована рассылка сообщений по email адресам с помощью [gomail](https://github.com/go-gomail/gomail)
- Настройка Docker контейнеров

## Для запуска приложения:

Склонируйте репозиторий:
   
   ``` bash
   git clone https://github.com/KrizzMU/coolback-alkol.git
   ```

### Конфигурация .env

Прежде чем запустить проект, убедитесь, что в вашей директории проекта создан файл `.env` с необходимыми переменными окружения. Далее представлен шаблон содержимого файла `.env`:

```plaintext
JWT_SIGNING_KEY=
DB_PASSWORD=
DB_NAME=
DB_USER=
ADM_LOGIN=
ADM_PASSWORD=
MAIL_PASSWORD=
```

### Запуск сервера

1. Если у вас UNIX подобная ОС (таких как Linux и macOS), вополните следующие команды:
      - Запуск:
   
         ```
         make build
         ```
      - Запуск уже сбилженного контейнера:
   
         ```
         make run
         ```
2. Если в вашей OC не установлена утилита make, то можно выполнить следующие команды:
      - Запуск:
   
         ```
         docker-compose up --build epta-app
         ```
      - Запуск уже сбилженного контейнера:
   
         ```
         docker-compose up epta-app
         ```

## Разработчики
Разработкой сервера занимались:
- [Mutalibov Alaudin](https://github.com/KrizzMU)
- [Paradeev Nikolay](https://github.com/Cr4z1k)
