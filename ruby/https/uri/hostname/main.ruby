require 'uri'


uri = "https://test:abcd123@techbyexample.com:8000/tutorials/intro?type=advance&compact=false#history"
pasrse_uri = URI(uri)

hostname = pasrse_uri.scheme + "://" + pasrse_uri.hostname + ":" + pasrse_uri.port.to_s

puts hostname
