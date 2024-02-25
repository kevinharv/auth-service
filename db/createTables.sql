-- SAML IdPs
CREATE TABLE IF NOT EXISTS saml_idps (
    idp_id          UUID            UNIQUE NOT NULL,
    domain          VARCHAR(63)     UNIQUE NOT NULL,
    metadata_url    VARCHAR(253)    NOT NULL,
    is_enabled      BOOLEAN         DEFAULT(TRUE),
    created_at      TIMESTAMP,
    updated_at      TIMESTAMP,

    PRIMARY KEY (idp_id)
);

-- Authentication Methods
CREATE TABLE IF NOT EXISTS auth_methods (
    auth_id     UUID            UNIQUE NOT NULL,
    auth_name   VARCHAR(128)    UNIQUE NOT NULL,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP,

    PRIMARY KEY (auth_id)
);

-- USERS
CREATE TABLE IF NOT EXISTS users (
    user_id             UUID            UNIQUE NOT NULL,
    userPrincipalName   VARCHAR(253)    UNIQUE NOT NULL,
    auth_method         UUID,
    first_name          VARCHAR(128),
    last_name           VARCHAR(128),
    middle_init         CHAR(1),
    display_name        VARCHAR(64),
    created_at          TIMESTAMP,
    updated_at          TIMESTAMP,

    PRIMARY KEY (user_id),
    FOREIGN KEY (auth_method) REFERENCES auth_methods(auth_id)
);