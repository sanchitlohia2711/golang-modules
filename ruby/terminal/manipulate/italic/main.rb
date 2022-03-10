class String
    def italic
        "\e[3m#{self}\e[23m"
     end
end

puts "\n"
puts "Italic String".italic
puts "\n"