# sbn-metadata-transform

transform source metadata (from sbn-metadata.db) and convert to simplified jsonl

```
~ go run . > sbn.jsonl 
```

load into DuckDB

```
~ duckdb sbn.duckdb "create table digital as select * from read_json_auto('sbn.jsonl');"
```

export to Parquet

```
~ duckdb sbn.duckdb "COPY digital TO 'sbn.digital.parquet' (FORMAT PARQUET);"
```

load from DuckDB shell https://shell.duckdb.org ([link](https://shell.duckdb.org/#queries=v0,CREATE-TABLE-digital-AS-FROM-'https%3A%2F%2Fatomotic.github.io%2Fdata%2Fsbn.digital.parquet'~,select-*-from-digital-limit-5~))

```
CREATE TABLE digital AS FROM 'https://atomotic.github.io/data/sbn.digital.parquet';

```

