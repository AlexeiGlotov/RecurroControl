import './App.css';

import {HashRouter, NavLink, Route, Routes} from "react-router-dom";
import Dashboard from "./components/dashboard";
import Home from "./components/home";
import Contact from "./components/contact";
import About from "./components/about";
import Login from "./components/login";
import { AuthProvider } from './components/AuthContext';
import React, { useContext } from 'react';
import { AuthContext } from './components/AuthContext';
import Navbar from "./components/navbar";

function App() {

    return (
        <AuthProvider>
        <div>
          <HashRouter>
            <Navbar></Navbar>
          <div className="content">
            <Routes>
              <Route path='/dashboard' element={<Dashboard/>}/>
              <Route path='/' element={<Home/>}/>
              <Route path='/contact' element={<Contact/>}/>
              <Route path='/about' element={<About/>}/>
              <Route path='/login' element={<Login/>}/>
            </Routes>
          </div>
        </HashRouter>
      </div>
     </AuthProvider>
  )
}

export default App;