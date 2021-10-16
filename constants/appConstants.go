package constants

const (
	LATITUDE                       = "latitude"
	LONGITUDE                      = "longitude"
	REVERSE_GEOCODING_BASE_URL     = "https://us1.locationiq.com/v1/reverse.php?key="
	FETCH_CURRENT_CASES_URL        = "https://api.covid19api.com/live/country/india/status/active?"
	FETCH_CURRENT_CASES_BY_COUNTRY = "https://api.covid19api.com/total/country/india?"
	DATABASE_NAME                  = "covid_cases"
	COUNTRY_COLLECTION             = "country_data"
	STATE_COLLECTION               = "state_data"
	REDIS_ADDRESS                  = "redis-18220.c264.ap-south-1-1.ec2.cloud.redislabs.com:18220"
	CACHE_TTL                      = 1800
	STATE_DATA_CACHE_KEY           = "state:data:"
	COUNTRY_DATA_CACHE_KEY         = "country:data:"
)
