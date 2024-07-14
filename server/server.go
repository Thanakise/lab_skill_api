package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanakize/lab_skill_api/server/router"
	"github.com/thanakize/lab_skill_api/usecase"
)

type Server struct{
	goServer  *http.Server
}

type GoServer interface {
	StartServer()
}

func InitServerGO(usecase usecase.IUseCase) GoServer {
	
	r := gin.Default()
	router.InitRoute(r, usecase)

	srv := http.Server{
		Addr: ":" + os.Getenv("PORT"),
		Handler:  r,
	}
	

	return Server{
		goServer : &srv,
	}
}

func (s Server) StartServer(){
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	go func ()  {
		<-ctx.Done()
		fmt.Println("Shutting down...")
		ctx, cancle := context.WithTimeout(context.Background(), time.Second * 5)
		defer cancle()

		if err := s.goServer.Shutdown(ctx); err != nil {
		if !errors.Is(err, http.ErrServerClosed){
			log.Println(err)
		}
		}

	}()
	if err := s.goServer.ListenAndServe(); err != nil{
		log.Println(err)
	}
}