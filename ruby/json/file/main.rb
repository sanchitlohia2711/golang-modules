require 'json'
file = File.read('temp.json')
parsed = JSON.parse(file)

puts parsed
puts parsed["a"]
puts parsed["b"]