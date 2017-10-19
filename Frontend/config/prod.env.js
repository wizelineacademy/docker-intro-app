
const address = process.env.API_PORT_8000_TCP_ADDR || "localhost";
const port = process.env.API_PORT_8000_TCP_PORT || "8000";

module.exports = {
  NODE_ENV: '"production"',
  API_URL: `'http://${address}:${port}'`
}
