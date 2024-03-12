import {NavLink} from "react-router-dom";
import React,{useContext} from "react";
import { AuthContext } from './AuthContext';

function Navbar() {
    const { isAuthenticated } = useContext(AuthContext);
    const { logout } = useContext(AuthContext);
    return (
        <div className="navbar">
            <ul>
                <li><NavLink to="/">Dashboard</NavLink></li>
                <li><NavLink to="/license-keys">License Keys</NavLink></li>
                <li><NavLink to="/access-keys">Access Keys</NavLink></li>
                <li><NavLink to="/cheats">Cheats</NavLink></li>
                <li>
                    {isAuthenticated ? <li>
                            <button onClick={logout}>Logout</button>
                        </li> :
                        <li>false</li>}
                </li>
            </ul>
        </div>
    )
}

export default Navbar;