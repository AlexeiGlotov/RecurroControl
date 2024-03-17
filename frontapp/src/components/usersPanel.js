import React, {useState, useEffect, useContext} from 'react';
import { axiosInstanceWithJWT } from "../api/axios";
import "react-toastify/dist/ReactToastify.css";
import {Card, Container, Table,Button} from "react-bootstrap";
import 'bootstrap-icons/font/bootstrap-icons.css';
import {AuthContext} from "./AuthContext";
import handleError from '../utils/errorHandler';

function ManagePanelUsers() {
    const [users, setUsers] = useState([]);
    const {role } = useContext(AuthContext);
    const [isLoading, setIsLoading] = useState(false);

    useEffect(() => {
        setIsLoading(true)
        const fetchUsers = async () => {
            try {
                const response = await axiosInstanceWithJWT.get('/api/users/getUsers');
                setUsers(response.data.users);
            } catch (error) {
                handleError(error);
            } finally {
                setIsLoading(true)
            }
        };

        fetchUsers();
    }, []);


    const handleBanUnbanUser = async (userId, isBanned) => {
        if (isBanned === 0) {
            try {
                await axiosInstanceWithJWT.post('/api/users/unban', {"id": userId});
                setUsers(users.map(user =>
                    user.id === userId ? { ...user, banned: 0 } : user));
            }
            catch(error) {
                handleError(error);
            }
            finally {

            }
        } else {
            try {
                await axiosInstanceWithJWT.post('/api/users/ban', {"id": userId});
                setUsers(users.map(user =>
                    user.id === userId ? { ...user, banned: 1 } : user));
            }
            catch(error) {
                handleError(error);
            }
            finally {
            }
        }
    };

    const handleDeleteUser = async (userId) => {
        try {
            await axiosInstanceWithJWT.post('/api/users/delete', {"id": userId});
            setUsers(users.map(user =>
                user.id === userId ? { ...user, is_deleted: 1 } : user));
        }
        catch(error) {
            handleError(error);
        }
        finally {

        }
    };

    return (
        <div>
            <Container className="my-4">
                <Card className="p-3">
                    <Table striped bordered hover>
                        <thead>
                        <tr>
                            {role === 'admin' && (
                                <th>ID</th>
                            )}
                            <th>Login</th>
                            <th>Role</th>
                            <th>Owner</th>
                            <th>Key-A</th>
                            <th>Key-G</th>
                            <th>Action</th>
                        </tr>
                        </thead>
                        <tbody>
                        {users && users.length > 0 ? (
                            users.filter(user => user.is_deleted !== 1).map((user) => (
                                <tr key={user.id}>
                                    {role === 'admin' && (
                                        <td>{user.id}</td>
                                    )}
                                    <td>{user.login}</td>
                                    <td>{user.role}</td>
                                    <td>{user.owner}</td>
                                    <td>{user.keys_activated}</td>
                                    <td>{user.keys_generated}</td>
                                    <td>
                                        {user.banned === 1 ? (
                                            <Button
                                                onClick={() => handleBanUnbanUser(user.id, 0)}
                                                variant="outline-danger"
                                                className="me-2"
                                            >
                                                <i className="bi bi-lock-fill"></i>
                                            </Button>
                                        ) : (
                                            <Button
                                                onClick={() => handleBanUnbanUser(user.id, 1)}
                                                variant="outline-success"
                                                className="me-2"
                                            >
                                                <i className="bi bi-unlock-fill"></i>
                                            </Button>
                                        )}

                                        {role === 'admin' && (
                                            <>
                                                <Button onClick={() => handleDeleteUser(user.id)} variant="danger"><i
                                                    className="bi bi-trash"></i></Button>
                                            </>
                                        )}
                                    </td>


                                </tr>
                            ))
                        ) : (
                            <>

                                {
                                    isLoading ? (
                                        <tr>
                                            <td colSpan="7">Loading.....</td>
                                        </tr>
                                    ) : (
                                        <tr>
                                            <td colSpan="7">No keys available</td>
                                        </tr>
                                    )
                                }

                            </>
                        )}
                        </tbody>
                    </Table>
                </Card>
            </Container>
        </div>
    );
}

export default ManagePanelUsers;