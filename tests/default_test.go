package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"path/filepath"
	_ "github.com/john-deng/k8s-cli-demo/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"net/url"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGet is a sample to run an endpoint test
func TestDeployment(t *testing.T) {
	data := url.Values{}
	data.Set("name", "k8s-cli-demo")
	data.Add("project", "demo-dev")
	r, _ := http.NewRequest("POST", "/k8s/deployment", strings.NewReader(data.Encode()))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
	        Convey("Status Code Should Be 200", func() {
	                So(w.Code, ShouldEqual, 200)
	        })
	        Convey("The Result Should Not Be Empty", func() {
	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
	        })
	})
}

