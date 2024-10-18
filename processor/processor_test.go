package processor_test

import (
	"errors"
	"test-ozon-2/mocks"
	"test-ozon-2/processor"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestProcessDocument(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Arrange
	tests := []struct {
		name        string
		jsonData    []byte
		setupLogger func(logger *mocks.MockLogger)
		setupDB     func(db *mocks.MockDatabase)
		expectedErr error
	}{
		{
			name:     "Успешная обработка",
			jsonData: []byte(`{"header": "Sample Header", "line_items": ["item1", "item2", "item3"]}`),
			setupLogger: func(logger *mocks.MockLogger) {
				logger.EXPECT().LogInfo("Document processed successfully").Times(1)
			},
			setupDB: func(db *mocks.MockDatabase) {
				db.EXPECT().WriteToDatabase(gomock.Any()).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name:     "Неверный формат JSON",
			jsonData: []byte(`{invalid-json}`),
			setupLogger: func(logger *mocks.MockLogger) {
				logger.EXPECT().LogError("Invalid JSON data").Times(1)
			},
			setupDB: func(db *mocks.MockDatabase) {
				// Ничего не ожидается от базы данных, т.к. до неё не дойдёт
			},
			expectedErr: errors.New("invalid character 'i' looking for beginning of object key string"),
		},
		{
			name:     "Ошибка валидации",
			jsonData: []byte(`{"header": "", "line_items": []}`),
			setupLogger: func(logger *mocks.MockLogger) {
				logger.EXPECT().LogError("Missing header or line items").Times(1)
			},
			setupDB: func(db *mocks.MockDatabase) {
				// Ничего не ожидается от базы данных, т.к. до неё не дойдёт
			},
			expectedErr: errors.New("validation error"),
		},
		{
			name:     "Ошибка базы данных",
			jsonData: []byte(`{"header": "Sample Header", "line_items": ["item1", "item2", "item3"]}`),
			setupLogger: func(logger *mocks.MockLogger) {
				logger.EXPECT().LogError("Database write error").Times(1)
			},
			setupDB: func(db *mocks.MockDatabase) {
				db.EXPECT().WriteToDatabase(gomock.Any()).Return(errors.New("database error")).Times(1)
			},
			expectedErr: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLogger := mocks.NewMockLogger(ctrl)
			mockDB := mocks.NewMockDatabase(ctrl)

			tt.setupLogger(mockLogger)
			tt.setupDB(mockDB)

			// Act
			p := processor.New(mockLogger, mockDB)
			err := p.ProcessDocument(tt.jsonData)

			// Assert
			assert.Equal(t, tt.expectedErr, err, "Ожидалась ошибка: %v, но была получена: %v", tt.expectedErr, err)
		})
	}
}
