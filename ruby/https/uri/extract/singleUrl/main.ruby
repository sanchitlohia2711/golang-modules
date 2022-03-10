require 'uri'


input = "The webiste is https://techbyexample.com:8000/tutorials/intro"
urls = URI.extract(input)

puts(urls)