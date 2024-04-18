require 'nokogiri'
require 'open-uri'
require 'json'

config = JSON.parse(File.read('config.json'))

def fetch_vouches_count(config)
  begin
    doc = Nokogiri::HTML(open(config['Vouch']['MyVouch_URL']))
    vouches_element = doc.at('p.social span:last-child')
    vouches_text = vouches_element.text.strip
    vouches_count = vouches_text.match(/\d+/)[0].to_i
    return vouches_count
  rescue StandardError => e
    puts "Failed to fetch the vouch count: #{e}"
    return -1
  end
end

def print_vouches_count(config)
  count = fetch_vouches_count(config)
  puts "Vouch count: #{count}"
end

print_vouches_count(config)
request_delay = config['Vouch']['Request_Delay']
sleep(request_delay)
loop do
  print_vouches_count(config)
  sleep(request_delay)
end
