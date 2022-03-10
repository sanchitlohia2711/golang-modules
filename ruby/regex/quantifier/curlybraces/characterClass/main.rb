regex = "[ab]{4}"

input = "ababb"
match = input.match(/#{regex}/)
puts match

input = "ababbc"
match = input.match(/#{regex}/)
puts match
