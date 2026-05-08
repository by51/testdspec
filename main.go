package main

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
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

// TimestampResponse 时间戳接口响应结构
type TimestampResponse struct {
	Unix     int64  `json:"unix"`
	UnixMs   int64  `json:"unix_ms"` // 毫秒级时间戳
	Readable string `json:"readable"`
}

// TimezoneResponse 时区接口响应结构
type TimezoneResponse struct {
	Timezone string `json:"timezone"` // 时区名称，如 Asia/Shanghai
	Offset   string `json:"offset"`   // 时区偏移，如 +0800
}

// Md5Request MD5 接口请求结构
type Md5Request struct {
	Input string `json:"input"` // 待加密的字符串
}

// Md5Response MD5 接口响应结构
type Md5Response struct {
	Hash string `json:"hash"` // MD5 哈希值（32位小写十六进制）
}

// Sha256Request SHA256 接口请求结构
type Sha256Request struct {
	Input string `json:"input"` // 待加密的字符串
}

// Sha256Response SHA256 接口响应结构
type Sha256Response struct {
	Hash string `json:"hash"` // SHA256 哈希值（64位小写十六进制）
}

// Sha1Request SHA1 接口请求结构
type Sha1Request struct {
	Input string `json:"input"` // 待加密的字符串
}

// Sha1Response SHA1 接口响应结构
type Sha1Response struct {
	Hash string `json:"hash"` // SHA1 哈希值（40位小写十六进制）
}

func timestampHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	now := time.Now()
	resp := TimestampResponse{
		Unix:     now.Unix(),
		UnixMs:   now.UnixMilli(),
		Readable: now.Format(time.RFC3339),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func timezoneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	name, offset := time.Now().Zone()
	// 将秒数转换为 +HHMM 格式
	sign := "+"
	if offset < 0 {
		sign = "-"
		offset = -offset
	}
	hours := offset / 3600
	minutes := (offset % 3600) / 60
	offsetStr := fmt.Sprintf("%s%02d%02d", sign, hours, minutes)

	resp := TimezoneResponse{
		Timezone: name,
		Offset:   offsetStr,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func md5Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req Md5Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	// 计算 MD5 哈希
	hash := md5.Sum([]byte(req.Input))
	resp := Md5Response{
		Hash: hex.EncodeToString(hash[:]),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func sha256Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req Sha256Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	// 计算 SHA256 哈希
	hash := sha256.Sum256([]byte(req.Input))
	resp := Sha256Response{
		Hash: hex.EncodeToString(hash[:]),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func sha1Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req Sha1Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	// 计算 SHA1 哈希
	// 注意：SHA1 不适用于安全敏感场景（如密码存储、数字签名），仅用于兼容性和校验场景
	hash := sha1.Sum([]byte(req.Input))
	resp := Sha1Response{
		Hash: hex.EncodeToString(hash[:]),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
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
	mux.HandleFunc("/timestamp", timestampHandler)
	mux.HandleFunc("/timezone", timezoneHandler)
	mux.HandleFunc("/md5", md5Handler)
	mux.HandleFunc("/sha256", sha256Handler)
	mux.HandleFunc("/sha1", sha1Handler)

	server := &http.Server{
		Addr:    ":10001",
		Handler: mux,
	}

	// 监听 SIGINT / SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Println("服务启动，监听 :10001")
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
