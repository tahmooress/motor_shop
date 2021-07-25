BEGIN;

CREATE TABLE items
(
    id            varchar(36) PRIMARY KEY                   NOT NULL,
    pelak_number  varchar(36)                               NOT NULL,
    factor_number varchar(36) NULL,
    created_at    TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,
    updated_at    TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP (6) NOT NULL,
    CONSTRAINT items_pelak_number_fk
        FOREIGN KEY (pelak_number) REFERENCES motors (pelak_number),
    CONSTRAINT items_factor_number_fk
        FOREIGN KEY (factor_number) REFERENCES factors (factor_number)
);

COMMIT;
