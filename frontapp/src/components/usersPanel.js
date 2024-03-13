import React, { useState, useEffect } from 'react';
import { axiosInstanceWithJWT } from "../api/axios";
import { toast } from 'react-toastify';
import "react-toastify/dist/ReactToastify.css";


function ManagePanelUsers() {
    const [users, setUsers] = useState([]);
    const [user, setUser] = useState([]);

    useEffect(() => {
        const fetchUsers = async () => {
            try {
                const response = await axiosInstanceWithJWT.get('/api/users/getUsers');
                setUsers(response.data.users);
                toast.success("loading page success")
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

    const handleInfoUser = async (userId) => {

        try {
            const response = await axiosInstanceWithJWT.post('/api/users/getUser', {"id": userId});
            setUser(response.data.user)
        }
        catch(error) {
            toast.error(`error: ${error.message}`);
        }
        finally {

        }

    };


    function UserInfoDisplay({ userInfo }) {
        return (
            <div>
            {Object.keys(user).length > 0 && (
                    // Вывод информации о пользователе, если объект user не пуст
                    <div>
                        <p>ID: {userInfo.id}</p>
                        <p>Login: {userInfo.login}</p>
                        <p>Role: {userInfo.role}</p>
                        <p>Keys Generated: {userInfo.keys_generated}</p>
                        <p>Keys Activated: {userInfo.keys_activated}</p>
                        <p>Banned: {userInfo.banned === 1 ? 'Yes' : 'No'}</p>
                        <p>Owner: {userInfo.owner}</p>
                        <p>Deleted: {userInfo.is_deleted === 1 ? 'Yes' : 'No'}</p>
                    </div>
                )
            }
            </div>


        );
    }

    return (
        <div>
            <h2>Manage Panel Users</h2>
            <table>
                <thead>
                <tr>
                    <th>Login</th>
                    <th>Info</th>
                    <th>Block</th>
                    <th>Delete</th>
                </tr>
                </thead>
                <tbody>
                {users.filter(user => user.is_deleted !== 1).map((user) => (
                    <tr key={user.id}>

                        <td>{user.login}</td>
                        <td>
                            <button onClick={() => handleInfoUser(user.id)}>Info</button>
                        </td>
                        <td>
                            {user.banned === 1 ? (
                                <button onClick={() => handleBanUnbanUser(user.id, 0)}>Unban</button>
                            ) : (
                                <button onClick={() => handleBanUnbanUser(user.id, 1)}>Ban</button>
                            )}
                        </td>
                        <td>


                            <button onClick={() => handleDeleteUser(user.id)}>Delete</button>

                        </td>
                    </tr>
                ))}
                </tbody>
            </table>

            <UserInfoDisplay userInfo={user}/>
        </div>

    );
}

export default ManagePanelUsers;