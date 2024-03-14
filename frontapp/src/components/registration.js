import React, { useState } from 'react';
import {Link, useNavigate} from 'react-router-dom';
import { axiosInstanceWithoutJWT } from '../api/axios';
import {toast} from "react-toastify";
import '../styles/auth.css'

import { Container, Form, Card, Button } from 'react-bootstrap';

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
        <Container className="d-flex align-items-center justify-content-center" style={{ minHeight: "100vh" }}>
            <Card className="w-100" style={{ maxWidth: "400px" }}>
                <Card.Body>
                    <h2 className="text-center mb-4">Registration</h2>
                    <Form>
                    <Form.Group controlId="formBasicLogin" className="mb-3">

                        <Form.Control
                            type="text"
                            name="login"
                            value={formData.login}
                            onChange={(e) => setFormData({...formData, login: e.target.value})}
                            placeholder="username"
                        />
                    </Form.Group>

                    <Form.Group controlId="formBasicPassword" className="mb-3">
                        <Form.Control
                            type="text"
                            name="login"
                            value={formData.password}
                           onChange={(e) => setFormData({...formData, password: e.target.value})}
                            placeholder="password"
                        />
                    </Form.Group>

                    <Form.Group controlId="formBasicRePassword" className="mb-3">

                        <Form.Control
                            type="text"
                            name="login"
                            value={formData.repassword}
                           onChange={(e) => setFormData({...formData, repassword: e.target.value})}
                            placeholder="repeat password"
                        />
                    </Form.Group>


                    <Form.Group controlId="formBasicAccessKey" className="mb-3">

                        <Form.Control
                            type="text"
                            name="login"
                            value={formData.access_key}
                            onChange={(e) => setFormData({...formData, access_key: e.target.value})}
                            placeholder="access key"
                        />
                    </Form.Group>

                        <Button variant="primary" onClick={handleRegistation}  className="w-100">
                            Register
                        </Button>
                    </Form>
                    <div className="text-center mt-3">
                        Already have an account? <Link to="/login">Login</Link>
                    </div>
                    </Card.Body>
            </Card>
        </Container>
    );
}

export default RegistrationForm;
