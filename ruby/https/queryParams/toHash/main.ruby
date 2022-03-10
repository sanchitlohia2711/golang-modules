input_sring = 'a=b&x=y'

query_params_hash = {}
input_string_split = input_sring.split('&')
input_string_split.each do |q|
    q_split = q.split('=')
    query_params_hash[q_split[0]] = q_split[1]
end 

puts query_params_hash
