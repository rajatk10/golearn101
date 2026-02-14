package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

// start with generate short url test, no dependency

func TestGenerateShortURLLength(t *testing.T) {
	result := generateShortURL()
	if len(result) != 6 {
		t.Errorf("Expected length 6, got %d ", len(result))
	}
}

func TestGenerateShortURLReturnType(t *testing.T) {
	short := generateShortURL()
	if reflect.TypeOf(short).Kind() != reflect.String {
		t.Errorf("Return Type Should be string")
	}
}

func TestUniqueShortURLUniqueness(t *testing.T) {
	seen := make(map[string]bool)
	for _ = range 50 { //trying 50 times for uniqueness check
		short := generateShortURL()
		if _, exists := seen[short]; exists {
			t.Errorf("Short URL %s is not unique", short)
		}
		seen[short] = true
	}
}

//Test shorthandler method/function

func TestShortenHandlerValid(t *testing.T) {
	body := `{"url": "https://github.com/ramkt10/golearn10"}`
	req := httptest.NewRequest(http.MethodPost, "/shorten", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	shortenHandler(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
	if !strings.Contains(w.Body.String(), "short_url") {
		t.Errorf("Response should contain short_url")
	}
}

func TestShortenHandlerInvalidHTTPMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/shorten", nil)
	w := httptest.NewRecorder()
	shortenHandler(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, w.Code)
	}
}

func TestShortenHandlerInvalidJSON(t *testing.T) {
	body := `{invalid json}`
	req := httptest.NewRequest(http.MethodPost, "/shorten", strings.NewReader(body))
	w := httptest.NewRecorder()
	shortenHandler(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestShortenInvalidURL(t *testing.T) {
	body := `{"url": "invalid-url"}`
	req := httptest.NewRequest(http.MethodPost, "/shorten", strings.NewReader(body))
	w := httptest.NewRecorder()
	shortenHandler(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}

// Now test redirect handler
func TestRedirectHandlerValid(t *testing.T) {
	mu.Lock()
	urlstore["test123"] = "https://github.com/testorg/golearn101"
	mu.Unlock()
	req := httptest.NewRequest(http.MethodGet, "/sh/test123", nil)
	w := httptest.NewRecorder()
	redirectHandler(w, req)
	if w.Code != http.StatusFound {
		t.Errorf("Expected status code %d, got %d", http.StatusFound, w.Code)
	}
	//check if redirect url worked
	location := w.Header().Get("Location")
	if location != "https://github.com/testorg/golearn101" {
		t.Errorf("Expected location %s, got %s", "https://github.com/testorg/golearn101", location)
	}
	mu.Lock()
	delete(urlstore, "test123")
	mu.Unlock()
}

func TestRedirectHandlerInvalid(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/sh/testInvalid123", nil)
	w := httptest.NewRecorder()
	redirectHandler(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestRedirectHandlerEmptyCode(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/sh/", nil)
	w := httptest.NewRecorder()
	redirectHandler(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestSaveURLShortStoreFileValid(t *testing.T) {
	originalFile := urlShortFile
	urlShortFile = "test_urlstore.json"
	defer func() {
		urlShortFile = originalFile
		os.Remove("test_urlstore.json")
	}()
	mu.Lock()
	urlstore = make(map[string]string)
	urlstore["test123"] = "https://github.com/testorg/golearn101"
	mu.Unlock()
	if err := SaveURLShortStoreFile(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if _, err := os.Stat("test_urlstore.json"); os.IsNotExist(err) {
		t.Errorf("Expected file to exist, got %v", err)
	}
	data, err := os.ReadFile("test_urlstore.json")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !strings.Contains(string(data), "test123") {
		t.Errorf("Expected file content to contain test123, got %s", string(data))
	}
	mu.Lock()
	delete(urlstore, "test123")
	mu.Unlock()

}
