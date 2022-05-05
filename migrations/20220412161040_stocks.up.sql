-- stocks
create table stocks (
  id serial,
  symbol VARCHAR(255) NOT NULL,
  name_ja VARCHAR(255) NULL DEFAULT '',
  ticker VARCHAR(16) NOT NULL,
  exchange VARCHAR(255) NULL DEFAULT '',
  exchange_short_name VARCHAR(64) NULL DEFAULT '',
  price DECIMAL NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(id)
);
