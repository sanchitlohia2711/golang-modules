require 'uri'


input = "The webiste is https://techbyexample.com:8000/tutorials/intro amd mail to mailto:contactus@techbyexample.com"
urls = URI.extract(input, "https")

puts(urls)