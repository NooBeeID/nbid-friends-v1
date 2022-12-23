# NBID - Be a Fullstack Development [Backend Session]
Halo, di repo ini saya tunjukin gimana struktur folder dan aplikasi yang biasa saya buat. Struktur folder ini saya gunakan jika aplikasi yg dibuat cukup complex

## How to run
Clone project ini dengan cara :
```bash
git clone -b main https://github.com/NooBeeID/nbid-friends-v1.git .
```

Lalu install dependensi yang akan digunakan :
```bash
go mod tidy
```

Setelah dependensi terinstall, silahkan copy file `cmd/.example.env` dan paste ke `cmd/.env`. Lalu isi dari file tersebut adalah :
```bash
DB_HOST=localhost       #isi dengan db host
DB_PORT=5432            #isi dengan port database
DB_USER=nbfriends       #isi dengan username database
DB_PASS=nbfriendlist    #isi dengan password database
DB_NAME=nbfriends       #isi dengan nama database
```

Jika telah membuat `environtment variable`, maka tinggal jalankan aplikasi dengan cara :
```bash
go run cmd/main.go
```

Jika berhasil, maka aplikasi akan berjalan di port `:4444`. Jika berhasil, maka akan muncul seperti output dibawah
```bash
go run cmd/main.go

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /v1/auth/register         --> backend/apps/domain/auth/controller.(*ControllerAPI).Register-fm (4 handlers)
[GIN-debug] POST   /v1/auth/login            --> backend/apps/domain/auth/controller.(*ControllerAPI).Login-fm (4 handlers)
[GIN-debug] POST   /v1/auth/search           --> backend/apps/domain/auth/controller.(*ControllerAPI).Search-fm (5 handlers)
[GIN-debug] GET    /v1/auth/profile          --> backend/apps/domain/auth/controller.(*ControllerAPI).Profile-fm (5 handlers)
[GIN-debug] POST   /v1/follow                --> backend/apps/domain/follow/controller.(*ControllerAPI).Follow-fm (5 handlers)
[GIN-debug] GET    /v1/follow                --> backend/apps/domain/follow/controller.(*ControllerAPI).GetMyFollowing-fm (5 handlers)
[GIN-debug] DELETE /v1/follow/               --> backend/apps/domain/follow/controller.(*ControllerAPI).UnfollFriend-fm (5 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :4444
```

## Struktur Folder
Saya menerapkan prinsip **Clean Architecture** dan **Domain Driven Design** versi saya. Jadi setiap services nya akan dipisah berdasarkan domainnya. Dan juga setiap layer akan mempunyai sebuah `interface` sebagai contract dengan layer lain.
```bash
├── apps
│   ├── commons
│   │   ├── middleware
│   │   │   └── gin.go
│   │   └── response
│   │       ├── error.go
│   │       └── success.go
│   ├── domain
│   │   ├── auth
│   │   │   ├── controller
│   │   │   │   └── api.go
│   │   │   ├── models
│   │   │   │   └── auth.go
│   │   │   ├── params
│   │   │   │   ├── request.go
│   │   │   │   └── response.go
│   │   │   ├── repositories
│   │   │   │   ├── postgres
│   │   │   │   │   └── auth.go
│   │   │   │   └── repositories.go
│   │   │   ├── route.go
│   │   │   └── services
│   │   │       ├── services.go
│   │   │       └── usecase
│   │   │           ├── base.go
│   │   │           ├── command.go
│   │   │           └── query.go
│   │   └── follow
│   │       ├── controller
│   │       │   └── api.go
│   │       ├── models
│   │       │   └── follow.go
│   │       ├── params
│   │       │   ├── request.go
│   │       │   └── response.go
│   │       ├── repositories
│   │       │   ├── postgres
│   │       │   │   └── follow.go
│   │       │   └── repositories.go
│   │       ├── route.go
│   │       └── services
│   │           ├── services.go
│   │           └── usecase
│   │               ├── base.go
│   │               ├── command.go
│   │               └── query.go
│   ├── router
│   │   └── gin.go
│   └── router.go
├── cmd
│   └── .env
│   └── main.go
├── config
│   ├── config.go
│   └── keys.go
├── deploy
│   ├── Dockerfile
│   ├── build.sh
│   └── run.sh
├── docs
│   ├── NBID-Friends.postman_collection.json
│   └── schema.sql
├── go.mod
├── go.sum
└── pkg
    ├── database
    │   ├── base.go
    │   ├── postgres.go
    │   └── postgres_test.go
    ├── encryption
    │   ├── hash.go
    │   └── hash_test.go
    └── token
        ├── payload.go
        ├── token.go
        └── token_test.go
```
Secara umum, ada beberapa folder utama yaitu :
- **cmd** - berfungsi untuk entry point dari aplikasi. 
- **apps** - berfungsi untuk meletakkan seluruh inti dari aplikasi
- **config** - berfungsi untuk menghandle configuration dari aplikasi
- **deploy** - berfungsi untuk menghandle command dan keperluan deployment
- **docs** - berfungsi dokumentasi yang di perlukan, seperti API Spec, Postman Collection, Schema database, dan lain lain.
- **pkg** - berfungsi untuk menyimpan custom package yg dibutuhkan.