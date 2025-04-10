# GymI(Жыми) - Микросервисное приложение для отслеживания загруженности тренажерных залов

GymI — это приложение для отслеживания загруженности тренажерных залов в реальном времени. Оно позволяет пользователям вступать в группы залов, взаимодействовать с друзьями и назначать совместные тренировки.

## Функциональность

### 1. Аутентификация и авторизация
- Регистрация и вход пользователей
- Управление сессиями (JWT/Refresh токены)
- Поддержка ролей (администратор, пользователь)

### 2. Отображение залов
- Просмотр доступных залов в городе
- Фильтрация и поиск залов по различным критериям

### 3. Группы пользователей
- Проверка наличия пользователя в зале
- Автоматическое добавление пользователя в группу соответствующего зала
- Отображение статусов активности пользователей в группе (активен/неактивен)

### 4. Сервис опросов
- Назначение встреч с друзьями
- Ответы на опросы ("YES"/"NO")

## Архитектура

Проект построен на микросервисной архитектуре. Каждый компонент приложения является независимым микросервисом, который может быть развернут и масштабирован отдельно. В проекте используются следующие микросервисы:

### Микросервисы:
- **Authentication Service** — управление пользователями и сессиями (регистрация, вход, управление токенами).
- **Gym Service** — управление залами, отображение доступных залов, фильтрация и поиск.
- **Group Service** — управление группами пользователей в залах, проверка активности.
- **Poll Service** — создание и управление опросами, назначение встреч с друзьями.

## Технологии

- **Backend**: Golang
- **База данных**: PostgreSQL
- **Аутентификация**: JWT, Refresh токены
- **API**: REST
- **Архитектура**: Микросервисная
- **Контейнеризация**: Docker (опционально)

## Установка и настройка

1. **Клонировать репозиторий**:
   ```bash
   git clone https://github.com/yourusername/gym-tracker.git
   cd gym-tracker
