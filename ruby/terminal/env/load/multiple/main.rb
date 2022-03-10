require 'dotenv'

Dotenv.load("local.env", "test.env")
puts ENV["STACK"]
puts ENV["DATABASE"]
puts ENV["TEST"]
