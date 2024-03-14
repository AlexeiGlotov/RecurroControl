// AuthContext.js
import React, { createContext, useState, useEffect } from 'react';
import {jwtDecode} from "jwt-decode";

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [authState, setAuthState] = useState({
        token: localStorage.getItem('token'),
        isAuthenticated: false,
        role: null
    });

    useEffect(() => {
        if (authState.token) {
            const decodedToken = jwtDecode(authState.token);
            setAuthState({...authState, isAuthenticated: true, role: decodedToken.Role});
        }
    }, [authState.token]);

        const login = (token) => {
            localStorage.setItem('token', token);
            const decodedToken = jwtDecode(token);
            setAuthState({ ...authState, token, isAuthenticated: true ,role: decodedToken.Role});
        };

        const logout = () => {
            localStorage.removeItem('token');
            setAuthState({ ...authState, token: null, role:null, isAuthenticated: false });
        };

        return (
            <AuthContext.Provider value={{ ...authState, login, logout }}>
                {children}
            </AuthContext.Provider>
        );
    };