# API Error Codes Documentation

This document provides a detailed explanation of error codes and their corresponding messages for the frontend to interpret. Each error code is designed to represent a specific type of error, making it easier to identify and troubleshoot issues.

## Error Code Format
Error codes follow a structured format:

- **Prefix**: Represents the category of the error (e.g., `WP` for authentication issues).
- **Number**: A unique identifier within the category.

### Example:
- `WP101`: Authentication error indicating "Wrong Password."

---

# Error Documentation

## Password-Related Errors

| Error Code | Error Name                    | Message                                                             | Description                                                      |
|------------|-------------------------------|---------------------------------------------------------------------|------------------------------------------------------------------|
| `WP100`    | `WrongPasswordBlank`          | `Password or confirm password cannot be blank.`                      | Missing password or confirmation password.                       |
| `WP101`    | `WrongPasswordMatchConfirm`   | `Passwords do not match confirm.`                                    | The password and confirmation password do not match.             |
| `WP102`    | `IncorrectPassword`           | `Error verify password.`                                             | Failed to verify the provided password.                          |
| `WP103`    | `AccountLocked`               | `The user account is temporarily locked due to too many failed login attempts.` | Account locked due to multiple failed login attempts. |

---

## Authentication Errors

| Error Code | Error Name               | Message                                                | Description                                   |
|------------|--------------------------|--------------------------------------------------------|-----------------------------------------------|
| `AU200`    | `AccessDenied`           | `The user does not have the necessary permissions.`    | Access denied for this operation.             |
| `AU201`    | `InvalidToken`           | `The token provided is invalid.`                      | Invalid or malformed token.                   |
| `AU202`    | `GentokenFailure`        | `Unable to generate access token. Please try again later.` | Failed to generate token due to internal error. |
| `AU203`    | `SaveTokenFailure`       | `Failed to save token.`                               | Error saving token to the database.           |
| `AU204`    | `TokenExpired`           | `Token expired.`                                      | Token is no longer valid.                     |
| `AU205`    | `TokenMissing`           | `Token does not exist.`                               | Request does not include a token.             |
| `AU206`    | `ValidateTokenFailure`   | `Token is not valid.`                                 | Token validation failed.                      |

---

## Validation Errors

| Error Code | Error Name            | Message                                                  | Description                               |
|------------|-----------------------|----------------------------------------------------------|-------------------------------------------|
| `VE300`    | `MissingField`        | `A required field is missing in the request.`            | Field missing in the input payload.       |
| `VE301`    | `InvalidFormat`       | `Validation failed, the input format is invalid (e.g., invalid email or phone number).` | Incorrect format provided.               |
| `VE302`    | `ValueOutOfRange`     | `A provided value exceeds allowed limits.`               | Value out of range for input field.       |
| `VE303`    | `InvalidRole`         | `Invalid role provided in request.`                      | Role is not recognized or supported.      |

---

## Database Errors

| Error Code | Error Name            | Message                                                  | Description                               |
|------------|-----------------------|----------------------------------------------------------|-------------------------------------------|
| `DB400`    | `ChangePassword`      | `Error during query change password.`                    | Database error during password change operation. |
| `DB401`    | `RegisterUser`        | `Error during query create user.`                        | Database error during user registration.  |
| `DB402`    | `NotFound`            | `Error during query find by phone.`                      | No record found for the given phone number. |

---

## Server Errors

| Error Code | Error Name            | Message                                                  | Description                               |
|------------|-----------------------|----------------------------------------------------------|-------------------------------------------|
| `SE500`    | `InternalServer`      | `An unexpected error occurred on the server.`            | Internal server error.                    |
| `SE501`    | `DatabaseError`       | `A database operation failed.`                           | Database error.                           |
| `SE502`    | `ServiceUnavailable`  | `The service is temporarily unavailable.`                | Service is currently unavailable.         |

---

### Notes:
- **Error Code Format**: `{Prefix}{Error Number}` (e.g., `WP100`, `AU200`).
- **Prefix Meaning**:
  - `WP`: Password-related errors
  - `AU`: Authentication errors
  - `VE`: Validation errors
  - `DB`: Database errors
  - `SE`: Server errors
- **Structure**:
  - **Error Name**: The symbolic name of the error.
  - **Message**: The user-facing message describing the error.
  - **Description**: A more detailed explanation for debugging purposes.

--- 

Tài liệu này được sắp xếp theo nhóm lỗi rõ ràng và đồng bộ với quy định mã lỗi của bạn. 😊