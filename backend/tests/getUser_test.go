package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"pchkaty_web/backend/config"
	"pchkaty_web/backend/controller"
	"pchkaty_web/backend/db"
	"testing"
	//"github.com/stretchr/testify/assert"
)

func init() {
	config.LoadEnv()
	db.InitDB()
}

type resp struct {
	ping string
}

func TestPingRoute(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	controller.GetUser(c)

	fmt.Printf("%s", w.Body.String())

	//err := json.Unmarshal(&got, w.Body.Bytes())

	//if err != nil {
	//	t.Fatal(err)
	//}
	//assert.Equal(t, want, got)

}
