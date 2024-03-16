// Login.js
import React, { useState,useContext } from 'react';
import { axiosInstanceWithoutJWT } from '../api/axios';
import { AuthContext } from './AuthContext';
import { useNavigate,Link } from 'react-router-dom';
import {toast} from "react-toastify";

import { Container, Form, Card, Button } from 'react-bootstrap';

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
            switch (error.response.data.message) {
                case "not all fields are filled in":
                    toast.error("not all fields are filled in");
                    break
                case "bad len login":
                    toast.error("bad len login 4-20");
                    break
                case "bad len password":
                    toast.error("bad len password 6-20");
                    break
                default:
                    toast.error("server error");
            }
        }
    };

    return (
        <Container className="d-flex align-items-center justify-content-center" style={{ minHeight: "100vh" }}>
            <Card className="w-100" style={{ maxWidth: "400px" }}>
                <Card.Body>
                    <h2 className="text-center mb-4">Login</h2>
                    <Form>
                        <Form.Group controlId="formUsername">
                            <Form.Control
                                type="text"
                                value={username}
                                onChange={(e) => setUsername(e.target.value)}
                                placeholder="login"
                            />
                        </Form.Group>

                        <Form.Group controlId="formPassword" className="mb-4 mt-3">
                            <Form.Control
                                type="password"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                                placeholder="password"
                            />
                        </Form.Group>

                        <Button variant="primary" onClick={handleLogin} className="w-100">
                            Login
                        </Button>
                    </Form>
                    <div className="text-center mt-3">
                        Not registered yet? <Link to="/registration">Register</Link>
                    </div>
                </Card.Body>
            </Card>
        </Container>
    );
};

export default Login;