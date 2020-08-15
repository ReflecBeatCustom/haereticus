package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"sync/atomic"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	// import not used
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"

	"github.com/ReflecBeatCustom/haereticus/pkg/config"
	"github.com/ReflecBeatCustom/haereticus/pkg/types"
)

const (
	// MethodGetFumens [...]
	MethodGetFumens = "GetFumens"
	// MethodGetPacks [...]
	MethodGetPacks = "GetPacks"
)

// HaereticusServer ...
type HaereticusServer struct {
	httpServer *http.Server
	healthy    int32
	config     *types.ServerConfig
	dbClient   *gorm.DB
}

// NewServer ...
func NewServer(configFile string) *HaereticusServer {
	server := &HaereticusServer{}
	router := chi.NewRouter()
	router.Group(func(router chi.Router) {
		router.Use(middleware.RequestID)
		router.Use(middleware.RealIP)
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)
		router.Mount("/debug", middleware.Profiler())
		router.Get("/api/healthz", server.Healthz)
		router.Post("/api", server.Root)
	})

	config, err := config.NewServerConfig(configFile)
	if err != nil {
		glog.Errorf("Generate config from file error: %v", err)
	}

	httpServer := &http.Server{
		Addr:         config.Server.ListenAddr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	dbClient, err := gorm.Open("mysql", filepath.Join(config.DB.DBAddr, config.DB.DBName)+"?charset=utf8mb4&parseTime=True&loc=Asia%2fShanghai")
	dbClient.DB().SetConnMaxLifetime(500)
	if err != nil {
		glog.Errorf("Open mysql DB error: %v", err)
		return nil
	}
	err = dbClient.DB().Ping()
	if err != nil {
		glog.Errorf("Connect to mysql DB error: %v", err)
		return nil
	}
	dbClient.SingularTable(true)

	server.httpServer = httpServer
	server.config = config
	server.dbClient = dbClient
	return server
}

// Start ...
func (server *HaereticusServer) Start() {
	glog.Infof("Server start")
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		glog.Infof("Server is shutting down..")
		atomic.StoreInt32(&server.healthy, 0)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		server.httpServer.SetKeepAlivesEnabled(false)
		if err := server.httpServer.Shutdown(ctx); err != nil {
		}
		close(done)
	}()
	glog.Infof("Server is ready to handle requests at %s", server.httpServer.Addr)
	atomic.StoreInt32(&server.healthy, 1)
	err := server.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		glog.Errorf("Failed to listen on %s: %v", server.httpServer.Addr, err)
	}
	<-done
	glog.Infof("Server closed")
}

// Healthz ...
func (server *HaereticusServer) Healthz(w http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	if atomic.LoadInt32(&server.healthy) == 1 {
		response := types.HealthzResponse{
			Result: true,
		}
		responseBytes, err := json.Marshal(response)
		if err != nil {
			glog.Errorf("Marshal response error: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}
		glog.Infof("Healthz response: %+v", response)
		responseOK(w, responseBytes)
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}

// Root ...
func (server *HaereticusServer) Root(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		glog.Errorf("Read body error: %v", err)
		return
	}
	glog.Infof("Receive request: %s", body)
	defer request.Body.Close()

	var header types.Header
	if err := json.Unmarshal(body, &header); err != nil {
		glog.Errorf("Decode request body failed: %v", err)
		return
	}
	glog.Infof("Request method: %s", header.Method)

	switch header.Method {
	case MethodGetFumens:
		server.GetFumens(writer, body)
	case MethodGetPacks:
		server.GetPacks(writer, body)
	}

	return
}

func responseOK(writer http.ResponseWriter, responseBytes []byte) {
	writer.Write(responseBytes)
	writer.Header().Set("Content-Type", "text/json; charset=utf-8")
	writer.Header().Set("X-Content-Type-Options", "nosniff")
	writer.WriteHeader(http.StatusOK)
}
