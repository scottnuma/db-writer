# db-writer

db-writer is a simple application that create and continually write timestamps into the stamps table. db-writer can be useful when you want to simulate continuous writes to a database.

db-writer requires setting the CONN_STRING env var. You can find this on Render as Postgres' External Database URL.

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy)

## Schema 

```sql
CREATE TABLE IF NOT EXISTS stamps (time timestamp);
```

```sql
INSERT into stamps (time) VALUES ($1);
```

```sql
SELECT * FROM stamps ORDER BY time DESC LIMIT 5;
```

