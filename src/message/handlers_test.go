package message_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"

	"mist-io/src/message"
	pb "mist-io/src/protos/v1/gen"
)

func TestUpdateJwtToken(t *testing.T) {
	t.Run("successfully_updates_wsc_token", func(t *testing.T) {
		// ARRANGE
		t1 := "before"
		t2 := "after"
		updateMessage := &pb.Input_UpdateJwtToken{UpdateJwtToken: &pb.UpdateJwtToken{Access: t2}}
		wsc := &message.WsConnection{JwtToken: t1}

		// ASSERT
		assert.Equal(t, t1, wsc.JwtToken)

		// ACT
		wsc.UpdateJwtToken(updateMessage)

		// ASSERT
		assert.Equal(t, t2, wsc.JwtToken)
	})
}

func TestAppserverListing(t *testing.T) {
	t.Run("is_successful", func(t *testing.T) {
		// ARRANGE
		server1 := "foo"
		server2 := "bar"

		mockResponse := &pb.GetUserAppserverSubsResponse{}
		mockResponse.Appservers = []*pb.AppserverAndSub{
			{Appserver: &pb.Appserver{Name: server1}},
			{Appserver: &pb.Appserver{Name: server2}},
		}
		mockService := new(MockService)
		mockService.On("GetUserAppserverSubs", mock.Anything, mock.Anything).Return(mockResponse, nil)

		mockClient := new(MockClient)
		mockClient.On("GetServerClient").Return(mockService)

		wsc := &message.WsConnection{Client: mockClient}

		// ACT
		response, err := wsc.AppserverListing(&pb.Input_AppserverListing{})

		// ASSERT
		assert.Nil(t, err)
		mockClient.AssertExpectations(t)

		output := &pb.Output{}
		err = proto.Unmarshal(response, output)
		appservers := output.Data.(*pb.Output_AppserverListing).AppserverListing.Appservers

		assert.Nil(t, err)
		assert.Equal(t, appservers[0].Appserver.Name, server1)
		assert.Equal(t, appservers[1].Appserver.Name, server2)
	})

	t.Run("on_error_returns_error", func(t *testing.T) {
		// ARRANGE
		mockService := new(MockService)
		mockResponse := &pb.GetUserAppserverSubsResponse{}
		mockService.On("GetUserAppserverSubs", mock.Anything, mock.Anything).Return(mockResponse, errors.New("boom"))

		mockClient := new(MockClient)
		mockClient.On("GetServerClient").Return(mockService)

		wsc := &message.WsConnection{Client: mockClient}

		// ACT
		response, err := wsc.AppserverListing(&pb.Input_AppserverListing{})

		// ASSERT
		assert.NotNil(t, err)
		assert.Nil(t, response)
	})
}

func TestAppserverDetails(t *testing.T) {
	t.Run("is_successful", func(t *testing.T) {
		// ARRANGE
		server1 := "foo"
		appserverId := uuid.NewString()
		mockResponse := &pb.GetByIdAppserverResponse{}
		mockResponse.Appserver = &pb.Appserver{Name: server1}
		mockService := new(MockService)
		mockService.On(
			"GetByIdAppserver", mock.Anything, &pb.GetByIdAppserverRequest{Id: appserverId}).Return(mockResponse, nil)

		mockClient := new(MockClient)
		mockClient.On("GetServerClient").Return(mockService)

		wsc := &message.WsConnection{Client: mockClient}

		// ACT
		response, err := wsc.AppserverDetails(
			&pb.Input_AppserverDetails{AppserverDetails: &pb.GetByIdAppserverRequest{Id: appserverId}})

		// ASSERT
		assert.Nil(t, err)
		mockClient.AssertExpectations(t)

		output := &pb.Output{}
		err = proto.Unmarshal(response, output)
		appserver := output.Data.(*pb.Output_AppserverDetails).AppserverDetails.Appserver

		assert.Nil(t, err)
		assert.Equal(t, appserver.Name, server1)
	})

	t.Run("on_error_returns_error", func(t *testing.T) {
		// ARRANGE
		mockService := new(MockService)
		mockResponse := &pb.GetByIdAppserverResponse{}
		mockService.On("GetByIdAppserver", mock.Anything, mock.Anything).Return(mockResponse, errors.New("boom"))

		mockClient := new(MockClient)
		mockClient.On("GetServerClient").Return(mockService)

		wsc := &message.WsConnection{Client: mockClient}

		// ACT
		response, err := wsc.AppserverDetails(
			&pb.Input_AppserverDetails{AppserverDetails: &pb.GetByIdAppserverRequest{Id: uuid.NewString()}})

		// ASSERT
		assert.NotNil(t, err)
		assert.Nil(t, response)
	})
}

