package api

import (
	"net/http/httptest"
	"testing"

	"github.com/foxbit19/todo-app/server/model"
	"github.com/foxbit19/todo-app/server/store"
	testingCommon "github.com/foxbit19/todo-app/server/testing"
	"gotest.tools/v3/assert"
)

func TestStoreItemsAndRetrieveThem(t *testing.T) {
	database, cleanDb := testingCommon.CreateTempFile(t, "[]")
	defer cleanDb()
	store, _ := store.NewFileSystemStore(database)
	server := NewTodoServer(store)

	server.ServeHTTP(httptest.NewRecorder(), testingCommon.NewPostTodoRequest(t, "I have to do some things, at first"))
	server.ServeHTTP(httptest.NewRecorder(), testingCommon.NewPostTodoRequest(t, "Next, I have other things to do"))
	server.ServeHTTP(httptest.NewRecorder(), testingCommon.NewPostTodoRequest(t, "Lastly, this is the last thing I have to do"))

	t.Run("Get items", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, testingCommon.NewGetAllTodosRequest())

		assert.Equal(t, response.Code, 200)
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