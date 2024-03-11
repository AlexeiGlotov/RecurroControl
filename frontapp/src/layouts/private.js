import React,{useContext} from "react";
import {HashRouter, Route, Routes,Navigate,useLocation} from "react-router-dom";
import Login from '../components/login';
import Contact from '../components/contact';
import Navbar from "../components/navbar";
import Dashboard from "../components/dashboard";
import Home from "../components/home";
import About from "../components/about";
import { AuthContext } from '../components/AuthContext';

// Компонент ProtectedRoute
const ProtectedRoute = ({ children }) => {
    const {isAuthenticated} = useContext(AuthContext);
    const location = useLocation();

    if (!isAuthenticated) {
        return <Navigate to="/login" replace state={{ from: location }} />;
    }

    return children;
};

function PrivateRoutes() {
    return (
        <div>
            <HashRouter>
                <Navbar></Navbar>
                <div className="content">
                    <Routes>
                        <Route path='/dashboard' element={<Dashboard/>}/>
                        <Route path='/' element={<Home/>}/>
                        <Route path='/contact' element={<Contact/>}/>
                        <Route path="/about"  element={<ProtectedRoute><About /></ProtectedRoute>} />
                        <Route path='/login' element={<Login/>}/>
                    </Routes>
                </div>
            </HashRouter>
        </div>
    )
}

export default PrivateRoutes;