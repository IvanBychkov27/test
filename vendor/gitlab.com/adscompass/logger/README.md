# Logger

Логгер с возможность динамически включать/отключать вывовод определенных логов в рантайме

Для того, чтобы управлять выводом с помощью тегов, необходимо использовать специальные функции Tagged<LEVEL>

В этом примере сообщение будет показано, только если включен тег `mytag1`

```go
log.TaggedInfo("mytag1", "message")
```

Уровень лога также учитывается. То есть, сообщение не будет отображено, если тег включен, но уровень установлен выше.

Можно указать несколько тегов, разделяя их с помощью `logger.TagsSeprator`. В этом случае, сообщение будет показано, если включен хотя бы один тег

```go
tags := strings.Join([]string{ 
    "mytag1", 
    "mytag2",
    "mytag3",
    },
    logger.TagsSeprator 
)

log.TaggedInfo(tags, "message")
```

Чтобы показать сообщение, толко если включены ВСЕ теги, следует в начале строки добавить символ `!`

```go
log.TaggedInfo("!tag1|tag2", "message") // logger.TagsSeprator == '|' 
```

Сообщение будет показано, только если включены оба тега - `tag1` и `tag2`

## API

> New() *Logger

Создание нового логгера

> logger.Zap() *zap.Logger

Возвращает ссылку на Zap Logger

> logger.SetLevel(zapcore.Level)

Установить уровень логгирования

> logger.TaggedError(tags string, msg string, fields zap.Field...)
> logger.TaggedWarn(tags string, msg string, fields zap.Field...)
> logger.TaggedInfo(tags string, msg string, fields zap.Field...)
> logger.TaggedDebug(tags string, msg string, fields zap.Field...)

Вывод сообщения с учетом строки тегов

> logger.TagOn(<TAGNAME>)

Включить тег. (Включить отображение сообщений с этим тегом)

> logger.TagOff(<TAGNAME>)

Выключить тег. (Отключить отображение сообщений с этим тегом)

> logger.TagsOnAll()

Включить все теги

> logger.TagsOffAll()

Выключить все теги

> logger.Tags() []string 

Получить вчключенные теги

## Пример

```go
log, err := New()

// Обычный вывод черех zap.Logger
log.Zap().Info("Hello")

// Это сообщение не будет показано, пока не включен хотя бы один тег: tag1 или tag2
log.Zap().Info("Hello", zap.String("__tags__", "tag1|tag2")

log.SetTag("tag1")

// А теперь будет показано
log.Zap().Info("Hello", zap.String("__tags__", "tag1|tag2")

log.DelTag("tag1") 
// или
log.DelTags()

// Теперь опять не будет показано
log.Zap().Info("Hello", zap.String("__tags__", "tag1|tag2")


```