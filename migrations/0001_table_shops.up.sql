BEGIN;

CREATE TABLE shops
(
    id         varchar(36) PRIMARY KEY                   NOT NULL,
    shop_name  varchar(36) UNIQUE                        NOT NULL,
    created_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,
    updated_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP (6) NOT NULL
);

COMMIT;