func TestAppserverCreate(t *testing.T) {
	t.Run("is_successful", func(t *testing.T) {
		// ARRANGE
		newserver := "new"
		server1 := "foo"
		server2 := "bar"
		mockCreateRequest := &pb.CreateAppserverRequest{Name: newserver}
		mockCreateResponse := &pb.CreateAppserverResponse{}
		mockResponse := &pb.GetUserAppserverSubsResponse{}
		mockResponse.Appservers = []*pb.AppserverAndSub{
			{Appserver: &pb.Appserver{Name: server1}},
			{Appserver: &pb.Appserver{Name: server2}},
		}
		mockService := new(MockService)
		mockService.On(
			"CreateAppserver", mock.Anything, mockCreateRequest,
		).Return(mockCreateResponse, nil)
		mockService.On("GetUserAppserverSubs", mock.Anything, mock.Anything).Return(mockResponse, nil)

		mockClient := new(MockClient)
		mockClient.On("GetServerClient").Return(mockService)

		wsc := &message.WsConnection{Client: mockClient}

		// ACT
		response, err := wsc.CreateAppserver(
			&pb.Input_CreateAppserver{CreateAppserver: mockCreateRequest},
		)

		// ASSERT
		assert.Nil(t, err)
		mockClient.AssertExpectations(t)

		output := &pb.Output{}
		err = proto.Unmarshal(response, output)
		appservers := output.Data.(*pb.Output_AppserverListing).AppserverListing.Appservers

		assert.Nil(t, err)
		assert.Equal(t, appservers[0].Appserver.Name, server1)
		assert.Equal(t, appservers[1].Appserver.Name, server2)
	})

	t.Run("on_error_when_creating_returns_error", func(t *testing.T) {
		// ARRANGE
		mockService := new(MockService)
		mockCreateRequest := &pb.CreateAppserverRequest{Name: "boom"}
		mockResponse := &pb.CreateAppserverResponse{}
		subResponse := &pb.GetUserAppserverSubsResponse{}
		mockService.On("CreateAppserver", mock.Anything, mock.Anything).Return(mockResponse, errors.New("boom"))
		mockService.On("GetUserAppserverSubs", mock.Anything, mock.Anything).Return(subResponse, errors.New("boom"))

		mockClient := new(MockClient)
		mockClient.On("GetServerClient").Return(mockService)

		wsc := &message.WsConnection{Client: mockClient}

		// ACT
		response, err := wsc.CreateAppserver(&pb.Input_CreateAppserver{CreateAppserver: mockCreateRequest})

		// ASSERT
		assert.NotNil(t, err)
		assert.Nil(t, response)
	})

	t.Run("on_error_when_fetching_subs_returns_error", func(t *testing.T) {
		// ARRANGE
		mockService := new(MockService)
		mockCreateRequest := &pb.CreateAppserverRequest{Name: "boom"}
		mockResponse := &pb.CreateAppserverResponse{}
		subResponse := &pb.GetUserAppserverSubsResponse{}
		mockService.On("CreateAppserver", mock.Anything, mock.Anything).Return(mockResponse, nil)
		mockService.On("GetUserAppserverSubs", mock.Anything, mock.Anything).Return(subResponse, errors.New("boom"))

		mockClient := new(MockClient)
		mockClient.On("GetServerClient").Return(mockService)

		wsc := &message.WsConnection{Client: mockClient}

		// ACT
		response, err := wsc.CreateAppserver(&pb.Input_CreateAppserver{CreateAppserver: mockCreateRequest})

		// ASSERT
		assert.NotNil(t, err)
		assert.Nil(t, response)
	})
}

