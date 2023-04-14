
package exception

import (
	"encoding/json"
	"WorkerWithCheckHealth/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, ok := err.(ValidationError)
	if ok {
		var obj interface{}
		_ = json.Unmarshal([]byte(err.Error()), &obj)
		return ctx.Status(400).JSON(model.Response{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   struct{}{},
			Error:  obj,
		})
	}
	if err.Error() == model.AUTHENTICATION_FAILURE_ERR_TYPE {
		return ctx.Status(401).JSON(model.Response{
			Code:   401,
			Status: model.UNAUTHORIZATION,
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.AUTHENTICATION_FAILURE_ERR_TYPE,
			},
		})
	}

	return ctx.Status(500).JSON(model.Response{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
