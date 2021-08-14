package tests

import (
	"fileserver/cmd/client"
	"fileserver/controller"
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
	Router *gin.Engine
}

func (t *ULSuite) SetupSuite() {
	log.Println("Setup")
	t.Url = "http://localhost:8080/upload"
	t.Router = controller.SetupRouter("./temp/")
}

func (t *ULSuite) TearDownSuite() {
	log.Println("TearDown")
	// file delete all
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
