require 'uri'


uri = "https://test:abcd123@techbyexample.com:8000/tutorials/intro?type=advance&compact=false#history"
pasrse_uri = URI(uri)

puts(pasrse_uri.scheme)  
puts(pasrse_uri.userinfo)  
puts(pasrse_uri.host)
puts(pasrse_uri.port)
puts(pasrse_uri.path)
puts(pasrse_uri.query)
puts(pasrse_uri.fragment)
puts(pasrse_uri.to_s)
