require 'dotenv'

Dotenv.load("local.env")
puts ENV["STACK"]
puts ENV["DATABASE"]
