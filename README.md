# check-password-strength
Checks whether a password is valid given a list of rules.

Options:

- minSize: password is at least n characters long.
- minUppercase: password contains at least n uppercase letters.
- minLowercase: password contains at least n lowercase letters.
- minDigit: password contains at least n digits (0-9).
- minSpecialChars: password contains at least n special characters (neither a letter nor a digit).
- noRepeated: password must not contain sequences of the same characters -- 'aba' is allowed, but 'aab' is not.
