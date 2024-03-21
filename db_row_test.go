package db

import (
	"testing"

	db_errors "github.com/fatihtatoglu/db-go/error"
)

func TestSetItemByIndex(t *testing.T) {
	// Arrange
	expectedValues := []interface{}{"Fatih", "Tatoğlu", 39}
	columnNames := []string{"firstname", "lastname", "age"}

	dt := createTestTableWithColumns(columnNames)

	sut := dt.CreateNewDBRow()

	// Act
	for i, value := range expectedValues {
		err := sut.SetItemByIndex(i, value)
		if err != nil {
			t.Errorf("Unexpected error occurred while setting item by index: %v", err)
		}
	}

	// Assert
	assertRowValues(t, sut, expectedValues)
}

func TestSetItemByIndex_IndexOutOfRange(t *testing.T) {
	// Test cases
	testCases := []struct {
		testName       string
		index          int
		expectedErrMsg string
	}{
		{"Negative index", -1, db_errors.Column_IndexOutOfRangeErrorMessage},
		{"Index greater than column count", 99, db_errors.Column_IndexOutOfRangeErrorMessage},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			// Arrange
			columnNames := []string{"firstname", "lastname", "age"}
			dt := createTestTableWithColumns(columnNames)
			sut := dt.CreateNewDBRow()

			// Act
			err := sut.SetItemByIndex(tc.index, "test")

			// Assert
			assertError(t, err, tc.expectedErrMsg)
		})
	}
}

func TestGetItemByIndex(t *testing.T) {
	// Arrange
	columnNames := []string{"firstname", "lastname", "age"}
	expectedValues := []interface{}{"Fatih", "Tatoğlu", 39}

	dt := createTestTableWithColumns(columnNames)

	row := dt.CreateNewDBRow()
	row.itemArray = expectedValues

	// Act & Assert
	assertRowValues(t, row, expectedValues)
}

func TestGetItemByIndex_IndexOutOfRange(t *testing.T) {
	// Test cases
	testCases := []struct {
		testName       string
		index          int
		expectedErrMsg string
	}{
		{"Negative index", -1, db_errors.Column_IndexOutOfRangeErrorMessage},
		{"Index greater than column count", 99, db_errors.Column_IndexOutOfRangeErrorMessage},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			// Arrange
			columnNames := []string{"firstname", "lastname", "age"}
			expectedValues := []interface{}{"Fatih", "Tatoğlu", 39}

			dt := createTestTableWithColumns(columnNames)
			row := dt.CreateNewDBRow()
			row.itemArray = expectedValues

			// Act
			_, err := row.GetItemByIndex(tc.index)

			// Assert
			assertError(t, err, tc.expectedErrMsg)
		})
	}
}

func TestSetItemByName(t *testing.T) {
	// Arrange
	expectedValues := []interface{}{"Fatih", "Tatoğlu", 39}
	columnNames := []string{"firstname", "lastname", "age"}

	dt := createTestTableWithColumns(columnNames)

	sut := dt.CreateNewDBRow()

	// Act
	for i, value := range expectedValues {
		columnName := columnNames[i]
		err := sut.SetItemByName(columnName, value)
		if err != nil {
			t.Errorf("Unexpected error occurred while setting item by index: %v", err)
		}
	}

	// Assert
	assertRowValues(t, sut, expectedValues)
}

func TestSetItemByName_NotExisting_Name(t *testing.T) {
	// Arrange
	columnNames := []string{"firstname", "lastname", "age"}

	dt := createTestTableWithColumns(columnNames)

	sut := dt.CreateNewDBRow()

	// Act
	err := sut.SetItemByName("not-existing-column-name", "test")

	// Assert
	assertError(t, err, db_errors.Column_NotFoundErrorMessage)
}

func TestGetItemByName(t *testing.T) {
	// Arrange
	columnNames := []string{"firstname", "lastname", "age"}
	expectedValues := []interface{}{"Fatih", "Tatoğlu", 39}

	dt := createTestTableWithColumns(columnNames)

	row := dt.CreateNewDBRow()
	row.itemArray = expectedValues

	// Act & Assert
	assertRowValuesByName(t, row, columnNames, expectedValues)
}

func TestGetItemByName_NotExisting_Name(t *testing.T) {
	// Arrange
	columnNames := []string{"firstname", "lastname", "age"}
	expectedValues := []interface{}{"Fatih", "Tatoğlu", 39}

	dt := createTestTableWithColumns(columnNames)

	row := dt.CreateNewDBRow()
	row.itemArray = expectedValues

	// Act
	_, err := row.GetItemByName("not-existing-column-name")

	// Assert
	assertError(t, err, db_errors.Column_NotFoundErrorMessage)
}

func createTestTableWithColumns(columnNames []string) *DBTable {
	dt := CreateNewDBTable()

	for i, name := range columnNames {
		col := CreateNewDBColumn(name, i)
		dt.AddDBColumn(col)
	}

	return &dt
}

func assertRowValues(t *testing.T, row DBRow, expectedValues []interface{}) {
	for i, expectedValue := range expectedValues {
		result, err := row.GetItemByIndex(i)
		if err != nil {
			t.Errorf("Unexpected error occurred while getting item by index: %v", err)
		}

		if result != expectedValue {
			t.Errorf("For index %d, expected value: %v, got value: %v", i, expectedValue, result)
		}
	}
}

func assertRowValuesByName(t *testing.T, row DBRow, columnNames []string, expectedValues []interface{}) {
	for i, expectedValue := range expectedValues {
		result, err := row.GetItemByName(columnNames[i])
		if err != nil {
			t.Errorf("Unexpected error occurred while getting item by index: %v", err)
		}

		if result != expectedValue {
			t.Errorf("For index %d, expected value: %v, got value: %v", i, expectedValue, result)
		}
	}
}

func assertError(t *testing.T, err error, expectedErrorMessage string) {
	if err == nil {
		t.Error("Expected error not occurred")
	}

	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message: %s, got: %s", expectedErrorMessage, err.Error())
	}
}
