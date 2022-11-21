/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/9/25 2:37 下午
 * @Desc: TODO
 */

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type serverInfo struct {
	Addr    string `json:"addr"`    // 服务器地址
	Version int    `json:"version"` // 客户端版本
}

func main() {
	r := mux.NewRouter()

	g := &game{}

	// 获取服务器信息
	r.HandleFunc("/server-info", g.serverInfo).Methods(http.MethodGet, http.MethodOptions)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("http serve failed: ", err)
	}
}

type game struct{}

// 服务器信息
func (game) serverInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&serverInfo{Addr: "127.0.0.1:3553", Version: 20161227})
}
