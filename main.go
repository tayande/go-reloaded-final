// package main

// import (
// 	"fmt"
// 	"os"
// )
// func main() {
// 	if len(os.Args) != 3 {
// 		fmt.Fprintln(os.Stderr, "Usage: go run . input.txt output.txt")
// 		os.Exit(2)
// 	}
// 	inputFile := os.Args[1]
// 	outputFile := os.Args[2]

// 	if inputFile == outputFile {
// 		fmt.Fprintln(os.Stderr, "Cannot use the same input file as output file")
// 		os.Exit(1)
// 	}

// 	info, err := os.ReadFile(inputFile)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, "Error. File does not exist")
// 		os.Exit(1)
// 	}
// 	if len(info) == 0 {
// 		fmt.Fprintln(os.Stderr, "File is empty")
// 		os.Exit(1)
// 	}
// 	data := string(info)
// 	processed := processor(data)
// 	newProcessed := []byte(processed)

// 	err = os.WriteFile(outputFile, newProcessed, 0644)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, "Error writing into output file")
// 		os.Exit(1)
// 	}
// 	fmt.Fprintf(os.Stdout, "Done. Output written into %s\n", outputFile)

// }


package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	// CLI mode: go run . input.txt output.txt
	if len(os.Args) == 3 {
		runCLI(os.Args[1], os.Args[2])
		return
	}

	// Web mode: go run .
	runServer()
}

// CLI mode

func runCLI(inputFile, outputFile string) {
	if inputFile == outputFile {
		fmt.Fprintln(os.Stderr, "Cannot use the same input file as output file")
		os.Exit(1)
	}

	info, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error. File does not exist")
		os.Exit(1)
	}
	if len(info) == 0 {
		fmt.Fprintln(os.Stderr, "File is empty")
		os.Exit(1)
	}

	data := string(info)
	processed := processor(data)

	err = os.WriteFile(outputFile, []byte(processed), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error writing into output file")
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Done. Output written into %s\n", outputFile)
}

// Web mode

type processRequest struct {
	Text string `json:"text"`
}

type processResponse struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func runServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "static/index.html")
	})
	mux.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/style.css")
	})
	mux.HandleFunc("/process", handleProcess)

	port := "8080"
	fmt.Printf("Server running. Open http://localhost:%s in your browser.\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Fprintln(os.Stderr, "Server error:", err)
		os.Exit(1)
	}
}

func handleProcess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(processResponse{Error: "Method not allowed"})
		return
	}

	var req processRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(processResponse{Error: "Invalid request body"})
		return
	}
	if strings.TrimSpace(req.Text) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(processResponse{Error: "Input text is empty"})
		return
	}

	result := processor(req.Text)

	// Also write to output.txt
	// at the same time it's shown on the web page.
	if err := os.WriteFile("output.txt", []byte(result), 0644); err != nil {
		fmt.Fprintln(os.Stderr, "Warning: could not write output.txt:", err)
	}

	json.NewEncoder(w).Encode(processResponse{Result: result})
}
