class String
    def bold
        "\e[1m#{self}\e[22m" 
     end
end

puts "\n"
puts "Bold String".bold
puts "\n"