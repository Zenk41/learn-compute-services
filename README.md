## **NOTE :**

### "***Untuk databasenya saya ganti dengan mysql karena saya belum pernah menggunakan sqlite***"

To generate mocks on folder businesses/users
```
cd businesses/users
mockery --all
```

To see the testing result 
```
go tool cover -html=coverage.out
             or
go tool cover -func=coverage.out
```