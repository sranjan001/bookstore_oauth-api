CREATE KEYSPACE oauth WITH REPLICATION ={'class': 'SimpleStrategy','replication_factor': 1}
USE oauth

CREATE TABLE access_tokens( access_token varchar PRIMARY KEY, USER_ID bigint, client_id bigint, expires bigint);
