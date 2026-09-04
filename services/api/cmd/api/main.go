package main

import (
	"fmt"
	"net/http"

	"github.com/itapurba0/strata/services/api/internal/config"
	"github.com/itapurba0/strata/services/api/internal/database"
	"github.com/itapurba0/strata/services/api/internal/organization"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "STARTA API is Live and Healthy...")
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Configuration error:", err)
		return
	}

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}
	defer db.Close()

	organizationRepository := organization.NewRepository(db)
	organizationService := organization.NewService(organizationRepository)
	organizationHandler := organization.NewHandler(organizationService)

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("POST /api/v1/organizations", organizationHandler.Create)
	http.HandleFunc("GET /api/v1/organizations/{id}", organizationHandler.GetByID)

	fmt.Println("STRATA API running on http://localhost:" + cfg.Port)

	err = http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
