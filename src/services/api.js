import axios from 'axios'

export default {
  getDefaultOptions (url) {
    return {
      json: true,
      withCredeintials: false,
      headers: {
      },
      url: process.env.API_URL + url
    }
  },

  get (url) {
    const options = this.getDefaultOptions(url)
    options['method'] = 'GET'
    return axios(options)
  },

  delete (url) {
    const options = this.getDefaultOptions(url)
    options['method'] = 'DELETE'
    return axios(options)
  },

  put (url, data) {
    const options = this.getDefaultOptions(url)
    options['method'] = 'PUT'
    options['data'] = data
    return axios(options)
  },

  post (url, data) {
    const options = this.getDefaultOptions(url)
    options['method'] = 'POST'
    options['data'] = data
    return axios(options)
  }
}
