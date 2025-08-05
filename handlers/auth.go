package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("my-secret-key") // 建議用 .env 讀取（這裡先寫死）

// LoginRequest 定義登入請求的格式
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler 處理 POST /login 請求
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "❌ 無法解析請求", http.StatusBadRequest)
		return
	}

	// ✅ 模擬帳號密碼驗證（之後會連 AD）
	if req.Username != "admin" || req.Password != "1234" {
		http.Error(w, "❌ 帳號或密碼錯誤", http.StatusUnauthorized)
		return
	}

	// ✅ 建立 JWT token，效期 8 小時
	expiration := time.Now().Add(8 * time.Hour)
	claims := jwt.MapClaims{
		"username": req.Username,
		"exp":      expiration.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "❌ Token 產生失敗", http.StatusInternalServerError)
		return
	}

	// ✅ 回傳 token 給前端
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
