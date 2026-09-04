package main

import (
	"fmt"
	"net/http"

	"github.com/itapurba0/strata/services/api/internal/database"
	"github.com/itapurba0/strata/services/api/internal/organization"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"STARTA API is Live and Healthy...")
}

func main() {
	db, err := database.Connect()
	if err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}
	defer db.Close()

	organizationRepository := organization.NewRepository(db)
	organizationService := organization.NewService(organizationRepository)
	organizationHandler := organization.NewHandler(organizationService)

	http.HandleFunc("/health", healthHandler)

	http.HandleFunc("/api/v1/organizations", organizationHandler.Create)

	fmt.Println("STRATA API running on http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}