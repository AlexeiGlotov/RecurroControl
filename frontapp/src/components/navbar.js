import React, { useContext } from "react";
import { AuthContext } from './AuthContext';
import '../styles/navbar.css';
import { Navbar, Nav, Button } from 'react-bootstrap';
import { NavLink } from 'react-router-dom';

function NavigationBar() {
    const { isAuthenticated, logout,role } = useContext(AuthContext);

    return (
        <Navbar collapseOnSelect expand="lg" bg="light" variant="light">
            <Navbar.Brand as={NavLink} to="/">BrandName</Navbar.Brand>
            <Navbar.Toggle aria-controls="responsive-navbar-nav" />
            <Navbar.Collapse id="responsive-navbar-nav">
                <Nav className="mr-auto">
                    <Nav.Link as={NavLink} to="/" exact>Dashboard</Nav.Link>
                    <Nav.Link as={NavLink} to="/list-license-keys">List keys</Nav.Link>

                    {(role === 'admin' || role === 'distributors' || role === 'reseller') && (
                        <>
                        <Nav.Link as={NavLink} to="/generate-license-keys">Generate key</Nav.Link>
                        <Nav.Link as={NavLink} to="/managepanelusers">ManageUsers</Nav.Link>
                            <Nav.Link as={NavLink} to="/access-keys">Access Keys</Nav.Link>
                        </>
                    )}

                    {role === 'admin' && (
                        <Nav.Link as={NavLink} to="/cheats">Cheats</Nav.Link>
                    )}
                </Nav>

                <Nav className="ms-auto">
                    <Navbar.Text className="me-3">{role}</Navbar.Text>

                    {isAuthenticated && (
                        <Button onClick={logout} variant="outline-danger" className="me-3">
                            Logout
                        </Button>
                    )}
                </Nav>
            </Navbar.Collapse>
        </Navbar>
    );
}

export default NavigationBar;