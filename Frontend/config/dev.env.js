var merge = require('webpack-merge')
var prodEnv = require('./prod.env')
const address = process.env.API_PORT_8000_TCP_ADDR || "localhost";
const port = process.env.API_PORT_8000_TCP_PORT || "8000";

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  API_URL: JSON.stringify(process.env.API_URL) || `'http://${address}:${port}'`
})
