import axios from 'axios'

const http = axios.create({
  baseURL: 'http://localhost:3000/api'
})

http.interceptors.request.use(config => {
  if (localStorage.token) {
    config.headers.Authorization = localStorage.token
  }
  return config;
}, (error) => {
  // Do something with request error
  return Promise.reject(error);
});

export default http
