import './App.css';
import Main from './layouts/main';
import {AuthProvider} from "./components/AuthContext";
import React from "react";

function App() {
    return (
        <AuthProvider>
            <Main></Main>
        </AuthProvider>
    );
}

export default App;