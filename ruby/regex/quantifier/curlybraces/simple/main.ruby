regex = "b{2}"

input = "bb"
match = input.match(/#{regex}/)
puts match

input = "bbb"
match = input.match(/#{regex}/)
puts match