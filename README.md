# ground-control

## Вышка наземного движения

### Спецификация

[![Swagger-UI](https://img.shields.io/badge/Swagger-UI-brightgreen?logo=swagger)](https://docs.reaport.ru/)

### Описание

Отвечает за то, как ездят машины и самолеты. Все машины, самолёты двигаются по одинаковым правилам, на всех нужна одна библиотека. Карта аэропорта - это граф. На каждой части графа может быть только 1 машина. Машина доезжает до точки и запрашивает, можно ли двигаться на опред точку? Вышка говорит можно/нельзя. Машина доезжает и передает вышке, что плечо свободно и на него может приезжать другая машина (для самолетов одни места, для машин другие). Машины должны выезжать и заезжать в гараж. Параллельное обслуживание 2х (больше?) самолётов.

### Пример схемы

![](/assets/scheme.png "Схема")![alt text](image.png)