import './App.css';
import Main from './layouts/main';
import {AuthProvider} from "./components/AuthContext";
import React from "react";
import {ToastContainer} from "react-toastify";
import 'bootstrap/dist/css/bootstrap.min.css';


function App() {
    return (
        <AuthProvider>
            <ToastContainer/>
            <Main></Main>
        </AuthProvider>
    );
}

export default App;