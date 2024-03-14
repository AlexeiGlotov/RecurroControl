// Login.js
import React, { useState,useContext } from 'react';
import { axiosInstanceWithoutJWT } from '../api/axios';
import { AuthContext } from './AuthContext';
import { useNavigate,Link } from 'react-router-dom';
import {toast} from "react-toastify";
import '../styles/auth.css'

const Login = () => {
    const { login } = useContext(AuthContext);
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    let navigate = useNavigate();

    const handleLogin = async (event) => {
        event.preventDefault();
        try {
            const response = await axiosInstanceWithoutJWT.post('/auth/sign-in',{username,password});
            const { token } = response.data;
            login(token)
            navigate('/');
        } catch (error) {
            toast.error(`error: ${error.message}`);
        }
    };

    return (
        <div className="form-container">
            <form >
                <h2>Login</h2>
                <label> Username:</label>
                <input type="text" className="form-input" value={username}
                       onChange={(e) => setUsername(e.target.value)}/>

                <label>Password:</label>
                <input type="password" className="form-input" value={password}
                       onChange={(e) => setPassword(e.target.value)}/>

                <button className="form-button" onClick={handleLogin}>Login</button>
                <p>Not registered yet? <Link to="/registration">Registration</Link></p>
            </form>
        </div>
);
};

export default Login;