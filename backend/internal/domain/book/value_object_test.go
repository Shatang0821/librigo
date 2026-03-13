package book

import (
	"errors"
	"testing"
)

func TestNewBookTitle(t *testing.T) {
	tests := map[string]struct {
		input   string
		wantErr error
	}{
		"正常系": {
			input:   "Go言語",
			wantErr: nil,
		},
		"異常系: 空文字": {
			input:   "",
			wantErr: ErrInvalidBookTitle,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := NewBookTitle(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewBookTitle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewBookPrice(t *testing.T) {
	tests := map[string]struct {
		input   int
		wantErr error
	}{
		"正常系": {
			input:   1000,
			wantErr: nil,
		},
		"異常系: マイナス": {
			input:   -1,
			wantErr: ErrInvalidBookPrice,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := NewBookPrice(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewBookPrice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewBookISBN(t *testing.T) {
	tests := map[string]struct {
		input   string
		wantErr error
	}{
		"正常系": {
			input:   "978-4-0000-0000-0",
			wantErr: nil,
		},
		"異常系: 空文字": {
			input:   "",
			wantErr: ErrInvalidBookISBN,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := NewBookISBN(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewBookISBN() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
