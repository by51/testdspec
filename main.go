package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"testdespec/mathutils"
	"testdespec/stringutils"
	"time"
)

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Items []string `json:"items"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": strings.Join(req.Items, "")})
}

func varianceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Items []float64 `json:"items"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	v, err := mathutils.Variance(req.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"variance": v})
}

func pingResponse(hour int) string {
	if hour >= 10 && hour < 18 {
		return "pong"
	}
	return "pang"
}

func randomNumberHandler(w http.ResponseWriter, r *http.Request) {
	n := rand.Int63n(1_000_000_000_000)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%012d", n)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, HTTP Server!")
	})
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, pingResponse(time.Now().Hour()))
	})
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "你好")
	})
	mux.HandleFunc("/add", addHandler)
	mux.HandleFunc("/variance", varianceHandler)
	mux.HandleFunc("/random", randomNumberHandler)
	mux.HandleFunc("/split", func(w http.ResponseWriter, r *http.Request) {
		s := r.URL.Query().Get("s")
		parts := stringutils.SplitBySpace(s)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string][]string{"parts": parts})
	})

	server := &http.Server{
		Addr:    ":9999",
		Handler: mux,
	}

	// 监听 SIGINT / SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Println("服务启动，监听 :9999")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务异常退出: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("收到退出信号，开始优雅关闭...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("优雅关闭失败: %v", err)
	}
	log.Println("服务已停止")
}
