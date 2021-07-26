BEGIN;

CREATE TABLE transactions
(
    id            varchar(36) PRIMARY KEY                   NOT NULL,
    shop_id       varchar(36)                               NOT NULL,
    type          enum ('PAYED','RECEIVED') NOT NULL,
    subject       enum ('EXPENSES', 'EQUITY') NOT NULl,
    amount        decimal(64)                               NOT NULL,
    description   TEXT NUll,
    factor_number varchar(36) NULL,
    created_at    TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,
    updated_at    TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP (6) NOT NULL,
    CONSTRAINT transactions_shop_id_fk
        FOREIGN KEY (shop_id) REFERENCES shops (id)
--     CONSTRAINT transactions_factor_number_fk
--         FOREIGN KEY (factor_number) REFERENCES factors (factor_number)
);

COMMIT;
