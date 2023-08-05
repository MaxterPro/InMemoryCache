# Пакет In-Memory Cache
Пакет предназначен для сохранения данных в кэше по ключу

## Интефейсные функции
- `New()` - создание кэша
- `Set(key string, value interface{})` - запись данных `value` в кеш по ключу `key`
- `Get(key string)` - получить данные по ключу `key`
- `Delete(key)` - удалить данные по ключу `key`

## Инсталляция
go get -u github.com/MaxterPro/InMemoryCache

## Использование
package main

import (
	"fmt"

	cache "github.com/MaxterPro/InMemoryCache"
)

func main() {

	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)
}
