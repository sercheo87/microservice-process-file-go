package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// Define mock data constants
	regions := []string{"North America", "Europe", "Asia", "South America", "Australia"}
	countries := []string{"USA", "Canada", "Mexico", "Germany", "France", "UK", "China", "Japan", "India", "Brazil", "Argentina", "Australia"}
	itemTypes := []string{"Clothing", "Electronics", "Food", "Furniture", "Toys", "Sports"}
	salesChannels := []string{"Online", "Offline"}
	orderPriorities := []string{"High", "Medium", "Low"}

	// Create the CSV file
	file, err := os.Create("mock-data/sales.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create the CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the CSV header
	header := []string{"Region", "Country", "Item Type", "Sales Channel", "Order Priority", "Order Date", "Order ID", "Ship Date", "Units Sold", "Unit Price", "Unit Cost", "Total Revenue", "Total Cost", "Total Profit"}
	if err := writer.Write(header); err != nil {
		panic(err)
	}

	// Generate mock data
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 1000; i++ {
		region := regions[rand.Intn(len(regions))]
		country := countries[rand.Intn(len(countries))]
		itemType := itemTypes[rand.Intn(len(itemTypes))]
		salesChannel := salesChannels[rand.Intn(len(salesChannels))]
		orderPriority := orderPriorities[rand.Intn(len(orderPriorities))]
		orderDate := time.Now().AddDate(0, 0, -rand.Intn(365)).Format("2006-01-02")
		orderID := fmt.Sprintf("%d", i)
		shipDate := time.Now().AddDate(0, 0, rand.Intn(7)+1).Format("2006-01-02")
		unitsSold := strconv.Itoa(rand.Intn(1000) + 1)
		unitPrice := strconv.FormatFloat(float64(rand.Intn(100)+1)/10.0, 'f', 2, 64)
		unitCost := strconv.FormatFloat(float64(rand.Intn(50)+1)/10.0, 'f', 2, 64)
		totalRevenue := strconv.FormatFloat(float64(rand.Intn(1000)+1)/10.0, 'f', 2, 64)
		totalCost := strconv.FormatFloat(float64(rand.Intn(500)+1)/10.0, 'f', 2, 64)
		totalProfit := strconv.FormatFloat(float64(rand.Intn(500)+1)/10.0, 'f', 2, 64)

		// Write the CSV row
		row := []string{region, country, itemType, salesChannel, orderPriority, orderDate, orderID, shipDate, unitsSold, unitPrice, unitCost, totalRevenue, totalCost, totalProfit}
		if err := writer.Write(row); err != nil {
			panic(err)
		}
	}
}
