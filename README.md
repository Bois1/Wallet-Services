# Wallet Transfer Service Test Code

## My Approach

### Dependency Injection
- Defined `WalletRepository` interface for data access.
- `WalletService` depends only on this interface.
- Implemented two repos: `InMemoryWalletRepo` (for tests) and `FailingWalletRepo` (to show dependency injection flexibility).
- Service is fully decoupled from storage implementation.

### Money Handling
- Used `int64` to represent **cents** (smallest currency unit).
- Wrapped in `Money` type with validation (positive amounts only).
- Avoids floating-point arithmetic entirely → no rounding errors.

### Concurrency & Safety
- Global mutex ensures **atomic** transfers (no race conditions).
- **Idempotency** via transfer ID tracking to prevent duplicate processing.
- Wallets are copied on load/save to avoid shared mutation.

## Running Tests
`go test service_test.go`