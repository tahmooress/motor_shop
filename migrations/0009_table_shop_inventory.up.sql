BEGIN;

CREATE TABLE shop_inventory
(
    id            varchar(36) PRIMARY KEY                   NOT NULL,
    shop_id       varchar(36)                               NOT NULL,
    motor_id      varchar(36)                               NOT NULL,
    factor_number varchar(36)                               NOT NULL,
    created_at    TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,
    updated_at    TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP (6) NOT NULL,
    CONSTRAINT shop_inventory_shop_id_fk
        FOREIGN KEY (shop_id) REFERENCES shops (id),
    CONSTRAINT shop_inventory_motor_id_fk
        FOREIGN KEY (motor_id) REFERENCES motors (id),
    CONSTRAINT shop_inventory_factor_number_fk
        FOREIGN KEY (factor_number) REFERENCES factors (factor_number)
);

COMMIT;
