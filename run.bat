go build -o  Project.exe ./cmd/web/
Project.exe -dbname=bookings -dbuser=postgres -dbpass=Chaxeaoc123 -cache=false -production=false

@REM    .\run.bat