// src/api/axios.js
import axios from 'axios';
import authEvent from '../components/authEvent';

const axiosInstanceWithJWT = axios.create({
    baseURL: process.env.REACT_APP_API_URL
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

axiosInstanceWithJWT.interceptors.response.use(
    response => response,
    error => {
        if (error.response && error.response.status === 401) {
            authEvent.emit('unauthorized');
        }
        return Promise.reject(error);
    }
);

const axiosInstanceWithoutJWT = axios.create({
    baseURL: process.env.REACT_APP_API_URL,
});

export { axiosInstanceWithJWT, axiosInstanceWithoutJWT };

/*
import { axiosInstanceWithJWT, axiosInstanceWithoutJWT } from './api/axios';

// Для запроса с JWT
axiosInstanceWithJWT.get('/your-endpoint-with-jwt').then(/!* ... *!/);

// Для запроса без JWT
axiosInstanceWithoutJWT.get('/your-endpoint-without-jwt').then(/!* ... *!/);
*/
