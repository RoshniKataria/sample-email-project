func isMaintenanceMode() bool {
    if _, err := os.Stat(maintenanceFlagPath); err == nil {
        return true
    }
    return false
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    if isMaintenanceMode() {
        tmpl.ExecuteTemplate(w, "maintenance.html", nil)
        return
    }

    log.Println("HomeHandler: Fetching product details")
    resp, err := http.Get("http://productlist:8081/products") // Replace with actual API URL
    if err != nil {
        log.Printf("ERROR: HomeHandler: Error fetching products - %v\n", err)
        http.Error(w, "Unable to fetch products", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        log.Printf("ERROR: HomeHandler: Unexpected status code %d from products API\n", resp.StatusCode)
        http.Error(w, "Error fetching products", http.StatusInternalServerError)
        return
    }

    var products []Product
    if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
        log.Printf("ERROR: HomeHandler: Error parsing product data - %v\n", err)
        http.Error(w, "Error parsing product data", http.StatusInternalServerError)
        return
    }

    log.Printf("HomeHandler: Successfully fetched %d products\n", len(products))
    tmpl.ExecuteTemplate(w, "home.html", products)
}

