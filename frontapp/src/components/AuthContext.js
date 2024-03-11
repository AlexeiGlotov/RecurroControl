// AuthContext.js
import React, { createContext, useState, useEffect } from 'react';

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [authState, setAuthState] = useState({
        token: localStorage.getItem('token'),
        isAuthenticated: false,
    });

    useEffect(() => {
        if (authState.token) {
            setAuthState({ ...authState, isAuthenticated: true });
        }
    }, [authState.token]);

    const login = (token) => {
        localStorage.setItem('token', token);
        setAuthState({ ...authState, token, isAuthenticated: true });
    };

    const logout = () => {
        localStorage.removeItem('token');
        setAuthState({ ...authState, token: null, isAuthenticated: false });
    };

    return (
        <AuthContext.Provider value={{ ...authState, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};