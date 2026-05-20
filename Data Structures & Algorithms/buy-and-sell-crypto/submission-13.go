
func maxProfit(prices []int) int {
    minPrice := prices[0]
    left := 0

    for _,price := range prices {

        if price < minPrice { 
            minPrice = price
        }

        if price - minPrice > left {
            left = price - minPrice
        }
    }

    return left
}
