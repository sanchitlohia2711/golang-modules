regex = "b+"

input = "bb"
match = input.match(/#{regex}/)
puts match

input = "bbb"
match = input.match(/#{regex}/)
puts match

input = "a"
match = input.match(/#{regex}/)
puts match