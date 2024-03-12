import React from 'react';
import { HashRouter, Routes, Route, Navigate } from 'react-router-dom';
import Login from '../components/login';
import RegistrationForm from "../components/registration";

function PublicRoutes() {
    return (
        <HashRouter>
            <Routes>
                <Route path='/login' element={<Login />} />
                <Route path='/registration' element={<RegistrationForm />} />
                <Route path="*" element={<Navigate to="/login" replace />} />
            </Routes>
        </HashRouter>
    );
}

export default PublicRoutes;