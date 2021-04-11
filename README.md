# Cutt.ly URL Shortener

Intuitive GO wrapper for [cuttly] (https://cutt.ly/)

# Pre-requisite
Create or Login to your account at https://cutt.ly/login

On the left panel `Your Profile`, you can see a list of menus. Click on `Edit Account` then to your right you will see an `API Key`.

# Installation 
```
go get https://github.com/ernestngugi/go-cuttly

```

# Shorten URL
```
credentials := &Client{
		Key: "",
		BaseURL: "",
	}
  
 shortURL, err := credentials.Shorten("https://www.avery.long/url","sample")
 ```
 
 # Get Statistics
 ```
 credentials := &Client{
		Key: "",
		BaseURL: "",
	}
  
  stats,err := credentials.GetStats("https://cutt.ly/qwerty")
  ```
