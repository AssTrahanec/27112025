## Запуск

```bash
go run .
```

Сервер запускается на порту 8080.

### POST /check

```bash
curl -X POST http://localhost:8080/check \
  -H "Content-Type: application/json" \
  -d '{"links": ["google.com", "yandex.ru"]}'
```
### POST /report

```bash
curl -X POST http://localhost:8080/report \
  -H "Content-Type: application/json" \
  -d '{"links_list": [1, 2]}' \
  --output report.pdf
```
## Особенности
- Данные сохраняются в `data.json`
- При остановке (Ctrl+C) данные автоматически сохраняются
