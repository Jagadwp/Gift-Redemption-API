# Testing Documentation

## Overview
This project includes unit tests for the service layer, focusing on business logic validation and edge case handling.

## Test Coverage

### Service Layer Tests

- `auth_service_test.go`: login with valid credentials
- `auth_service_test.go`: login with wrong password
- `auth_service_test.go`: login with non-existent user
- `user_service_test.go`: get user by ID (success & not found)
- `user_service_test.go`: create user (success & duplicate email)
- `user_service_test.go`: update user
- `user_service_test.go`: delete user (success & not found)
- `gift_service_test.go`: get all gifts with pagination
- `gift_service_test.go`: get gift by ID (success & not found)
- `gift_service_test.go`: create gift
- `gift_service_test.go`: patch gift (partial update)
- `gift_service_test.go`: star rating rounding logic (table-driven tests)
- `redemption_service_test.go`: gift not found validation
- `redemption_service_test.go`: rate gift validation (not redeemed, gift not found)
- `redemption_service_test.go`: score validation (1-5 range)
- `redemption_service_test.go`: note on transaction logic (stock deduction, rating stats update) requires integration tests with real DB

## Running Tests

### Run all tests
```bash
make test
```

### Run with coverage
```bash
make test-coverage
```
This will generate `coverage.html` that you can open in browser.

### Run specific test
```bash
go test -v ./internal/service -run TestAuthService_Login_Success
```

### Run with verbose output
```bash
go test -v ./internal/service/...
```

## Test Architecture

### Mocking Strategy
We use manual mocks (not code generation) for simplicity and control. Mock repositories are located in `internal/repository/mocks/`.

**Advantages:**

- No external dependencies for mock generation
- Easy to understand and modify
- Full control over mock behavior

### Test Structure
Each test follows this pattern:
1. **Arrange** - Setup mocks and test data
2. **Act** - Call the function being tested
3. **Assert** - Verify results and mock calls

Example:
```go
func TestGiftService_GetByID_Success(t *testing.T) {
    // Arrange
    mockGiftRepo := new(mocks.MockGiftRepository)
    giftService := NewGiftService(mockGiftRepo)
    
    gift := &model.Gift{ID: 1, Name: "Test"}
    mockGiftRepo.On("FindByID", uint(1)).Return(gift, nil)
    
    // Act
    result, err := giftService.GetByID(1)
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, "Test", result.Name)
    mockGiftRepo.AssertExpectations(t)
}
```

## Key Test Cases

### Star Rating Rounding
Tests verify that `avg_rating` is correctly rounded to nearest 0.5:
- 3.2 → 3.0
- 3.6 → 3.5
- 3.9 → 4.0

### Error Handling
All service tests verify proper error propagation from repository layer:
- `ErrNotFound` - Resource doesn't exist
- `ErrDuplicateEntry` - Unique constraint violation
- `ErrInsufficientStock` - Stock validation
- `ErrNotRedeemed` - Rating validation

## Limitations

### Transaction Testing
Redemption and rating services use database transactions which cannot be easily mocked:
- **Current approach:** Test only validation logic before transactions
- **Why:** `gorm.DB.Transaction()` requires actual database connection
- **Solution for full coverage:** 
  - Integration tests with testcontainers (real PostgreSQL)
  - Or use sqlmock for mocking database queries

**What's tested:**

- ✅ Pre-transaction validation (gift exists, user has redeemed)
- ⏭️ Stock deduction logic (requires integration test)
- ⏭️ Rating stats update (requires integration test)

### Coverage Goals
Current coverage focuses on:

- ✅ Business logic validation
- ✅ Error handling paths
- ✅ Edge cases (rounding, validation)
- ⏭️ Transaction logic (future: integration tests)
- ⏭️ Concurrent redemption scenarios (future: load tests)

## Future Improvements
1. Add integration tests with testcontainers
2. Add handler layer tests (HTTP request/response)
3. Add repository layer tests with sqlmock
4. Implement CI/CD pipeline with automated testing
