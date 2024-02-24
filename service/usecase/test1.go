package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const APIKey = "qwerklj1230dsa350123l2k1j4kl1j24"

func MiddlewareValidateAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("api-key")
		if apiKey != APIKey {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

type Node struct {
	Value interface{}
	Next  *Node
}

func ArrayToNode(arr []interface{}) *Node {
	var head *Node
	var prev *Node

	for _, val := range arr {
		node := &Node{Value: val}
		if head == nil {
			head = node
		} else {
			prev.Next = node
		}
		prev = node
	}

	return head
}

func Test1(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Array []interface{} `json:"Array"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 轉換陣列
	node := ArrayToNode(requestData.Array)

	curr := node
	index := 0
	for curr != nil {
		// 第一個點為 head
		if index == 0 {
			fmt.Fprintf(w, "head -> %v\n", curr.Value)

			// 移動到下一個節點
			index++
			curr = curr.Next
			continue
		}

		// 紀錄上一個節點
		prev := curr
		// 移動到下一個節點
		curr = curr.Next

		// 最後一個點為 tail
		if curr == nil {
			fmt.Fprintf(w, "tail -> %v\n", prev.Value)
		} else {
			// 中間點
			fmt.Fprintf(w, "node%d -> %v\n", index, prev.Value)
		}
		index++
	}
}
