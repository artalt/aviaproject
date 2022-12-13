Накатить все миграции:

**migrate -source file://migrations/postgresql -database "postgres://root:pass@localhost:5433/digdb?sslmode=disable" up**

Накатить одну миграцию:

**migrate -source file://migrations/postgresql -database "postgres://root:pass@localhost:5433/digdb?sslmode=disable" up 1**

Откатить все миграции:

**migrate -source file://migrations/postgresql -database "postgres://root:pass@localhost:5433/digdb?sslmode=disable" down**

Откатить одну миграцию:

**migrate -source file://migrations/postgresql -database "postgres://root:pass@localhost:5433/digdb?sslmode=disable" down 1**

Сгенерировать мок для тестирования

**mockgen -source=source.go > ./mock/source_mock.go**
