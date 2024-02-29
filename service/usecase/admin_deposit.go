package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"meepshop_project/database"
	"meepshop_project/service/handler"
	"meepshop_project/service/request"
)

func AdminDeposit(ctx context.Context, req *request.AdminDeposit) (*handler.ResponseWithData, error) {

	tx := func(dc *gorm.DB) error {

		return nil
	}

	if err := database.Db.Transaction(tx); err != nil {
		fmt.Println(fmt.Errorf("[AdminDeposit]transaction err: %v", err))
		return nil, errors.New("unknown error")
	}
	return &handler.ResponseWithData{
		Data: "success",
	}, nil
}
