package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/rafaelb13/full-cycle-hexagonal/application"
	mockApplication "github.com/rafaelb13/full-cycle-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

// Função que testa o Service de Produtos
func TestProductService_Get(t *testing.T) {
	// Cria um novo controlador de mock
	ctrl := gomock.NewController(t)
	// Adia a execução do controlador
	defer ctrl.Finish()

	// Cria um novo mock de produto
	product := mockApplication.NewMockProductInterface(ctrl)
	// Cria um novo mock para a persistência de produtos
	persistence := mockApplication.NewMockProductPersistenceInterface(ctrl)

	// Define que o produto deve retornar um produto e nenhum erro
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	// Cria um novo serviço de produtos
	service := application.ProductService{
		// Define a persistência do serviço passando o mock de persistência
		Persistence: persistence,
	}

	// Chama o método Get do serviço passando um ID de produto
	result, err := service.Get("abc")
	// Verifica se não houve erro
	require.Nil(t, err)
	// Verifica se o produto retornado é igual ao produto esperado
	require.Equal(t, product, result)
}
