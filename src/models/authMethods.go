package models

import (
)

type AUTH_METHOD struct {
    SAML bool
    MSFT bool
    GOOG bool
}

func GetAuthMethod(userID string) AUTH_METHOD {
    var a AUTH_METHOD

    // SELECT * FROM auth_methods JOIN users ON auth_methods.auth_id = users.auth_method WHERE user_id = $1
    
    // Set booleans accordingly

    return a
}

func GetSAMLIdP(auth_method string) string {
    // Query for IdP string given auth method string
    return ""
}
