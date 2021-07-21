BEGIN;

CREATE TABLE motors
(
    id           varchar(36) PRIMARY KEY                   NOT NULL,
    model_name   varchar(255)                              NOT NULL,
    pelak_number varchar(9) UNIQUE                         NOT NULL,
    body_number  varchar(36) UNIQUE                        NOT NULL,
    color        varchar(36) NULL,
    model_year   varchar(4) NULL,
    created_at   TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,
    updated_at   TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP (6) NOT NULL
);

COMMIT;
