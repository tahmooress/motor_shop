BEGIN;

CREATE TABLE accessibility
(
    id         varchar(36) PRIMARY KEY                   NOT NULL,
    admin_id   varchar(36)                               NOT NULL,
    shop_id    varchar(36)                               NOT NULL,
    created_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,
    updated_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP (6) NOT NULL,
    CONSTRAINT accessibility_admin_id_fk
        FOREIGN KEY (admin_id) REFERENCES admin_users (id) ON DELETE CASCADE,
    CONSTRAINT accessibility_shop_id_fk
        FOREIGN KEY (shop_id) REFERENCES shops (id) ON DELETE CASCADE
);

COMMIT;
