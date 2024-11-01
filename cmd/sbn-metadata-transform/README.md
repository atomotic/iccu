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