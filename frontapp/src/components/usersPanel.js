import React, {useState, useEffect, useContext} from 'react';
import { axiosInstanceWithJWT } from "../api/axios";
import { toast } from 'react-toastify';
import "react-toastify/dist/ReactToastify.css";
import {Card, Container, Table,Button} from "react-bootstrap";
import 'bootstrap-icons/font/bootstrap-icons.css';
import {AuthContext} from "./AuthContext";

function ManagePanelUsers() {
    const [users, setUsers] = useState([]);
    const {id,role } = useContext(AuthContext);

    useEffect(() => {
        const fetchUsers = async () => {
            try {
                const response = await axiosInstanceWithJWT.get('/api/users/getUsers');
                setUsers(response.data.users);
            } catch (error) {
                toast.error(`error: ${error.message}`);
            } finally {
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
                toast.error(`error: ${error.message}`);
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
                toast.error(`error: ${error.message}`);
            }
            finally {
            }
        }
        // Отправка запроса на сервер для ban/unban
    };

    const handleDeleteUser = async (userId) => {
        try {
            await axiosInstanceWithJWT.post('/api/users/delete', {"id": userId});
            setUsers(users.map(user =>
                user.id === userId ? { ...user, is_deleted: 1 } : user));
        }
        catch(error) {
            toast.error(`error: ${error.message}`);
        }
        finally {

        }
    };

    return (
        <div>
            <Container className="my-4">
                <Card className="p-3">
                    <Table striped bordered hover> {/* Измененный стиль таблицы */}
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
                                                variant="primary"
                                                className="me-2"
                                                disabled={user.id === id} // Отключаем кнопку, если user.id равен id
                                            >
                                                <i className="bi bi-unlock-fill"></i>
                                            </Button>
                                        ) : (
                                            <Button
                                                onClick={() => handleBanUnbanUser(user.id, 1)}
                                                variant="primary"
                                                className="me-2"
                                                disabled={user.id === id} // Отключаем кнопку, если user.id равен id
                                            >
                                                <i className="bi bi-lock-fill"></i>
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
                            <tr>
                                <td colSpan="7">No keys available</td>
                            </tr>
                        )}
                        </tbody>
                    </Table>
                </Card>
            </Container>





        </div>

    );
}

export default ManagePanelUsers;