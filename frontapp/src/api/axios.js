// src/api/axios.js
import axios from 'axios';

// Экземпляр Axios для запросов с JWT
const axiosInstanceWithJWT = axios.create({
    baseURL: 'http://go-app:23678',
});

axiosInstanceWithJWT.interceptors.request.use(
    config => {
        const token = localStorage.getItem('token');
        if (token) {
            config.headers['Authorization'] = 'Bearer ' + token;
        }
        return config;
    },
    error => {
        return Promise.reject(error);
    }
);

// Экземпляр Axios для запросов без JWT
const axiosInstanceWithoutJWT = axios.create({
    baseURL: 'http://go-app:23678',
    // Здесь нет необходимости добавлять JWT в заголовки
});

export { axiosInstanceWithJWT, axiosInstanceWithoutJWT };

/*
import { axiosInstanceWithJWT, axiosInstanceWithoutJWT } from './api/axios';

// Для запроса с JWT
axiosInstanceWithJWT.get('/your-endpoint-with-jwt').then(/!* ... *!/);

// Для запроса без JWT
axiosInstanceWithoutJWT.get('/your-endpoint-without-jwt').then(/!* ... *!/);
*/
