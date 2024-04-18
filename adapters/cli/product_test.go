package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/rafaelb13/full-cycle-hexagonal/adapters/cli"
	mock_application "github.com/rafaelb13/full-cycle-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 10.22
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)

	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
		productId,
		productName,
		productPrice,
		productStatus,
	)

	result, err := cli.Run(service, "create", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s with the name %s has been enabled",
		productId,
		productName,
	)

	result, err = cli.Run(service, "enable", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s with the name %s has been disabled",
		productId,
		productName,
	)

	result, err = cli.Run(service, "disable", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s with the name %s has the status %s and price %f",
		productId,
		productName,
		productStatus,
		productPrice,
	)

	result, err = cli.Run(service, "get", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
