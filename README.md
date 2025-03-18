# go-scaffold



## How to start project

#### 1. Install dependencies
```sh
make install-deps
cp dev/bruno/.env.example dev/bruno/.env
cp docker-compose.example.yaml docker-compose.yaml
docker compose up -d
docker compose logs -f app
```



## How to run tests and linters

#### 1. Generate mocks, docs and other dependencies with
```sh
make generate
```

#### 2. Run test
```sh
make test
```

#### 3. Run lint
```sh
make lint
```




## How to simulate app flows

### 1. receiving a book-created event and storing it on database

#### 1.1. Put `contracts/in/ms-books.book_created.json` into `ms-scaffold.books.created` queue
#### 1.2. Watch output into database. There must have a book item and chapter item
#### 1.3. Watch output into monitoring queue `ms-scaffold.*`. There must have a chapter-created message

### 2. creating an extra chapter through HTTP API

#### 2.1. Open /dev/bruno into bruno rest client
#### 2.2. Send post command (`Chapters/Create`)
#### 2.3. Watch output into database. There must have another chapter item
#### 2.4. Watch output into monitoring queue `ms-scaffold.*`. There must have a chapter-created message

