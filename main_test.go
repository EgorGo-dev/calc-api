package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculate_AllOperations(t *testing.T) {
	tests := []struct {
		a, b int
		op   string
		want int
	}{
		{5, 3, "add", 8},
		{5, 3, "sub", 2},
		{5, 3, "mul", 15},
		{6, 3, "div", 2},
		{-5, -3, "add", -8},
		{5, 0, "mul", 0},
	}

	for _, tt := range tests {
		got, err := calculate(tt.a, tt.b, tt.op)
		if err != nil {
			t.Fatalf("неожиданная ошибка: %v", err)
		}
		if got != tt.want {
			t.Errorf("calculate(%d, %d, %q) = %d; ожидалось %d",
				tt.a, tt.b, tt.op, got, tt.want)
		}
	}
}

func TestCalculate_DivisionByZero(t *testing.T) {
	_, err := calculate(5, 0, "div")
	if err == nil {
		t.Fatal("ожидалась ошибка при делении на ноль")
	}
	if err.Error() != "деление на ноль" {
		t.Errorf("неверный текст ошибки: %q", err.Error())
	}
}

func TestCalculate_UnknownOperation(t *testing.T) {
	_, err := calculate(5, 3, "mod")
	if err == nil {
		t.Fatal("ожидалась ошибка для неизвестной операции")
	}
	if err.Error() != "неизвестная операция" {
		t.Errorf("неверный текст ошибки: %q", err.Error())
	}
}

func TestCalcHandler_ValidRequest(t *testing.T) {
	body := `{"a":5,"b":3,"op":"add"}`
	req := httptest.NewRequest(http.MethodPost, "/calc", strings.NewReader(body))
	w := httptest.NewRecorder()

	calcHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d", w.Code)
	}

	var resp CalcResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("не удалось распарсить ответ: %v", err)
	}
	if resp.Result != 8 {
		t.Errorf("результат = %d; ожидалось 8", resp.Result)
	}
}

func TestCalcHandler_DivisionByZero(t *testing.T) {
	body := `{"a":5,"b":0,"op":"div"}`
	req := httptest.NewRequest(http.MethodPost, "/calc", strings.NewReader(body))
	w := httptest.NewRecorder()

	calcHandler(w, req)

	var resp CalcResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("не удалось распарсить ответ: %v", err)
	}
	if resp.Error != "деление на ноль" {
		t.Errorf("ожидалась ошибка деления на ноль, получено: %q", resp.Error)
	}
}

func TestCalcHandler_BadJSON(t *testing.T) {
	body := `{это не json}`
	req := httptest.NewRequest(http.MethodPost, "/calc", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	calcHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("ожидался статус 400, получен %d", w.Code)
	}
}

func TestCalcHandler_WrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/calc", nil)
	w := httptest.NewRecorder()

	calcHandler(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("ожидался статус 405, получен %d", w.Code)
	}
}