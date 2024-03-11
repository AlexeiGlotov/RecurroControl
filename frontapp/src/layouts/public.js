import React from 'react';
import { HashRouter, Routes, Route, Navigate } from 'react-router-dom';
import Login from '../components/login';
import Contact from '../components/contact';

function PublicRoutes() {
    return (
        <HashRouter>
            <Routes>
                <Route path='/login' element={<Login />} />
                <Route path='/contact' element={<Contact />} />
                <Route path="*" element={<Navigate to="/login" replace />} />
            </Routes>
        </HashRouter>
    );
}

export default PublicRoutes;