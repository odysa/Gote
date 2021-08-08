package errno

var (
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	PageNotFound        = &Errno{Code: 404, Message: "Path is not valid"}

	// user errors
	UserNotFound        = &Errno{Code: 20102, Message: "The user was not found."}
	UserNotLogin        = &Errno{Code: 20103, Message: "The use does not login"}
	UserInvalidPassword = &Errno{Code: 20104, Message: "Password is not valid"}
)
