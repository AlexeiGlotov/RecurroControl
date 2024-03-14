import React, { useState } from 'react';
import {Link, useNavigate} from 'react-router-dom';
import { axiosInstanceWithoutJWT } from '../api/axios';
import {toast} from "react-toastify";
import '../styles/auth.css'

function RegistrationForm() {
    const [formData, setFormData] = useState({login: '', password: '', repassword: '', access_key: ''});
    const navigate = useNavigate();

    const handleRegistation = async (event) => {
        event.preventDefault();
        if (formData.password !== formData.repassword) {
            toast.error(`password != repassword`);
            return;
        }
        try {
                await axiosInstanceWithoutJWT.post('/auth/sign-up',formData);
                setFormData({ });
                navigate('/login', { replace: true });
                toast.success("Successful registration")

        } catch (error) {
            toast.error(`error: ${error.message}`);
        }
    };

    return (
        <div className="form-container" >
            <form>
                <label>Login:</label>
                <input className="form-input"
                       type="text"
                       name="login"
                       value={formData.login}
                       onChange={(e) => setFormData({...formData, login: e.target.value})}
                />

                <label>Password:</label>
                <input className="form-input"
                       type="password"
                       name="password"
                       value={formData.password}
                       onChange={(e) => setFormData({...formData, password: e.target.value})}
                />

                <label>Re-enter Password:</label>
                <input className="form-input"
                       type="password"
                       name="repassword"
                       value={formData.repassword}
                       onChange={(e) => setFormData({...formData, repassword: e.target.value})}
                />

                <label>Access Key:</label>
                <input className="form-input"
                       type="text"
                       name="access_key"
                       value={formData.access_key}
                       onChange={(e) => setFormData({...formData, access_key: e.target.value})}
                />

                <button className="form-button" onClick={handleRegistation}>Register</button>
                <p>Already have an account? <Link to="/login">Login</Link></p>
            </form>
        </div>
    );
}

export default RegistrationForm;
