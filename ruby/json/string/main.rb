require 'json'
parsed = JSON.parse('{"a":"x","b":"y"}')

puts parsed
puts parsed["a"]
puts parsed["b"]