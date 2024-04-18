const cheerio = require('cheerio');
const fs = require('fs');

const config = JSON.parse(fs.readFileSync('config.json'));

function fetchVouchesCount() {
  const myvouch = config.Vouch['MyVouch_URL']; 
  return fetch(myvouch)
    .then(response => response.text())
    .then(data => {
      const $ = cheerio.load(data);
      const vouchesElement = $('p.social span:last-child');
      const vouchesText = vouchesElement.text().trim();
      const vouchesCount = parseInt(vouchesText.match(/\d+/)[0], 10);
      return vouchesCount;
    })
    .catch(error => {
      console.error('Failed to fetch the vouch count:', error);
      return -1; 
    });
}

function printVouchesCount() {
  fetchVouchesCount().then(count => {
    console.log('Vouch count:', count);
  });
}

const requestDelay = config.Vouch.Request_Delay * 1000;
printVouchesCount(); 
setInterval(printVouchesCount, requestDelay); 
