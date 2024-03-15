import React,{useContext} from "react";
import {HashRouter, Route, Routes,Navigate,useLocation} from "react-router-dom";
import AccessKeys from '../components/accessKeys';
import Navbar from "../components/navbar";
import Dashboard from "../components/dashboard";
import Cheats from "../components/cheats";
import { AuthContext } from '../components/AuthContext';
import ManagePanelUsers from "../components/usersPanel";
import GenerateLicenseKeys from "../components/generateLicenseKeys";
import ListLicenseKeys from "../components/listLicenseKeys";

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
                        <Route path='/' element={<ProtectedRoute><Dashboard /></ProtectedRoute>} />
                        <Route path='/list-license-keys' element={<ProtectedRoute><ListLicenseKeys /></ProtectedRoute>} />
                        <Route path='/generate-license-keys' element={<ProtectedRoute><GenerateLicenseKeys /></ProtectedRoute>} />
                        <Route path='/access-keys' element={<ProtectedRoute><AccessKeys /></ProtectedRoute>} />
                        <Route path="/cheats"  element={<ProtectedRoute><Cheats /></ProtectedRoute>} />
                        <Route path="/managepanelusers"  element={<ProtectedRoute><ManagePanelUsers /></ProtectedRoute>} />
                    </Routes>
                </div>
            </HashRouter>
        </div>
    )
}

export default PrivateRoutes;