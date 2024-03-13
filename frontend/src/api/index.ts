import axios from 'axios'
import qs from 'qs'

const jsonHttp = axios.create()
const formHttp = axios.create()

formHttp.interceptors.request.use(
  (config) => {
    // 放ngrok的網址
    config.baseURL = 'http://127.0.0.1:8081/'
    config.headers['Content-Type'] = 'application/x-www-form-urlencoded'
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)
jsonHttp.interceptors.request.use(
  (config) => {
    // 放ngrok的網址
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
  return jsonHttp.post(`user/findUserByNameAndPassword?${queryString}`)
}

export const apiPostRegister = (data: any) => {
  const queryString = qs.stringify(data)
  return jsonHttp.get(`user/createUser?${queryString}`)
}

export const apiPostSearchFriends = (data: any) => {
  return formHttp.post(`searchFriends`, data)
}

export const apiPostUploadImg = (data: any) => {
  return formHttp.post(`attach/upload`, data)
}

export const apiPostaddFriend = (data: any) => {
  return formHttp.post(`contact/addFriend`, data)
}
export const apiPostCreateGroup = (data: any) => {
  return formHttp.post(`contact/createCommunity`, data)
}
export const apiPostLoadGroup = (data: any) => {
  return formHttp.post(`contact/loadCommunity`, data)
}
export const apiPostJoinGroup = (data: any) => {
  return formHttp.post(`contact/joinGroup`, data)
}
