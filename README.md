# users
Ini repository untuk service users

1. Lakukan copy file ".env.example" beberapa environment variable bisa diakses package godotenv melalui command line seperti di bawah ini :
```
cp .env.example .env
```
2. Bila belum ada database yang sama seperti di file ".env" . Lakukan create database (bisa dilakukan di aplikasi database management seperti DBeaver)
3. Bila belum dilakukan migration database (table belum dicreate). Execute command di terminal seperti di bawah ini :
```
go run main.go -migrate=migrate
```
4. Happy coding :)
