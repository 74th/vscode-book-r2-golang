package server

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/74th/vscode-book-r2-golang/domain/entity"
	"github.com/74th/vscode-book-r2-golang/domain/usecase"
	"github.com/74th/vscode-book-r2-golang/repository"
)

// Server サーバAPI
type Server struct {
	server     http.Server
	interactor usecase.Interactor
}

// New サーバAPIのインスタンスを作成する
func New(addr string, webroot string) *Server {
	s := &Server{
		server: http.Server{
			Addr: addr,
		},
		interactor: usecase.Interactor{
			Repository: repository.New(),
		},
	}

	s.setRouter(webroot)

	return s
}

// Serve サーバを開始する
func (s *Server) Serve() {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("could not start server: %s", err.Error())
	}
}

func (s *Server) setRouter(webroot string) {
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/tasks", s.list)
	api.POST("/tasks", s.create)
	api.POST("/tasks/:id/done", s.done)

	router.StaticFile("/", filepath.Join(webroot, "index.html"))
	router.Static("/js", filepath.Join(webroot, "js"))
	s.server.Handler = router
}

// list GET /tasks
func (s *Server) list(c *gin.Context) {
	tasks, err := s.interactor.ShowTasks()
	if err != nil {
		log.Print("error", err)
		writeErrorResponse(c)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// create POST /tasks
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
		writeErrorResponse(c)
		return
	}

	c.JSON(200, task)
}

// done POST /tasks/:id/done
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

func writeErrorResponse(c *gin.Context) {
	c.Status(500)
}
