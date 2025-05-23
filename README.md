# Ogen для генерации кода по OpenAPI

В данном репозитории представлен тестовый пример взаимодействия с `ogen` для генерации кода по `OpenAPI` спецификации. Была поставлена следующая задача: разобраться в генераторе кода по `OpenAPI` - `Ogen`. И реализовать изменение формата сгенерированного кода - добавить комментарий `openapi` спецификацией к каждому методу.

Что для этого было сделано:

1) Составлен базовый `openapi` в файле `api.yml`. В нём есть один эндпоинт - **register**, который позволяет зарегистрировать пользователя по **email** и **password**. При успешной регистрации возвращается токен (заглушка под **JWT** токен). Если **email** уже существует, то возвращается ошибка "email уже существует".
2) С помощью `ogen` была сгенерирована дирректория `gen` командой `ogen --target gen --clean api.yml`
3) В файле **server.go** была реализована структура **userService**, соответствующая интерфейсу **Handler**. Она содержит in-memory хранилище пар (email, password) и реализует хендлер регистрации **RegisterPost**. В функции **main** запускается сам сервер
4) В файле **client.go** был реализован тестоый клиент, чтобы посылать запросы на сервер. Также в конце файла есть **curl** запрос на сервер

Вторая часть задачи - добавить нужные поля в `openapi`, чтобы были видны комментарии к методам.

За это отвечают поля **summary** и **description**. Правда при генерации **summary** используется, если не указан **description**

1) *summary* является кратким описанием метода. В моём случае это "Регистрация нового пользователя". В сгенерированном коде это можно увидеть, посмотрев интерфейс **Handler** или метод **handleRegisterPostRequest**:

2) *description* является более подробным описанием и используется при генерации по умолчанию. В моём случае это "Регистрирует нового пользователя по email и password". Аналогично описание прописывается в интерфейсе **Handler** и методе **handleRegisterPostRequest**

Интерфейс **Handler** находится в файле `oas_server_gen.go`, метод **handleRegisterPostRequest** в файле `oas_handlers_gen.go`

Для удобства были сделаны скриншоты соответсвующих комментариев в коде и добавлены в папку `images`