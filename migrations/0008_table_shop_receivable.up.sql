BEGIN;

CREATE TABLE shop_receivable
(
    id            varchar(36) PRIMARY KEY                   NOT NULL,
    customer_id   varchar(36)                               NOT NULl,
    factor_number varchar(36)                               NOT NULL,
    shop_id       varchar(36)                               NOT NULL,
    amount        decimal(64)                               NOT NULL,
    status        enum ('DEBTOR','CLEAR','DEFERRED') NOT NULL,
    clear_date    TIMESTAMP(6)                              NOT NULL,
    created_at    TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,
    updated_at    TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP (6) NOT NULL,
    CONSTRAINT shop_receivable_customer_id_fk
        FOREIGN KEY (customer_id) REFERENCES customers (id),
    CONSTRAINT shop_receivable_factor_number_fk
        FOREIGN KEY (factor_number) REFERENCES factors (factor_number),
    CONSTRAINT shop_receivable_shop_id_fk
        FOREIGN KEY (shop_id) REFERENCES shops (id)
);

COMMIT;
