import json
import requests
from bs4 import BeautifulSoup
import re
import time

with open('config.json') as f:
    config = json.load(f)

def fetchVouchesCount():
    try:
        response = requests.get(config['Vouch']['MyVouch_URL'])
        soup = BeautifulSoup(response.text, 'html.parser')
        vouches_element = soup.find('p', class_='social').find_all('span')[-1]
        vouches_text = vouches_element.get_text().strip()
        vouches_count = int(re.search(r'\d+', vouches_text).group())
        return vouches_count
    except Exception as e:
        print('Failed to fetch the vouch count:', e)
        return -1

def printVouchesCount():
    count = fetchVouchesCount()
    print('Vouch count:', count)

printVouchesCount()
request_delay = config['Vouch']['Request_Delay']
time.sleep(request_delay)
while True:
    printVouchesCount()
    time.sleep(request_delay)
