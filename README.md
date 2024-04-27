# image-converter

Подпрограмма проекта J.A.R.V.I.S. для конвертации графических файлов.

Оригинальный проект: https://github.com/Priler/jarvis

Поддерживаемые форматы файлов: .jpg .png .jfif 

Преобразование:

Из jpg, jfif  в →  png

Из png, jfif  в → jpg

Использование:
Скачайте и установите Goland  https://go.dev/doc/install

После скачивания архива перейдите в директорию ahk

Установите библиотеку для конвертации графических файлов:

go get -u github.com/disintegration/imaging

Для компиляции программ используйте команды:

cd to_jpg

go build -ldflags="-s -w" -gcflags="all=-l" main.go

cd ..

cd to_png

go build -ldflags="-s -w" -gcflags="all=-l" main.go

После компиляции скопируйте "convert_files" в  проект JARVIS в папку commands
