package err

const (
	// Skip 0
	_ = iota

	// Password-related errors (100+)
	PasswordErrorStart = 100
	WrongPasswordBlank = PasswordErrorStart + iota
	WrongPasswordMatchConfirm
	IncorrectPassword
	AccountLocked

	// Authentication errors (200+)
	AuthErrorStart = 200
	AccessDenied   = AuthErrorStart + iota
	InvalidToken
	GentokenFailure
	SaveTokenFailure
	TokenExpired
	TokenMissing
	ValidateTokenFailure

	// Validation errors (300+)
	ValidationErrorStart = 300
	MissingField         = ValidationErrorStart + iota
	InvalidFormat
	ValueOutOfRange
	InvalidRole

	// Database errors (400+)
	DatabaseErrorStart = 400
	ChangePassword     = DatabaseErrorStart + iota
	RegisterUser
	NotFound

	// Server errors (500+)
	ServerErrorStart = 500
	InternalServer   = ServerErrorStart + iota
	DatabaseError
	ServiceUnavailable
)

var (
	// Password-related errors
	ErrWrongPasswordBlank        = newError(WrongPasswordBlank, "Password or confirm password cannot be blank", nil, "Missing password or confirmation password.")
	ErrWrongPasswordMatchConfirm = newError(WrongPasswordMatchConfirm, "Passwords do not match confirm", nil, "The password and confirmation password do not match.")
	ErrIncorrectPassword         = newError(IncorrectPassword, "Error verify password", nil, "Failed to verify the provided password.")
	ErrAccountLocked             = newError(AccountLocked, "The user account is temporarily locked due to too many failed login attempts.", nil, "Account locked due to multiple failed login attempts.")

	// Authentication errors
	ErrAccessDenied         = newError(AccessDenied, "The user does not have the necessary permissions.", nil, "Access denied for this operation.")
	ErrInvalidToken         = newError(InvalidToken, "The token provided is invalid.", nil, "Invalid or malformed token.")
	ErrGentokenFailure      = newError(GentokenFailure, "Unable to generate access token. Please try again later.", nil, "Failed to generate token due to internal error.")
	ErrSaveTokenFailure     = newError(SaveTokenFailure, "Failed to save token.", nil, "Error saving token to the database.")
	ErrTokenExpired         = newError(TokenExpired, "Token expired.", nil, "Token is no longer valid.")
	ErrTokenMissing         = newError(TokenMissing, "Token does not exist.", nil, "Request does not include a token.")
	ErrValidateTokenFailure = newError(ValidateTokenFailure, "Token is not valid.", nil, "Token validation failed.")

	// Validation errors
	ErrMissingField    = newError(MissingField, "A required field is missing in the request.", nil, "Field missing in the input payload.")
	ErrInvalidFormat   = newError(InvalidFormat, "Validation failed, the input format is invalid (e.g., invalid email or phone number).", nil, "Incorrect format provided.")
	ErrValueOutOfRange = newError(ValueOutOfRange, "A provided value exceeds allowed limits.", nil, "Value out of range for input field.")
	ErrInvalidRole     = newError(InvalidRole, "Invalid role provided in request.", nil, "Role is not recognized or supported.")

	// Database errors
	ErrChangePassword = newError(ChangePassword, "Error during query change password", nil, "Database error during password change operation.")
	ErrRegisterUser   = newError(RegisterUser, "Error during query create user", nil, "Database error during user registration.")
	ErrNotFound       = newError(NotFound, "Error during query find by phone", nil, "No record found for the given phone number.")

	// Server errors
	ErrInternalServer     = newError(InternalServer, "An unexpected error occurred on the server.", nil, "Internal server error.")
	ErrDatabaseError      = newError(DatabaseError, "A database operation failed.", nil, "Database error.")
	ErrServiceUnavailable = newError(ServiceUnavailable, "The service is temporarily unavailable.", nil, "Service is currently unavailable.")
)
