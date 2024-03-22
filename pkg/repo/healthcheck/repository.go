package healthcheck

// HealthCheckRepository is an interface that defines the methods for checking
// the health of the database.
type Repository interface {
	DBHealth() error
}