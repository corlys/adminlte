Make some .env file like this

```bash
DB_HOST=localhost
DB_USER=postgres
DB_PASS=velaryon
DB_NAME=adminlte
DB_PORT=5432

SESSION_SECRET=xxxVelaryonxxx
```

Make sure to install Go

and then do the package install

```bash
go mod tidy
```

and then make sure to install templ in terminal

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

Then run with this command

```bash
templ generate && go build -o adminlte ./cmd && ./adminlte
```
