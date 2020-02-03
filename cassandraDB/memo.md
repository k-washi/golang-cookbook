# casanndra 

# mac install

```bash
brew install cassandra
sudo pip install cql

casanndra

cqlsh
```

# cmd

```bash
describe keyspaces;
>system_traces  system_schema  system_auth  system  system_distributed

CREATE KEYSPACE oauth WITH REPLICATION = {'class':'SimpleStrategy', 'replication_factor':1};

describe keyspaces;
>system_schema  system_auth  system  system_distributed  oauth  system_traces

USE oauth;

CREATE TABLE access_tokens( access_token varchar PRIMARY KEY, user_id bigint, expires bigint);

describe tables;
>access_tokens

SELECT * from access_tokens where access_token='doNaDonaEtc';

 access_token | expires | user_id
--------------+---------+---------


```

+ [replication_factor](https://www.intra-mart.jp/document/library/iap/public/imbox/cassandra_administrator_guide/texts/cluster/index.html)

# [gocql](https://github.com/gocql/gocql)

```bash
go get github.com/gocql/gocql
```