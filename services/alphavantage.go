package services

const (
	apiKey                 = "BJCFZLMCBG5UJXZD"
	alphaVantageURL        = "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=1min&apikey=%s&datatype=csv"
	alphaVantageFXURL      = "https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=%s&to_currency=%s&apikey=%s"
	alphaVantageFXDailyURL = "https://www.alphavantage.co/query?function=FX_DAILY&from_symbol=%s&to_symbol=%s&apikey=%s"
)
