### 初回

```
make build
```

### 次回以降

```
make up
```

### migrate file作成

```
make create-migrate ARGS=filename
```

db/migrations配下に以下のファイルが作成される
xxxxxxx_filename.up.sql
xxxxxxx_filename.down.sql

### migrate up

```
make migrate-up
```
