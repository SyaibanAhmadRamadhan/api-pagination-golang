package exception

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type ResponseError struct {
	Errors *any `json:"errors,omitempty"`
}

func Err(c *fiber.Ctx, err error) error {
	switch err.(type) {
	case *RespError:
	case *json.UnmarshalTypeError:
		message := fmt.Sprintf("type field must be %s", err.(*json.UnmarshalTypeError).Type.String())
		err = UnprocessableEntity(map[string]map[string]string{
			err.(*json.UnmarshalTypeError).Field: {
				"unprocess_entity": message,
			},
		})
	case *json.SyntaxError:
		err = BadRequest(map[string]map[string]string{
			"unexpected": {
				"JSON": "unexpected end of JSON input",
			},
		})
	case *mysql.MySQLError:
		err = InternalServerError(err.Error())
	default:
		if errors.Is(err, context.DeadlineExceeded) {
			err = RequestTimeOut("request time out")
		} else {
			err = InternalServerError(err.Error())
		}
	}

	return c.Status(err.(*RespError).Code).JSON(ResponseError{Errors: &err.(*RespError).Message})
}
