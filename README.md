# check-password-strength
Checks whether a password is valid given a list of rules.

### Options:

- minSize: password is at least n characters long.
- minUppercase: password contains at least n uppercase letters.
- minLowercase: password contains at least n lowercase letters.
- minDigit: password contains at least n digits (0-9).
- minSpecialChars: password contains at least n special characters (neither a letter nor a digit).
- noRepeated: password must not contain sequences of the same characters -- 'aba' is allowed, but 'aab' is not. The value for this rule is always 0 (See example for clarification).

### Usage:

Send a POST request with a GraphQL query to http://localhost:8080/graphql containing the password and list of rules.

### Example:

```
Request:

query {
    verify(password: "IsPasswordStrong!123&", rules: [
      {rule: "minSize", value: 8},
      {rule: "minSpecialChars", value: 2},
      {rule: "noRepeated", value: 0},
      {rule: "minDigit", value: 4}
    ])  {
         verify
         noMatch
    }
}
```
```
Response:

{
    "data": {
        "verify": {
            "verify": false,
            "noMatch": [
                "noRepeated",
                "minDigit"
            ]
        }
    }
}
``` 
