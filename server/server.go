package server

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/74th/vscode-book-r2-golang/domain/entity"
	"github.com/74th/vscode-book-r2-golang/domain/usecase"
	"github.com/74th/vscode-book-r2-golang/memdb"
)

// サーバAPI
type Server struct {
	server     http.Server
	interactor usecase.Interactor
}

// サーバAPIのインスタンスを作成する
func New(addr string, webroot string) *Server {
	s := &Server{
		server: http.Server{
			Addr: addr,
		},
		interactor: usecase.Interactor{
			Database: memdb.NewDB(),
		},
	}

	s.setRouter(webroot)

	return s
}

// サーバを開始する
func (s *Server) Serve() error {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("could not start server: %s", err.Error())
		return err
	}
	return nil
}

// ルータの設定
func (s *Server) setRouter(webroot string) {
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/tasks", s.list)
	api.POST("/tasks", s.create)
	api.PATCH("/tasks/:id/done", s.done)
	api.POST("/tasks/alldone", s.done)

	router.StaticFile("/", filepath.Join(webroot, "index.html"))
	router.Static("/js", filepath.Join(webroot, "js"))
	router.Static("/css", filepath.Join(webroot, "css"))
	s.server.Handler = router
}

// GET /tasks
// タスク一覧
func (s *Server) list(c *gin.Context) {
	tasks, err := s.interactor.ShowTasks()
	if err != nil {
		log.Print("error", err)
		c.Status(500)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// POST /tasks
// タスクの追加
func (s *Server) create(c *gin.Context) {
	task := new(entity.Task)

	err := c.ShouldBindJSON(task)
	if err != nil {
		log.Print("deserialize error", err)
		c.Status(401)
		return
	}

	task, err = s.interactor.CreateTask(task)
	if err != nil {
		log.Print("error", err)
		c.Status(500)
		return
	}

	c.JSON(200, task)
}

// POST /tasks/:id/done
// タスク完了
func (s *Server) done(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(400)
		return
	}

	task, err := s.interactor.DoneTask(id)
	if err != nil {
		c.Status(404)
		return
	}

	c.JSON(200, task)
}

// POST /tasks/alldone
// タスク完了
func (s *Server) allDone(c *gin.Context) {
	tasks, err := s.interactor.AllDoneTask()
	if err != nil {
		c.Status(404)
		return
	}

	c.JSON(200, tasks)
}
