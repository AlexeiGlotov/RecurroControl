import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { axiosInstanceWithoutJWT } from '../api/axios';
function RegistrationForm() {
    const [formData, setFormData] = useState({
        login: '',
        password: '',
        repassword: '',
        access_key: ''
    });

    const navigate = useNavigate();

    const handleChange = (e) => {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        if (formData.password !== formData.repassword) {
            alert('Passwords do not match');
            return;
        }

        try {
            const response = await axiosInstanceWithoutJWT.post('/auth/sign-up',formData);
            //const response = await axios.post('http://localhost:8080/auth/sign-in', { username, password });
           // const { id } = response.data;
            navigate('/login');
        } catch (error) {
            console.error('registation failed', error);
        }

    };

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <label>Login:</label>
                <input
                    type="text"
                    name="login"
                    value={formData.login}
                    onChange={handleChange}
                />
            </div>

            <div>
                <label>Password:</label>
                <input
                    type="password"
                    name="password"
                    value={formData.password}
                    onChange={handleChange}
                />
            </div>

            <div>
                <label>Re-enter Password:</label>
                <input
                    type="password"
                    name="repassword"
                    value={formData.repassword}
                    onChange={handleChange}
                />
            </div>

            <div>
                <label>Access Key:</label>
                <input
                    type="text"
                    name="access_key"
                    value={formData.access_key}
                    onChange={handleChange}
                />
            </div>

            <button type="submit">Register</button>
        </form>
    );
}

export default RegistrationForm;
