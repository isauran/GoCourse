# Курс языка Go

[![Go Report Card](https://goreportcard.com/badge/github.com/RedHatOfficial/GoCourse)](https://goreportcard.com/report/github.com/RedHatOfficial/GoCourse)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/RedHatOfficial/GoCourse)

Этот репозиторий содержит материалы курса языка Go в формате Go [present](https://godoc.org/golang.org/x/tools/present). Чтобы просмотреть слайды, запустите команду present в корневой директории.

## Пошаговое руководство для просмотра слайдов локально

Предполагается, что у вас установлены и добавлены в $PATH компилятор go и git.

```shell
git clone https://github.com/RedHatOfficial/GoCourse.git
cd GoCourse
go run golang.org/x/tools/cmd/present
```

После этого подключитесь браузером к указанному адресу. Для остановки сервера используйте Ctrl+C.

## Предоставление доступа к слайдам другим людям

Возможно запустить службу, которая раздаёт слайды другим компьютерам через HTTP.
```shell
go run golang.org/x/tools/cmd/present -http=:3999 -play=false
```
