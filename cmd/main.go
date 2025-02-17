// product: Asus DUAL EVO OC GeForce RTX 4070 12 GB Video Card, price: $818.16 CAD (make sure canadian)
// product 2: AMD Ryzen 7 7700X 4.5GHz 8-Core Processor, price: $435.50 CAD
// product 3: Corsair iCUE H150i ELITE CAPELLIX 65.57 CFM Liquid CPU Cooler, price: $309.02 CAD
package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
)

func sendEmail(productName, productPrice string) {
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")
	to := os.Getenv("EMAIL_TO")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := fmt.Sprintf("Subject: Price Alert\n\nProduct: %s, Price: %s", productName, productPrice)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Printf("Error sending email: %s", err.Error())
	} else {
		log.Printf("Email sent successfully")
	}
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Instantiate default collector
	c := colly.NewCollector()
	fmt.Printf("Outside OnHTML @@@@@@@@@@\n")
	// On every div element with a class that includes "main-wrapper" call callback
	c.OnHTML("div[class*='main-wrapper']", func(e *colly.HTMLElement) {
		fmt.Printf("Inside OnHTML @@@@@@@@@@\n")
		// Extract product name
		productName := e.ChildText("h1.pageTitle")
		if productName == "Asus DUAL EVO OC GeForce RTX 4070 12 GB Video Card" {
			productName = "GPU found: Asus DUAL EVO OC GeForce RTX 4070 12 GB Video Card"
		} else if productName == "AMD Ryzen 7 7700X 4.5GHz 8-Core Processor" {
			productName = "CPU found: AMD Ryzen 7 7700X 4.5GHz 8-Core Processor"
		} else if productName == "Corsair iCUE H150i ELITE CAPELLIX 65.57 CFM Liquid CPU Cooler" {
			productName = "Liquid Cooler found: Corsair iCUE H150i ELITE CAPELLIX 65.57 CFM Liquid CPU Cooler"
		}
		// Extract product price from td element with class td__finalPrice
		productPrice := e.ChildText("tr td.td__finalPrice")

		// Print product name and price
		fmt.Printf("Product: %s, Price: %s\n", productName, productPrice)

		// Check if the price is lower than $820.00
		priceStr := strings.Replace(productPrice, "$", "", -1)
		priceStr = strings.Replace(priceStr, ",", "", -1)
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Printf("Error parsing price: %s", err.Error())
			return
		}

		if price < 820.00 {
			sendEmail(productName, productPrice)
		}
	})

	// List of URLs to scrape
	urls := []string{
		"https://ca.pcpartpicker.com/product/4mbRsY/asus-dual-evo-oc-geforce-rtx-4070-12-gb-video-card-dual-rtx4070-o12g-evo-white",
		"https://example.com/another-product-page",
		"https://example.com/yet-another-product-page",
	}

	// Visit each URL
	for _, url := range urls {
		fmt.Printf("Visiting URL: %s\n", url)
		c.Visit(url)
		fmt.Print("Visited URL\n")
	}
}
