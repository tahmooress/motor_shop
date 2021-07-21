BEGIN;

CREATE TABLE customers
(
    id            varchar(36) PRIMARY KEY                                                  NOT NULL,
    name          varchar(255)                                                             NOT NULL,
    last_name     varchar(255)                                                             NOT NULL,
    mobile        varchar(11) UNIQUE                                                       NOT NULL,
    national_code varchar(36) UNIQUE                                                       NULL,
    company_name  varchar(255)                                                             NULL,
    created_at    TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6)                                NOT NULL,
    updated_at    TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) NOT NULL
);

COMMIT;
