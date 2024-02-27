-- SAML IdPs
CREATE TABLE IF NOT EXISTS saml_idps (
    idp_id          UUID            UNIQUE NOT NULL DEFAULT gen_random_uuid(),
    display_name    VARCHAR(64)     UNIQUE NOT NULL,
    domain          VARCHAR(63)     UNIQUE NOT NULL,
    metadata_url    VARCHAR(253)    NOT NULL,
    is_enabled      BOOLEAN         DEFAULT(TRUE),
    created_at      TIMESTAMP       DEFAULT now(),
    updated_at      TIMESTAMP       DEFAULT now(),

    PRIMARY KEY (idp_id)
);

-- Authentication Methods
CREATE TABLE IF NOT EXISTS auth_methods (
    auth_id     UUID            UNIQUE NOT NULL DEFAULT gen_random_uuid(),
    auth_name   VARCHAR(128)    UNIQUE NOT NULL,
    isSAML      BOOLEAN         DEFAULT(FALSE),
    isMSFT      BOOLEAN         DEFAULT(FALSE),
    isGOOG      BOOLEAN         DEFAULT(FALSE),
    saml_idp_id UUID,
    created_at  TIMESTAMP       DEFAULT now(),
    updated_at  TIMESTAMP       DEFAULT now(),

    PRIMARY KEY (auth_id),
    FOREIGN KEY (saml_idp_id) REFERENCES saml_idps(idp_id)
);

-- Users
CREATE TABLE IF NOT EXISTS users (
    user_id             UUID            UNIQUE NOT NULL DEFAULT gen_random_uuid(),
    userPrincipalName   VARCHAR(253)    UNIQUE NOT NULL,
    auth_method         UUID,
    first_name          VARCHAR(128),
    last_name           VARCHAR(128),
    middle_init         CHAR(1),
    display_name        VARCHAR(64),
    created_at          TIMESTAMP       DEFAULT now(),
    updated_at          TIMESTAMP       DEFAULT now(),

    PRIMARY KEY (user_id),
    FOREIGN KEY (auth_method) REFERENCES auth_methods(auth_id)
);