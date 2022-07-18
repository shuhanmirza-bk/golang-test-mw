package go_mw_service

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mw/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error

	globalConfig, err := util.LoadGlobalConfig("../")
	if err != nil {
		log.Fatal("can not load global config", err)
	}

	fmt.Println(globalConfig)

	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
