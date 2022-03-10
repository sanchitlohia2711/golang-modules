regex = "b+"

input = "bb"
match = input.match(/#{regex}/)
puts match

input = "bbb"
match = input.match(/#{regex}/)
puts match

input = "abb"
match = input.match(/a#{regex}/)
puts match