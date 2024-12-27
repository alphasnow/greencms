del /q .\build\main.exe
rd /q /s .\build\website
rd /q /s .\build\storage
mkdir .\build\website
cd .\server\
go build -o main.exe main.go
move .\main.exe ..\build\main.exe
xcopy .\storage ..\build\storage /E /I /C
copy .\config.yaml ..\build\config.yaml
cd ..\
cd .\admin\
npm run build
move .\dist ..\build\website\admin
cd ..\
cd .\web\
npm run build
move .\out ..\build\website\web
cd ..\