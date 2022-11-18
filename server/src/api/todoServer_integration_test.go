package api

import (
	"net/http/httptest"
	"testing"

	"github.com/foxbit19/todo-app/server/src/model"
	"github.com/foxbit19/todo-app/server/src/store"
	testingCommon "github.com/foxbit19/todo-app/server/src/testing"
)

func TestStoreItemsAndRetrieveThem(t *testing.T) {
	database, cleanDb := testingCommon.CreateTempFile(t, "[]")
	defer cleanDb()
	server := NewTodoServer(&store.FileSystemStore{
		Database :database,
	})

	server.ServeHTTP(httptest.NewRecorder(), newPostTodoRequest(t, "I have to do some things, at first"))
	server.ServeHTTP(httptest.NewRecorder(), newPostTodoRequest(t, "Next, I have other things to do"))
	server.ServeHTTP(httptest.NewRecorder(), newPostTodoRequest(t, "Lastly, this is the last thing I have to do"))

	t.Run("Get items", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetAllTodosRequest())

		assertResponseStatus(t, response.Code, 200)
		assertAndGetAllJsonResponse(t, response.Body, &[]model.Item{
			{
				Id: 1,
				Description: "I have to do some things, at first",
				Order: 1,
			},
			{
				Id: 2,
				Description: "Next, I have other things to do",
				Order: 2,
			},
			{
				Id: 3,
				Description: "Lastly, this is the last thing I have to do",
				Order: 3,
			},
		})
	})
}