import {NavLink} from "react-router-dom";
import React,{useContext} from "react";
import { AuthContext } from './AuthContext';

function Navbar() {
    const { isAuthenticated } = useContext(AuthContext);
    return (
        <div className="navbar">
            <ul>
                <li><NavLink to="/">Home</NavLink></li>
                <li><NavLink to="/dashboard">Dashboard</NavLink></li>
                <li><NavLink to="/contact">Contact</NavLink></li>
                <li><NavLink to="/about">About</NavLink></li>
                <div>
                    {isAuthenticated ? <li>true</li> :
                        <li>false</li>} 
                </div>
            </ul>
        </div>
    )
}

export default Navbar;