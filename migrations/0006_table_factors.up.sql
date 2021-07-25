BEGIN;

CREATE TABLE factors
(
    id            varchar(36) PRIMARY KEY                   NOT NULL,
    customer_id   varchar(36)                               NOT NULL,
    type          enum ('BUY','SELL') NOT NULL,
    shop_id       varchar(36)                               NOT NULl,
    factor_number varchar(36) UNIQUE                        NOT NULL,
    total_amount  decimal(64)                               NOT NULL,
    payed_amount  decimal(64) NULL,
    created_at    TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,
    updated_at    TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP (6) NOT NULL,
    CONSTRAINT factors_customer_id_fk
        FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE CASCADE,
        FOREIGN KEY (shop_id) REFERENCES shops (id) ON DELETE CASCADE
);

COMMIT;
