require 'uri'


uri = "https://test:abcd123@techbyexample.com:8000/tutorials/intro?type=advance&compact=false#history"
pasrse_uri = URI(uri)

puts(pasrse_uri.query)
