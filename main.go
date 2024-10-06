package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

// Sorting algorithms
func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func insertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	left := []int{}
	right := []int{}
	for _, val := range arr[1:] {
		if val <= pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}
	return append(quickSort(left), append([]int{pivot}, quickSort(right)...)...)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}
	return result
}

// Request handler for sorting
func sortHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse input
	var requestBody struct {
		Array     []int  `json:"array"`
		Algorithm string `json:"algorithm"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var sortedArray []int

	switch requestBody.Algorithm {
	case "bubble":
		sortedArray = bubbleSort(requestBody.Array)
	case "selection":
		sortedArray = selectionSort(requestBody.Array)
	case "insertion":
		sortedArray = insertionSort(requestBody.Array)
	case "quick":
		sortedArray = quickSort(requestBody.Array)
	case "merge":
		sortedArray = mergeSort(requestBody.Array)
	default:
		http.Error(w, "Invalid sorting algorithm", http.StatusBadRequest)
		return
	}

	// Return sorted array as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sortedArray)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/sort", sortHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, nil)
}
