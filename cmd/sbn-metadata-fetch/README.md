# sbn-metadata-fetch


oauth keys
```
export CLIENT_ID=
export CLIENT_SECRET=
```

run

```
go run .
```

query

```
~ sqlite3 sbn-metadata.db

sqlite> .schema
CREATE TABLE sbn (
                bid TEXT GENERATED ALWAYS AS (json_extract(doc, '$.unimarc.fields[1].003')) VIRTUAL, 
                doc json
        );
CREATE INDEX bid_idx on sbn(bid);

```