func TestAppserverDelete(t *testing.T) {
	t.Run("is_successful", func(t *testing.T) {
		// ARRANGE
		someid := "someid"
		server1 := "s1"
		mockDeleteRequest := &pb.DeleteAppserverRequest{Id: someid}
		mockDeleteResponse := &pb.DeleteAppserverResponse{}
		mockResponse := &pb.GetUserAppserverSubsResponse{}
		mockResponse.Appservers = []*pb.AppserverAndSub{
			{Appserver: &pb.Appserver{Name: server1}},
		}
		mockService := new(MockService)
		mockService.On(
			"DeleteAppserver", mock.Anything, mockDeleteRequest,
		).Return(mockDeleteResponse, nil)
		mockService.On("GetUserAppserverSubs", mock.Anything, mock.Anything).Return(mockResponse, nil)

		mockClient := new(MockClient)
		mockClient.On("GetServerClient").Return(mockService)

		wsc := &message.WsConnection{Client: mockClient}

		// ACT
		response, err := wsc.DeleteAppserver(
			&pb.Input_DeleteAppserver{DeleteAppserver: mockDeleteRequest},
		)

		// ASSERT
		assert.Nil(t, err)
		mockClient.AssertExpectations(t)

		output := &pb.Output{}
		err = proto.Unmarshal(response, output)
		appservers := output.Data.(*pb.Output_AppserverListing).AppserverListing.Appservers

		assert.Nil(t, err)
		assert.Equal(t, appservers[0].Appserver.Name, server1)
	})

	t.Run("on_error_when_deleting_returns_error", func(t *testing.T) {
		// ARRANGE
		mockService := new(MockService)
		mockDeleteRequest := &pb.DeleteAppserverRequest{Id: "someid"}
		mockResponse := &pb.DeleteAppserverResponse{}
		subResponse := &pb.DeleteAppserverResponse{}
		mockService.On("DeleteAppserver", mock.Anything, mock.Anything).Return(mockResponse, errors.New("boom"))
		mockService.On("GetUserAppserverSubs", mock.Anything, mock.Anything).Return(subResponse, errors.New("boom"))

		mockClient := new(MockClient)
		mockClient.On("GetServerClient").Return(mockService)

		wsc := &message.WsConnection{Client: mockClient}

		// ACT
		response, err := wsc.DeleteAppserver(
			&pb.Input_DeleteAppserver{DeleteAppserver: mockDeleteRequest},
		)

		// ASSERT
		assert.NotNil(t, err)
		assert.Nil(t, response)
	})

	t.Run("on_error_when_fetching_subs_returns_error", func(t *testing.T) {
		// ARRANGE
		mockService := new(MockService)
		mockDeleteRequest := &pb.DeleteAppserverRequest{Id: "someid"}
		mockResponse := &pb.DeleteAppserverResponse{}
		subResponse := &pb.GetUserAppserverSubsResponse{}
		mockService.On("DeleteAppserver", mock.Anything, mock.Anything).Return(mockResponse, nil)
		mockService.On("GetUserAppserverSubs", mock.Anything, mock.Anything).Return(subResponse, errors.New("boom"))

		mockClient := new(MockClient)
		mockClient.On("GetServerClient").Return(mockService)

		wsc := &message.WsConnection{Client: mockClient}

		// ACT
		response, err := wsc.DeleteAppserver(
			&pb.Input_DeleteAppserver{DeleteAppserver: mockDeleteRequest},
		)

		// ASSERT
		assert.NotNil(t, err)
		assert.Nil(t, response)
	})
}
