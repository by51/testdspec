package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"syscall"
	"testing"
	"time"
	"unicode"
)

func TestAddHandlerIntegration(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", addHandler)
	srv := httptest.NewServer(mux)
	defer srv.Close()

	resp, err := http.Post(srv.URL+"/add", "application/json", strings.NewReader(`{"items":["foo","bar","baz"]}`))
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("期望状态码 200，实际: %d", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	if !strings.Contains(string(body), "foobarbaz") {
		t.Fatalf("期望响应包含 foobarbaz，实际: %s", body)
	}
}

func TestGracefulShutdown(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	server := &http.Server{
		Addr:    ":19999",
		Handler: mux,
	}

	started := make(chan struct{})
	done := make(chan error, 1)

	go func() {
		close(started)
		err := server.ListenAndServe()
		done <- err
	}()

	<-started
	time.Sleep(20 * time.Millisecond) // 等待端口就绪

	// 验证服务已启动
	resp, err := http.Get("http://localhost:19999/")
	if err != nil {
		t.Fatalf("服务未正常启动: %v", err)
	}
	resp.Body.Close()

	// 模拟收到 SIGINT 后的关闭流程
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		t.Fatalf("Shutdown 返回错误: %v", err)
	}

	// ListenAndServe 应返回 ErrServerClosed
	select {
	case err := <-done:
		if err != http.ErrServerClosed {
			t.Fatalf("期望 ErrServerClosed，实际: %v", err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("服务未在预期时间内停止")
	}

	// 验证关闭后无法建立新连接
	_, err = http.Get("http://localhost:19999/")
	if err == nil {
		t.Fatal("服务关闭后仍能接受连接")
	}

	_ = syscall.SIGINT // 确认 signal 包正确引用
}

func TestTestHandler(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "你好")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("期望状态码 200，实际: %d", w.Code)
	}

	body, _ := io.ReadAll(w.Body)
	if string(body) != "你好\n" {
		t.Fatalf("期望响应体 '你好\\n'，实际: %q", string(body))
	}
}

func TestAddHandler(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantResult string
	}{
		{
			name:       "多字符串拼接",
			body:       `{"items":["hello"," ","world"]}`,
			wantStatus: http.StatusOK,
			wantResult: "hello world",
		},
		{
			name:       "单字符串",
			body:       `{"items":["only"]}`,
			wantStatus: http.StatusOK,
			wantResult: "only",
		},
		{
			name:       "空数组",
			body:       `{"items":[]}`,
			wantStatus: http.StatusOK,
			wantResult: "",
		},
		{
			name:       "错误请求体",
			body:       `not-json`,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader(tc.body))
			w := httptest.NewRecorder()
			addHandler(w, req)

			if w.Code != tc.wantStatus {
				t.Fatalf("期望状态码 %d，实际: %d", tc.wantStatus, w.Code)
			}
			if tc.wantStatus == http.StatusOK {
				body, _ := io.ReadAll(w.Body)
				// JSON 响应包含 result 字段
				if !strings.Contains(string(body), tc.wantResult) {
					t.Fatalf("期望结果包含 %q，实际响应: %s", tc.wantResult, body)
				}
			}
		})
	}
}

func TestVarianceHandlerIntegration(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/variance", varianceHandler)
	srv := httptest.NewServer(mux)
	defer srv.Close()

	// 正常计算
	resp, err := http.Post(srv.URL+"/variance", "application/json", strings.NewReader(`{"items":[2,4,4,4,5,5,7,9]}`))
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("期望 200，实际: %d", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	if !strings.Contains(string(body), "4") {
		t.Fatalf("期望方差为 4，实际响应: %s", body)
	}

	// 空列表返回 400
	resp2, err := http.Post(srv.URL+"/variance", "application/json", strings.NewReader(`{"items":[]}`))
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp2.Body.Close()
	if resp2.StatusCode != http.StatusBadRequest {
		t.Fatalf("期望 400，实际: %d", resp2.StatusCode)
	}
}

func TestRandomNumberHandler(t *testing.T) {
	t.Run("返回值长度始终为12位", func(t *testing.T) {
		for i := 0; i < 20; i++ {
			req := httptest.NewRequest(http.MethodGet, "/random", nil)
			w := httptest.NewRecorder()
			randomNumberHandler(w, req)

			if w.Code != http.StatusOK {
				t.Fatalf("期望状态码 200，实际: %d", w.Code)
			}
			body, _ := io.ReadAll(w.Body)
			s := string(body)
			if len(s) != 12 {
				t.Fatalf("第 %d 次调用：期望长度 12，实际长度 %d，值: %q", i+1, len(s), s)
			}
		}
	})

	t.Run("返回值为纯数字", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/random", nil)
		w := httptest.NewRecorder()
		randomNumberHandler(w, req)

		body, _ := io.ReadAll(w.Body)
		for _, c := range string(body) {
			if !unicode.IsDigit(c) {
				t.Fatalf("返回值包含非数字字符: %q", string(body))
			}
		}
	})

	t.Run("多次调用大概率不同", func(t *testing.T) {
		results := make(map[string]bool)
		for i := 0; i < 100; i++ {
			req := httptest.NewRequest(http.MethodGet, "/random", nil)
			w := httptest.NewRecorder()
			randomNumberHandler(w, req)
			body, _ := io.ReadAll(w.Body)
			results[string(body)] = true
		}
		// 100 次调用至少应有 50 个不同值
		if len(results) < 50 {
			t.Fatalf("随机性不足：100 次调用只有 %d 个不同值", len(results))
		}
	})
}

func TestRandomNumberHandlerIntegration(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/random", randomNumberHandler)
	srv := httptest.NewServer(mux)
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/random")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("期望状态码 200，实际: %d", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	if len(body) != 12 {
		t.Fatalf("期望响应长度 12，实际: %d，值: %q", len(body), string(body))
	}
}

func TestPingResponse(t *testing.T) {
	tests := []struct {
		hour int
		want string
	}{
		{hour: 10, want: "pong"}, // 工作时段起点（含）
		{hour: 17, want: "pong"}, // 工作时段末尾
		{hour: 9, want: "pang"},  // 工作时段之前
		{hour: 18, want: "pang"}, // 工作时段结束（不含）
	}
	for _, tc := range tests {
		got := pingResponse(tc.hour)
		if got != tc.want {
			t.Errorf("pingResponse(%d) = %q，期望 %q", tc.hour, got, tc.want)
		}
	}
}
