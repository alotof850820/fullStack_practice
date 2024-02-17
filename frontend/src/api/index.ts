import axios from 'axios'
import qs from 'qs'

axios.interceptors.request.use(
  (config) => {
    config.baseURL = 'http://127.0.0.1:8081/'
    config.headers['Content-Type'] = 'application/json'
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

export const apiPostlogin = (data: any) => {
  const queryString = qs.stringify(data)
  return axios.post(`user/findUserByNameAndPassword?${queryString}`)
}

export const apiPostRegister = (data: any) => {
  const queryString = qs.stringify(data)
  return axios.get(`user/createUser?${queryString}`)
}
