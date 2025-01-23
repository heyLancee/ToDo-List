package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// ListItem 代表一个列表项
type ListItem struct {
	Value       string `json:"value"`
	IsCompleted bool   `json:"isCompleted"`
}

// Response 定义标准的API响应格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Request struct {
	ID string `json:"id"`
}

// item_map 存储列表项
var item_map map[int]ListItem

var curr_id int

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")                   // 允许所有域名访问
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // 允许的 HTTP 方法
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")       // 允许的请求头
}

func getListHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w) // 启用 CORS
	if r.Method == "OPTIONS" {
		return // 预检请求直接返回
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item_map) // 返回整个 map
}

func addListHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w) // 启用 CORS
	if r.Method == "OPTIONS" {
		return // 预检请求直接返回
	}

	w.Header().Set("Content-Type", "application/json")

	var newItem ListItem
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		response := Response{
			Code:    400,
			Message: "添加失败：" + err.Error(),
			Data:    item_map,
		}
		w.WriteHeader(http.StatusBadRequest) // 设置 HTTP 状态码为 400
		json.NewEncoder(w).Encode(response)
		return
	}

	// 随机生成一个唯一的 ID
	item_map[curr_id] = newItem
	curr_id++

	response := Response{
		Code:    201,
		Message: "任务添加成功！",
		Data:    newItem,
	}
	w.WriteHeader(http.StatusCreated) // 设置 HTTP 状态码为 201
	json.NewEncoder(w).Encode(response)
}

func updateListHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w) // 启用 CORS
	if r.Method == "OPTIONS" {
		return // 预检请求直接返回
	}

	w.Header().Set("Content-Type", "application/json")

	var req Request
	err1 := json.NewDecoder(r.Body).Decode(&req)
	id, err2 := strconv.Atoi(req.ID)
	if err1 != nil || err2 != nil {
		response := Response{
			Code:    400,
			Message: "更新失败",
			Data:    nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if item, exists := item_map[id]; exists {
		item.IsCompleted = !item.IsCompleted // 切换任务状态
		item_map[id] = item                  // 更新 map 中的任务

		response := Response{
			Code:    200,
			Message: "任务状态已更新！",
			Data:    item,
		}
		json.NewEncoder(w).Encode(response)
	} else {
		response := Response{
			Code:    404,
			Message: "未找到该任务！",
			Data:    nil,
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
	}
}

func deleteListHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w) // 启用 CORS
	if r.Method == "OPTIONS" {
		return // 预检请求直接返回
	}
	w.Header().Set("Content-Type", "application/json")

	var req Request
	err1 := json.NewDecoder(r.Body).Decode(&req)
	id, err2 := strconv.Atoi(req.ID)
	if err1 != nil || err2 != nil {
		response := Response{
			Code:    400,
			Message: "删除失",
			Data:    nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if _, exists := item_map[id]; exists {
		delete(item_map, id) // 从 map 中删除任务

		response := Response{
			Code:    200,
			Message: "任务已删除！",
			Data:    nil,
		}
		json.NewEncoder(w).Encode(response)
	} else {
		response := Response{
			Code:    404,
			Message: "未找到该任务！",
			Data:    nil,
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	item_map = make(map[int]ListItem)
	curr_id = 0
	item_map[curr_id] = ListItem{Value: "学习 Go 语言", IsCompleted: false}
	curr_id++

	http.HandleFunc("/api/get_list", getListHandler)
	http.HandleFunc("/api/add_list", addListHandler)
	http.HandleFunc("/api/update_list", updateListHandler)
	http.HandleFunc("/api/delete_list", deleteListHandler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
