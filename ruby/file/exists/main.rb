File.open("temp.txt", "w") do |f|     
    f.write("Test")   
  end
  
puts File.exist?('temp.txt')