package tests

import (
	"fileserver/cmd/client"
	"fileserver/controller"
	"fileserver/service"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ULSuite struct {
	suite.Suite
	Url    string
	Store  *service.FileStore
	Router *gin.Engine
}

func (t *ULSuite) SetupSuite() {
	log.Println("Setup")
	folder := "./temp/"
	t.Url = "http://localhost:8080/upload"
	t.Store = &service.FileStore{
		Folder: folder,
	}
	router := gin.Default()
	controller.SetupUL(router, folder)
	controller.SetupDL(router, folder)
	t.Router = router
}

func (t *ULSuite) TearDownSuite() {
	log.Println("TearDown")
	// file delete all
	t.Store.DeleteAll()
}

func (t *ULSuite) TestHello() {
	log.Println("Hello")
}

func (t *ULSuite) TestUpload() {
	content, body, err := client.LoadFile("1628312808803.gif")
	require.NoError(t.T(), err)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/upload", body)
	// log.Println(contentType)
	req.Header.Add("Content-Type", content)
	require.NoError(t.T(), err)
	t.Router.ServeHTTP(w, req)
	require.Equal(t.T(), 200, w.Code)
	log.Println(w.Body)
}

func (t *ULSuite) TestHealth() {

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/health", nil)
	require.NoError(t.T(), err)
	t.Router.ServeHTTP(w, req)
	require.Equal(t.T(), 200, w.Code)
}

func TestULSuite(t *testing.T) {
	suite.Run(t, new(ULSuite))
}
