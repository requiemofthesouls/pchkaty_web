package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"pchkaty_web/config"
	"pchkaty_web/controller"
	"pchkaty_web/db"
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
	controller.Pong(c)

	fmt.Printf("%s", w.Body.String())

	//err := json.Unmarshal(&got, w.Body.Bytes())

	//if err != nil {
	//	t.Fatal(err)
	//}
	//assert.Equal(t, want, got)

}
