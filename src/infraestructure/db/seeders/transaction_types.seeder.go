package seeders

import (
	"log"

	"github.com/ssssshel/sp-api/src/domain/entities"
	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
)

func TransactionTypes(dbConnection *infra_db.DBConnections) {
	db := dbConnection.DBConn

	tableTypes := []entities.TableType{}

	if err := db.Find(&tableTypes).Error; err != nil {
		log.Fatal("Error finding table types => ", err)
	}

	tableTypeMap := make(map[string]*entities.TableType)
	for i := range tableTypes {
		tableTypeMap[tableTypes[i].Key] = &tableTypes[i]
	}

	transactionTypes := []entities.TransactionType{
		// -------------------------
		// GASTOS (expenses)
		// -------------------------
		{Key: "rent", Description: "Rent / Mortgage", TableType: tableTypeMap["expenses"]},
		{Key: "utilities", Description: "Utilities (water, electricity, internet, etc.)", TableType: tableTypeMap["expenses"]},
		{Key: "transport", Description: "Transport (fuel, public transport, maintenance)", TableType: tableTypeMap["expenses"]},
		{Key: "insurance", Description: "Insurance (health, car, home)", TableType: tableTypeMap["expenses"]},
		{Key: "education", Description: "Education (tuition, courses, books)", TableType: tableTypeMap["expenses"]},
		{Key: "food", Description: "Food (restaurants, delivery)", TableType: tableTypeMap["expenses"]},
		{Key: "groceries", Description: "Groceries (supermarket, market)", TableType: tableTypeMap["expenses"]},
		{Key: "entertainment", Description: "Entertainment (cinema, streaming, hobbies)", TableType: tableTypeMap["expenses"]},
		{Key: "clothing", Description: "Clothing & Accessories", TableType: tableTypeMap["expenses"]},
		{Key: "health", Description: "Health (medicines, medical visits)", TableType: tableTypeMap["expenses"]},
		{Key: "travel", Description: "Travel & Tourism", TableType: tableTypeMap["expenses"]},
		{Key: "technology", Description: "Technology & Gadgets", TableType: tableTypeMap["expenses"]},
		{Key: "gifts", Description: "Gifts & Celebrations", TableType: tableTypeMap["expenses"]},
		{Key: "repairs", Description: "Repairs & Maintenance", TableType: tableTypeMap["expenses"]},
		{Key: "other_expenses", Description: "Other Expenses", TableType: tableTypeMap["expenses"]},

		// -------------------------
		// INGRESOS (invoices)
		// -------------------------
		{Key: "salary", Description: "Salary", TableType: tableTypeMap["invoices"]},
		{Key: "bonus", Description: "Bonuses & Commissions", TableType: tableTypeMap["invoices"]},
		{Key: "freelance", Description: "Freelance / Side Projects", TableType: tableTypeMap["invoices"]},
		{Key: "rent_income", Description: "Rental Income", TableType: tableTypeMap["invoices"]},
		{Key: "royalties", Description: "Royalties (books, music, licenses)", TableType: tableTypeMap["invoices"]},
		{Key: "interests", Description: "Bank Interests", TableType: tableTypeMap["invoices"]},
		{Key: "dividends", Description: "Dividends & Stocks", TableType: tableTypeMap["invoices"]},
		{Key: "business_income", Description: "Business / Entrepreneurship", TableType: tableTypeMap["invoices"]},
		{Key: "crypto_income", Description: "Crypto / Digital Assets", TableType: tableTypeMap["invoices"]},
		{Key: "gifts_income", Description: "Gifts / Donations", TableType: tableTypeMap["invoices"]},
		{Key: "other_invoices", Description: "Other invoices", TableType: tableTypeMap["invoices"]},

		// -------------------------
		// AHORROS (savings)
		// -------------------------
		{Key: "emergency_fund", Description: "Emergency Fund", TableType: tableTypeMap["savings"]},
		{Key: "retirement", Description: "Retirement / Pension Fund", TableType: tableTypeMap["savings"]},
		{Key: "education_savings", Description: "Education Savings", TableType: tableTypeMap["savings"]},
		{Key: "house_savings", Description: "House / Real Estate Savings", TableType: tableTypeMap["savings"]},
		{Key: "vehicle_savings", Description: "Vehicle Savings", TableType: tableTypeMap["savings"]},
		{Key: "travel_savings", Description: "Travel Savings", TableType: tableTypeMap["savings"]},
		{Key: "investment_fund", Description: "Investment Fund (stocks, ETFs, etc.)", TableType: tableTypeMap["savings"]},
		{Key: "business_savings", Description: "Business / Startup Capital", TableType: tableTypeMap["savings"]},
		{Key: "crypto_savings", Description: "Crypto Savings", TableType: tableTypeMap["savings"]},
		{Key: "other_savings", Description: "Other Savings", TableType: tableTypeMap["savings"]},
	}

	SeedData(db, transactionTypes, "key", func(transactionType entities.TransactionType) interface{} {
		return transactionType.Key
	})
}
