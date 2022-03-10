regex = "ab{2,4}"

input = "abb"
match = input.match(/#{regex}/)
puts match

input = "abbb"
match = input.match(/#{regex}/)
puts match

input = "abbbb"
match = input.match(/#{regex}/)
puts match