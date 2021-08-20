BEGIN;

CREATE TABLE admin_users
(
    id         varchar(36) PRIMARY KEY                   NOT NULL,
    user_name  varchar(36) UNIQUE                        NOT NULL,
    password   varchar(255)                              NOT NULL,
    is_admin   tinyint(1) DEFAULT 0 NOT NULL,
    created_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,
    updated_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP (6) NOT NULL
);

COMMIT;
